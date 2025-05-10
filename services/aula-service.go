package services

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"sistema-alunos-go/database"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

func CadastrarAula(aula *models.Aula) (*models.Aula, *utils.RestErr) {
	disciplina, restErr := getDisciplina(aula.DisciplinaId)
	if restErr != nil {
		return nil, restErr
	}

	disciplina.CargaHorariaRealizada += aula.QuantidadeHoras

	if err := database.DB.Save(disciplina).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao atualizar disciplina", err)
	}

	var aulaExist models.Aula
	err := database.DB.Where("disciplina_id = ? AND numero = ?", aula.DisciplinaId, aula.Numero).First(&aulaExist).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar aulas", err)
	}
	if aulaExist.Id != "" {
		return nil, utils.NewRestErr(http.StatusBadRequest, "Aula com esse número já cadastrada para a disciplina", nil)
	}

	if err := database.DB.Create(aula).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao criar aula", err)
	}

	var alunosMatriculados []models.AlunoDisciplina
	err = database.DB.Where("disciplina_id = ?", aula.DisciplinaId).Find(&alunosMatriculados).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar alunos da disciplina", err)
	}

	presentesMap := map[string]bool{}
	for _, aa := range aula.AlunoAula {
		presentesMap[aa.AlunoId] = true
	}

	var registrosPresenca []models.AlunoAula
	for _, am := range alunosMatriculados {
		registrosPresenca = append(registrosPresenca, models.AlunoAula{
			AulaId:   aula.Id,
			AlunoId:  am.AlunoId,
			Presenca: presentesMap[am.AlunoId],
		})
	}

	err = database.DB.Create(&registrosPresenca).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao registrar presença dos alunos", err)
	}

	return aula, nil
}

func ListarAulasDisciplina(id string) ([]models.Aula, *utils.RestErr) {
	var err error

	if _, err := getDisciplina(id); err != nil {
		return nil, err
	}

	var aulas []models.Aula
	err = database.DB.Where("disciplina_id = ?", id).Find(&aulas).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar aulas", err)
	}

	return aulas, nil
}

func GetAula(id string) (*models.Aula, *utils.RestErr) {
	var err error

	var aula *models.Aula
	err = database.DB.Find(&aula, id).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar aulas", err)
	}

	return aula, nil
}

func getDisciplina(id string) (*models.Disciplina, *utils.RestErr) {
	var err error
	var disciplina models.Disciplina
	err = database.DB.First(&disciplina, id).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar disciplina", err)
		}
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Disciplina não existe", err)
	}
	return &disciplina, nil
}
