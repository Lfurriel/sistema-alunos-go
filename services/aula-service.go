package services

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"sistema-alunos-go/database"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

// CadastrarAula registra uma nova aula para uma disciplina
//
// A função recebe os dados da aula e o ID da disciplina à qual ela pertence
// Antes de cadastrar, verifica se já existe uma aula com o mesmo número naquela disciplina
// Também atualiza a carga horária realizada da disciplina
//
// Retorna a aula cadastrada ou um erro, caso haja falha de validação ou de persistência
func CadastrarAula(aula *models.Aula, disciplinaId string) (*models.Aula, *utils.RestErr) {
	aula.DisciplinaId = disciplinaId

	disciplina, restErr := buscaDisciplina(aula.DisciplinaId)
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

	return aula, nil
}

// ListarAulasDisciplina retorna todas as aulas relacionadas a uma disciplina
//
// A função busca todas as aulas associadas ao ID da disciplina fornecida, incluindo a lista de presenças (`AlunoAula`)
// e os respectivos dados dos alunos
//
// Retorna uma lista de aulas com suas respectivas presenças, ou erro caso a disciplina não exista ou a consulta falhe.
func ListarAulasDisciplina(id string) ([]models.Aula, *utils.RestErr) {
	var err error

	if _, err := buscaDisciplina(id); err != nil {
		return nil, err
	}

	var aulas []models.Aula
	err = database.DB.Preload("AlunoAula.Aluno").Where("disciplina_id = ?", id).Find(&aulas).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar aulas", err)
	}

	return aulas, nil
}

// GetAula retorna os detalhes de uma aula específica.
//
// A função busca uma aula pelo seu ID, incluindo os registros de presença (`AlunoAula`) e os dados dos alunos presentes.
//
// Retorna a aula encontrada ou um erro caso ocorra falha na busca.
func GetAula(id string) (*models.Aula, *utils.RestErr) {
	var err error

	var aula *models.Aula
	err = database.DB.Preload("AlunoAula.Aluno").Where("id = ?", id).Find(&aula).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar aula", err)
	}

	return aula, nil
}
