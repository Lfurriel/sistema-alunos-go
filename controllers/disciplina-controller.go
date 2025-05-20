package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistema-alunos-go/models"
	"sistema-alunos-go/services"
	"sistema-alunos-go/utils"
	"sistema-alunos-go/validations"
)

// CadastrarDisciplina trata a requisição de criação de uma nova disciplina.
//
// Obtém o ID do professor autenticado, valida os dados enviados no corpo da requisição e chama o serviço para salvar
// a disciplina
//
// Retorna a disciplina criada com status 201 ou um erro em caso de falha.
func CadastrarDisciplina(ctx *gin.Context) {
	professorId := getProfessorId(ctx)
	if professorId == "" {
		return
	}

	var disciplina models.Disciplina
	if !validations.DisciplinaValida(&disciplina, ctx) {
		return
	}

	result, restErr := services.CadastrarDisciplina(disciplina, professorId)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Disciplina cadastrada com sucesso",
		http.StatusCreated,
		result,
	))
}

// MatricularAluno associa um aluno a uma disciplina.
//
// Recebe os IDs via query string (`disciplinaId` e `alunoId`), chama o serviço e retorna o vínculo criado com status 201.
//
// Retorna erro em caso de falha na matrícula.
func MatricularAluno(ctx *gin.Context) {
	disciplinaId := ctx.Query("disciplinaId")
	alunoId := ctx.Query("alunoId")

	result, restErr := services.Matricular(disciplinaId, alunoId)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Aluno matriculado com sucesso",
		http.StatusCreated,
		result,
	))
}

// AdicionarAvaliacao registra uma nova avaliação (prova ou trabalho) para uma disciplina.
//
// O ID da disciplina é passado via parâmetro de rota Valida os dados da avaliação e chama o serviço responsável pelo cadastro.
//
// Retorna a avaliação criada com status 201 ou erro, se houver falha.
func AdicionarAvaliacao(ctx *gin.Context) {
	disciplinaId := ctx.Param("disciplinaId")

	var avaliacao models.Avaliacao
	if !validations.AvaliacaoValida(&avaliacao, ctx) {
		return
	}

	result, restErr := services.AdicionarAvaliacao(avaliacao, disciplinaId)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Avaliação cadastrada com sucesso",
		http.StatusCreated,
		result,
	))
}

// AdicionarNotaAvaliacao registra as notas dos alunos para uma avaliação específica
//
// Os IDs da disciplina e da avaliação são passados via rota. Recebe uma lista de notas no corpo da requisição e
// chama o serviço de persistência.
//
// Retorna as notas cadastradas com status 201 ou erro em caso de falha
func AdicionarNotaAvaliacao(ctx *gin.Context) {
	disciplinaId := ctx.Param("disciplinaId")
	avaliacaoId := ctx.Param("avaliacaoId")

	var alunosNota []models.AlunoAvaliacao
	if !validations.NotaValida(&alunosNota, ctx) {
		return
	}

	result, restErr := services.AdicionarNotaAvaliacao(alunosNota, avaliacaoId, disciplinaId)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Notas adicionadas com sucesso",
		http.StatusCreated,
		result,
	))
}

// ListarDisciplinas retorna todas as disciplinas do professor autenticado.
//
// Retorna a lista de disciplinas com status 201 ou erro em caso de falha.
func ListarDisciplinas(ctx *gin.Context) {
	professorId := getProfessorId(ctx)
	if professorId == "" {
		return
	}

	result, restErr := services.ListarDisciplinas(professorId)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Disciplina obtidas com sucesso",
		http.StatusCreated,
		result,
	))
}

// FecharSemestre finaliza o semestre de uma disciplina e calcula os resultados dos alunos.
//
// Retorna uma lista de resultados finais com status 201 ou erro.
func FecharSemestre(ctx *gin.Context) {
	disciplinaId := ctx.Param("disciplinaId")

	result, restErr := services.FecharSemestre(disciplinaId)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Disciplina obtidas com sucesso",
		http.StatusCreated,
		result,
	))
}

// getProfessorId é uma função auxiliar que extrai o ID do professor autenticado a partir do contexto da requisição
//
// # Utiliza os dados salvos pelo middleware de autenticação JWT
//
// Se não estiver presente, retorna erro 401 e encerra a execução do handler
func getProfessorId(ctx *gin.Context) string {
	professor, exists := ctx.Get("professor")
	if !exists {
		restErr := utils.NewRestErr(http.StatusUnauthorized, "Professor não autenticado", nil)
		utils.RespondRestErr(restErr, ctx)
		return ""
	}

	return professor.(string)
}
