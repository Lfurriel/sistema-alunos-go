package services

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"sistema-alunos-go/database"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

func CadastrarAula(aula models.Aula) (*models.Aula, *utils.RestErr) {
	var err error

	var disciplina models.Disciplina
	err = database.DB.First(&disciplina, aula.DisciplinaId).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar disciplina", err)
		}
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Disciplina não existe", err)
	}

	var aulaExist *models.Aula
	err = database.DB.Where("numero = ?", aula.Numero).First(&aulaExist).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar aulas", err)
	}

	if aulaExist != nil && aulaExist.Id != "" {
		return nil, utils.NewRestErr(http.StatusBadRequest, "Aula com esse número já cadastrada", nil)
	}

	err = database.DB.Create(&aula).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao criar professor", err)
	}

	return &aula, nil
}
