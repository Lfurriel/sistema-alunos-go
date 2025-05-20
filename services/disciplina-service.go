package services

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"sistema-alunos-go/database"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

// CadastrarDisciplina registra uma nova disciplina associada a um professor
//
// Define o ID do professor na disciplina e salva no banco
// Após o cadastro, busca os dados do professor para retornar no payload
//
// Retorna a disciplina criada ou erro, caso ocorra falha ao salvar ou buscar dados
func CadastrarDisciplina(disciplina models.Disciplina, professorId string) (*models.Disciplina, *utils.RestErr) {
	disciplina.ProfessorId = professorId
	if err := database.DB.Create(&disciplina).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao cadastrar disciplina", err)
	}

	var restErr *utils.RestErr
	disciplina.Professor, restErr = buscaProfessor(professorId)
	disciplina.Professor.Senha = ""
	disciplina.Professor.Disciplinas = nil

	if restErr != nil {
		return nil, restErr
	}

	return &disciplina, nil
}

// Matricular associa um aluno a uma disciplina
//
// Cria um registro em aluno_disciplina e atualiza o contador de alunos da disciplina
//
// Retorna o vínculo criado ou erro em caso de falha
func Matricular(disciplinaId string, alunoId string) (*models.AlunoDisciplina, *utils.RestErr) {
	disciplina, restErr := buscaDisciplina(disciplinaId)
	if restErr != nil {
		return nil, restErr
	}

	if _, restErr := buscaAluno(alunoId); restErr != nil {
		return nil, restErr
	}

	alunoDisciplina := models.AlunoDisciplina{
		DisciplinaId: disciplina.Id,
		AlunoId:      alunoId,
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

// AdicionarAvaliacao adiciona uma nova avaliação (prova ou trabalho) a uma disciplina
//
// # Atualiza os contadores de provas ou trabalhos na disciplina com base no tipo de avaliação
//
// Retorna a avaliação criada ou erro em caso de falha
func AdicionarAvaliacao(avaliacao models.Avaliacao, disciplinaId string) (*models.Avaliacao, *utils.RestErr) {
	disciplina, restErr := buscaDisciplina(disciplinaId)
	if restErr != nil {
		return nil, restErr
	}

	avaliacao.DisciplinaId = disciplina.Id

	if avaliacao.Tipo == "P" {
		disciplina.QuantidadeProvas += 1
	} else {
		disciplina.QuantidadeTrabalhos += 1
	}

	if err := database.DB.Save(&disciplina).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao atualizar disciplina", err)

	}

	if err := database.DB.Create(&avaliacao).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao inserir avaliação", err)

	}

	return &avaliacao, nil
}

// AdicionarNotaAvaliacao associa uma lista de notas de alunos a uma determinada avaliação
//
// A função insere múltiplos registros na tabela aluno_avaliacao com as notas fornecidas
//
// Retorna a lista salva ou erro em caso de falha de validação ou persistência
func AdicionarNotaAvaliacao(alunosNota []models.AlunoAvaliacao, avaliacaoId string, disciplinaId string) ([]models.AlunoAvaliacao, *utils.RestErr) {
	if _, restErr := buscaDisciplina(disciplinaId); restErr != nil {
		return nil, restErr
	}

	if _, restErr := buscaAvaliacao(avaliacaoId); restErr != nil {
		return nil, restErr
	}

	for i := range alunosNota {
		alunosNota[i].AvaliacaoId = avaliacaoId
		alunosNota[i].DisciplinaId = disciplinaId
	}

	if err := database.DB.Create(&alunosNota).Error; err != nil {
		return nil, utils.NewRestErr(500, "Erro ao salvar notas dos alunos", err)
	}

	return alunosNota, nil
}

// ListarDisciplinas retorna todas as disciplinas associadas a um professor
//
// # A resposta inclui as relações com alunos, aulas e avaliações
//
// Retorna uma lista de disciplinas ou erro em caso de falha
func ListarDisciplinas(professorId string) ([]models.Disciplina, *utils.RestErr) {
	var disciplinas []models.Disciplina

	err := database.DB.Preload("Alunos").Preload("Aulas").Preload("Avaliacoes").
		Where("professor_id = ?", professorId).Find(&disciplinas).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar disciplinas", err)
	}

	return disciplinas, nil
}

// FecharSemestre finaliza o semestre de uma disciplina calculando média e frequência dos alunos
//
// # Calcula a média ponderada com base nas avaliações e a frequência baseada nas presenças
//
// Retorna a lista de AlunoMedia com aprovação e dados finais ou erro em caso de falha
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

// buscaDisciplina é uma função auxiliar para buscar uma disciplina pelo ID
//
// Retorna a disciplina encontrada ou erro, caso não exista ou ocorra falha na consulta
func buscaDisciplina(id string) (*models.Disciplina, *utils.RestErr) {
	var disciplina models.Disciplina
	err := database.DB.Where("id = ?", id).First(&disciplina).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRestErr(http.StatusNotFound, "Disciplina não encontrada", err)
		}
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar disciplina", err)
	}

	return &disciplina, nil
}

// extractAulaIds extrai os IDs de uma lista de aulas
//
// Retorna um slice de strings contendo os IDs das aulas fornecidas
func extractAulaIds(aulas []models.Aula) []string {
	var ids []string
	for _, a := range aulas {
		ids = append(ids, a.Id)
	}
	return ids
}

// extractAvaliacaoIds extrai os IDs de uma lista de avaliações
//
// Retorna um slice de strings contendo os IDs das avaliações fornecidas
func extractAvaliacaoIds(avaliacoes []models.Avaliacao) []string {
	var ids []string
	for _, a := range avaliacoes {
		ids = append(ids, a.Id)
	}
	return ids
}

// findNota busca a nota atribuída a um aluno para uma avaliação específica
//
// Percorre a lista de notas e retorna a nota correspondente ao ID da avaliação fornecida.
//
// Se não encontrar, retorna 0.0
func findNota(notas []models.AlunoAvaliacao, avaliacaoId string) float64 {
	for _, n := range notas {
		if n.AvaliacaoId == avaliacaoId {
			return n.Nota
		}
	}
	return 0.0
}

// buscaAvaliacao recupera uma avaliação pelo ID
//
// Se a avaliação não for encontrada, retorna erro 404
// Em caso de falha de consulta, retorna erro interno
func buscaAvaliacao(id string) (*models.Avaliacao, *utils.RestErr) {
	var avaliacao models.Avaliacao
	if err := database.DB.Where("id = ?", id).Find(&avaliacao).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRestErr(http.StatusNotFound, "Avaliação não encontrada", err)
		}
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar avaliação", err)
	}

	return &avaliacao, nil
}
