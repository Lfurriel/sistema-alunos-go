package services

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"sistema-alunos-go/database"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

func CadastrarAluno(aluno models.Aluno) (*models.Aluno, *utils.RestErr) {
	aluno.Ativo = false
	err := database.DB.Create(&aluno).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao cadastrar aluno", err)
	}
	return &aluno, nil
}

func AtualizarAluno(alunoId string, ativo bool) (*models.Aluno, *utils.RestErr) {
	var err error
	aluno := models.Aluno{}

	err = database.DB.Find(&aluno, alunoId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRestErr(http.StatusNotFound, "Aluno não encontrado", err)
		}
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar professor", err)
	}

	aluno.Ativo = ativo
	err = database.DB.Save(&aluno).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao atualizar aluno", err)
	}
	return &aluno, nil
}
