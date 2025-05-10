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
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Disciplina cadastrada com sucesso",
		http.StatusCreated,
		result,
	))
}

func MatricularAluno(ctx *gin.Context) {
	var alunoDisciplina models.AlunoDisciplina
	if !validations.AlunoDisciplinaValido(&alunoDisciplina, ctx) {
		return
	}

	result, restErr := services.Matricular(alunoDisciplina)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Disciplina cadastrada com sucesso",
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
	}

	ctx.JSON(http.StatusCreated, utils.NewAppMessage(
		"Disciplina obtidas com sucesso",
		http.StatusCreated,
		result,
	))
}

func FecharSemestre(ctx *gin.Context) {
	disciplinaId := ctx.Param("id")

	result, restErr := services.FecharSemestre(disciplinaId)

	if restErr != nil {
		utils.RespondRestErr(restErr, ctx)
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
