package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistema-alunos-go/models"
	"sistema-alunos-go/services"
	"sistema-alunos-go/utils"
	"sistema-alunos-go/validations"
)

// CadastrarAula trata a requisição de criação de uma nova aula para uma disciplina
//
// Valida o corpo da requisição, obtém o ID da disciplina via parâmetro de rota e chama o serviço para salvar a aula.
// Também registra a presença dos alunos
//
// Retorna a aula criada com status 201 ou erro, se houver falha
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

// ListarAulasDisciplina retorna todas as aulas cadastradas para uma disciplina específica
//
// O ID da disciplina é obtido via parâmetro de rota. A resposta inclui também os alunos presentes em cada aula
//
// Retorna status 201 com os dados ou erro em caso de falha
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

// GetAula retorna os dados detalhados de uma aula específica pelo seu ID
//
// # A resposta inclui a presença dos alunos e demais informações da aula
//
// Retorna status 201 com os dados ou erro em caso de falha
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
