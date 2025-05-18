package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistema-alunos-go/models"
	"sistema-alunos-go/services"
	"sistema-alunos-go/utils"
	"sistema-alunos-go/validations"
)

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

func getProfessorId(ctx *gin.Context) string {
	professor, exists := ctx.Get("professor")
	if !exists {
		restErr := utils.NewRestErr(http.StatusUnauthorized, "Professor não autenticado", nil)
		utils.RespondRestErr(restErr, ctx)
		return ""
	}

	return professor.(string)
}
