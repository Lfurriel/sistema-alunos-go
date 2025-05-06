package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistema-alunos-go/models"
	"sistema-alunos-go/services"
	"sistema-alunos-go/utils"
	"sistema-alunos-go/validations"
)

// TODO: a aula deve receber os alunos presentes, e assim marcar a presença de cada aluno e falta dos que não aparecerem
func CadastrarAula(ctx *gin.Context) {
	var aula models.Aula
	if !validations.AulaValida(&aula, ctx) {
		return
	}

	result, restErr := services.CadastrarAula(aula)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Aula cadastrada com sucesso",
		http.StatusCreated,
		result,
	))
}

func ListarAulasDisciplina(ctx *gin.Context) {

}
