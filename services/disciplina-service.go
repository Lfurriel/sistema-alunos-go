package services

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"sistema-alunos-go/database"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

func CadastrarDisciplina(disciplina models.Disciplina, professorId string) (*models.Disciplina, *utils.RestErr) {
	disciplina.ProfessorId = professorId
	if err := database.DB.Create(&disciplina).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao cadastrar disciplina", err)
	}
	return &disciplina, nil
}

func Matricular(alunoDisciplina models.AlunoDisciplina) (*models.AlunoDisciplina, *utils.RestErr) {
	disciplina, restErr := buscaDisciplina(alunoDisciplina.DisciplinaId)
	if restErr != nil {
		return nil, restErr
	}

	if _, restErr := buscaAluno(alunoDisciplina.AlunoId); restErr != nil {
		return nil, restErr
	}

	if err := database.DB.Create(&alunoDisciplina).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao matricular aluno", err)
	}

	disciplina.QuantidadeAlunos = disciplina.QuantidadeAlunos + 1
	if err := database.DB.Save(&disciplina).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao atualizar quantidade de alunos", err)
	}

	return &alunoDisciplina, nil
}

func ListarDisciplinas(professorId string) ([]models.Disciplina, *utils.RestErr) {
	var disciplinas []models.Disciplina
	err := database.DB.Where("professor_id = ?", professorId).Find(&disciplinas).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar disciplina", err)
	}
	return disciplinas, nil
}

func FecharSemestre(disciplinaId string) ([]models.AlunoMedia, *utils.RestErr) {
	disciplina, restErr := buscaDisciplina(disciplinaId)
	if restErr != nil {
		return nil, restErr
	}

	if disciplina.CargaHorariaRealizada < disciplina.CargaHorariaPrevista {
		return nil, utils.NewRestErr(400, "Carga horária realizada menor que a prevista", nil)
	}

	var alunosDisciplina []models.AlunoDisciplina
	if err := database.DB.Where("disciplina_id = ?", disciplinaId).Find(&alunosDisciplina).Error; err != nil {
		return nil, utils.NewRestErr(500, "Erro ao buscar alunos da disciplina", err)
	}

	var aulas []models.Aula
	if err := database.DB.Where("disciplina_id = ?", disciplinaId).Find(&aulas).Error; err != nil {
		return nil, utils.NewRestErr(500, "Erro ao buscar aulas da disciplina", err)
	}

	totalAulas := len(aulas)
	if totalAulas == 0 {
		return nil, utils.NewRestErr(400, "Disciplina não possui aulas registradas", nil)
	}

	var avaliacoes []models.Avaliacao
	if err := database.DB.Where("disciplina_id = ?", disciplinaId).Find(&avaliacoes).Error; err != nil {
		return nil, utils.NewRestErr(500, "Erro ao buscar avaliações da disciplina", err)
	}

	if len(avaliacoes) == 0 {
		return nil, utils.NewRestErr(400, "A disciplina não possui avaliações cadastradas", nil)
	}

	var medias []models.AlunoMedia

	for _, ad := range alunosDisciplina {
		var presencas int64
		if err := database.DB.
			Model(&models.AlunoAula{}).
			Where("aluno_id = ? AND presenca = true AND aula_id IN (?)", ad.AlunoId, extractAulaIds(aulas)).
			Count(&presencas).Error; err != nil {
			return nil, utils.NewRestErr(500, "Erro ao calcular frequência", err)
		}
		frequencia := float64(presencas) / float64(totalAulas) * 100

		var notas []models.AlunoAvaliacao
		if err := database.DB.
			Where("aluno_id = ? AND avaliacao_id IN (?)", ad.AlunoId, extractAvaliacaoIds(avaliacoes)).
			Find(&notas).Error; err != nil {
			return nil, utils.NewRestErr(500, "Erro ao buscar notas do aluno", err)
		}

		var soma float64
		var pesoTotal float64

		for _, avaliacao := range avaliacoes {
			nota := findNota(notas, avaliacao.Id)
			soma += nota * avaliacao.Peso
			pesoTotal += avaliacao.Peso
		}

		if pesoTotal == 0 {
			return nil, utils.NewRestErr(400, "Peso total das avaliações é zero", nil)
		}
		mediaFinal := soma / pesoTotal

		aprovado := mediaFinal >= disciplina.NotaMinima && frequencia >= disciplina.FrequenciaMinima
		medias = append(medias, models.AlunoMedia{
			AlunoId:      ad.AlunoId,
			DisciplinaId: disciplinaId,
			MediaFinal:   mediaFinal,
			Frequencia:   frequencia,
			Aprovado:     aprovado,
		})
	}

	if err := database.DB.Create(&medias).Error; err != nil {
		return nil, utils.NewRestErr(500, "Erro ao salvar médias dos alunos", err)
	}

	return medias, nil
}

func buscaDisciplina(id string) (*models.Disciplina, *utils.RestErr) {
	var disciplina models.Disciplina
	err := database.DB.Find(&disciplina, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRestErr(http.StatusNotFound, "Disciplina não encontrada", err)
		}
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar disciplina", err)
	}

	return &disciplina, nil
}

func extractAulaIds(aulas []models.Aula) []string {
	var ids []string
	for _, a := range aulas {
		ids = append(ids, a.Id)
	}
	return ids
}

func extractAvaliacaoIds(avaliacoes []models.Avaliacao) []string {
	var ids []string
	for _, a := range avaliacoes {
		ids = append(ids, a.Id)
	}
	return ids
}

func findNota(notas []models.AlunoAvaliacao, avaliacaoId string) float64 {
	for _, n := range notas {
		if n.AvaliacaoId == avaliacaoId {
			return n.Nota
		}
	}
	return 0.0
}
