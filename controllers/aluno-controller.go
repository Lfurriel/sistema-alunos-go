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

func DesativarAluno(ctx *gin.Context) {
	// TODO: atualizar a quantidade de alunos matriculados nas disciplinas que esse aluno esatava matriculado (-1)
	id := ctx.Param("id")

	result, restErr := services.AtualizarAluno(id, false)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
	}

	ctx.JSON(http.StatusOK, utils.NewAppMessage(
		"Auluno desativado com sucesso",
		http.StatusOK,
		result,
	))
}

func ReativarAluno(ctx *gin.Context) {
	// TODO: atualizar a quantidade de alunos matriculados nas disciplinas que esse aluno esatava matriculado (+1)
	id := ctx.Param("id")

	result, restErr := services.AtualizarAluno(id, true)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
	}

	ctx.JSON(http.StatusOK, utils.NewAppMessage(
		"Aluno reativado com sucesso",
		http.StatusOK,
		result,
	))
}

func RemoverAluno(ctx *gin.Context) {
	/* TODO: atualizar a quantidade de alunos matriculados nas disciplinas que esse aluno esatava matriculado (-1) e apagar os alunos-disciplina, aluno-avaliacao etc (talvez o delete on cascade resolva)*/
	id := ctx.Param("id")
	restErr := services.RemoverAluno(id)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
	}

	ctx.JSON(http.StatusNoContent, utils.NewAppMessage(
		"Aluno removido com sucesso",
		http.StatusNoContent,
		nil,
	))
}
