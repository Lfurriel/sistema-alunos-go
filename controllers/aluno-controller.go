package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistema-alunos-go/models"
	"sistema-alunos-go/services"
	"sistema-alunos-go/utils"
	"sistema-alunos-go/validations"
)

func CadastrarAluno(ctx *gin.Context) {
	var aluno models.Aluno
	if !validations.AlunoValido(&aluno, ctx) {
		return
	}

	result, restErr := services.CadastrarAluno(aluno)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Aula cadastrada com sucesso",
		http.StatusCreated,
		result,
	))
}

func TrancarAluno(ctx *gin.Context) {
	id := ctx.Param("id")

	result, restErr := services.AtualizarAluno(id, true)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Aula cadastrada com sucesso",
		http.StatusCreated,
		result,
	))
}

func DestrancarAluno(ctx *gin.Context) {
	id := ctx.Param("id")

	result, restErr := services.AtualizarAluno(id, false)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Aula cadastrada com sucesso",
		http.StatusCreated,
		result,
	))
}
