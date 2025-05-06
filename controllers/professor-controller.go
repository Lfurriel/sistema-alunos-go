package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistema-alunos-go/models"
	"sistema-alunos-go/services"
	"sistema-alunos-go/utils"
	"sistema-alunos-go/validations"
)

func CadastrarProfessor(ctx *gin.Context) {
	var professor models.Professor
	if !validations.ProfessorValido(&professor, ctx) {
		return
	}

	result, restErr := services.CadastrarProfessor(professor)
	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Professor cadastrado com sucesso",
		http.StatusCreated,
		result,
	))
}

func Login(ctx *gin.Context) {
	var login models.Login
	if !validations.LoginValido(&login, ctx) {
		return
	}
	token, cliente, restErr := services.Login(login)
	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"professor": cliente,
		"token":     token,
	})
}
