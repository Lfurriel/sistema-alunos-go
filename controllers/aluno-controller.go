package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistema-alunos-go/models"
	"sistema-alunos-go/services"
	"sistema-alunos-go/utils"
	"sistema-alunos-go/validations"
)

// CadastrarAluno trata a requisição de cadastro de um novo aluno.
//
// # Valida os dados recebidos via JSON no corpo da requisição, chama o serviço de cadastro e retorna o aluno criado com status 201
//
// Em caso de erro de validação ou persistência, retorna um erro estruturado.
func CadastrarAluno(ctx *gin.Context) {
	var aluno models.Aluno
	if !validations.AlunoValido(&aluno, ctx) {
		return
	}

	result, restErr := services.CadastrarAluno(aluno)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Aluno cadastrado com sucesso",
		http.StatusCreated,
		result,
	))
}

// DesativarAluno trata a requisição para desativar um aluno (trancar matrícula).
//
// O ID do aluno é obtido via parâmetro de rota. Marca o aluno como inativo e retorna a entidade atualizada com status 200
//
// Em caso de erro, retorna uma resposta padronizada com código e mensagem.
func DesativarAluno(ctx *gin.Context) {
	id := ctx.Param("id")

	result, restErr := services.AtualizarAluno(id, false)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusOK, utils.NewAppMessage(
		"Aluno desativado com sucesso",
		http.StatusOK,
		result,
	))
}

// ReativarAluno trata a requisição para reativar um aluno (destrancar matrícula).
//
// O ID do aluno é obtido via parâmetro de rota. Marca o aluno como ativo e retorna a entidade atualizada com status 200
//
// Em caso de erro, retorna uma resposta padronizada com código e mensagem.
func ReativarAluno(ctx *gin.Context) {
	id := ctx.Param("id")

	result, restErr := services.AtualizarAluno(id, true)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusOK, utils.NewAppMessage(
		"Aluno reativado com sucesso",
		http.StatusOK,
		result,
	))
}

// RemoverAluno trata a requisição de exclusão definitiva de um aluno
//
// O ID do aluno é obtido via parâmetro de rota. Remove o aluno do banco de dados e retorna status 204 (No Content)
func RemoverAluno(ctx *gin.Context) {
	id := ctx.Param("id")
	restErr := services.RemoverAluno(id)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusNoContent, utils.NewAppMessage(
		"Aluno removido com sucesso",
		http.StatusNoContent,
		nil,
	))
}
