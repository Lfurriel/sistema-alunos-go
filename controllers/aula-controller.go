package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistema-alunos-go/models"
	"sistema-alunos-go/services"
	"sistema-alunos-go/utils"
	"sistema-alunos-go/validations"
)

func CadastrarAula(ctx *gin.Context) {
	disciplinaId := ctx.Param("disciplinaId")

	var aula models.Aula
	if !validations.AulaValida(&aula, ctx) {
		return
	}

	result, restErr := services.CadastrarAula(&aula, disciplinaId)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Aula cadastrada com sucesso",
		http.StatusCreated,
		result,
	))
}

func ListarAulasDisciplina(ctx *gin.Context) {
	disciplinaId := ctx.Param("disciplinaId")

	result, restErr := services.ListarAulasDisciplina(disciplinaId)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Aula resgatadas com sucesso",
		http.StatusCreated,
		result,
	))
}

func GetAula(ctx *gin.Context) {
	id := ctx.Param("id")

	result, restErr := services.GetAula(id)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Aula resgatadas com sucesso",
		http.StatusCreated,
		result,
	))
}
