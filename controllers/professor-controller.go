package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistema-alunos-go/models"
	"sistema-alunos-go/services"
	"sistema-alunos-go/utils"
	"sistema-alunos-go/validations"
)

// CadastrarProfessor trata a requisição de criação de um novo professor
//
// # Valida os dados recebidos no corpo da requisição, chama o serviço de cadastro e retorna o professor criado com status 201
//
// Retorna erro em caso de falha na validação ou persistência
func CadastrarProfessor(ctx *gin.Context) {
	var professor models.Professor
	if !validations.ProfessorValido(&professor, ctx) {
		return
	}

	result, restErr := services.CadastrarProfessor(professor)
	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Professor cadastrado com sucesso",
		http.StatusCreated,
		result,
	))
}

// Login autentica um professor com base em suas credenciais
//
// Valida o corpo da requisição (email e senha), chama o serviço de autenticação e retorna um token JWT juntamente com os dados do professor (sem senha)
//
// Retorna erro 401 se as credenciais forem inválidas
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

// RemoverProfessor trata a requisição de exclusão de um professor do sistema
//
// # O ID é obtido via parâmetro de rota
//
// Remove o professor do banco de dados, e retorna status 204 (No Content) se a exclusão for bem-sucedida
func RemoverProfessor(ctx *gin.Context) {
	id := ctx.Param("id")

	if restErr := services.RemoverProfessor(id); restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusNoContent, utils.NewAppMessage(
		"Professor removido com sucesso",
		http.StatusNoContent,
		nil,
	))
}
