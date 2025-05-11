package validations

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"regexp"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
	"strconv"
	"strings"
)

func DisciplinaValida(disciplina *models.Disciplina, ctx *gin.Context) bool {
	return utils.BindAndValidate(disciplina, ctx)
}

func AlunoDisciplinaValido(alunoDisciplina *models.AlunoDisciplina, ctx *gin.Context) bool {
	return utils.BindAndValidate(alunoDisciplina, ctx)
}

func AnoSemestre(fl validator.FieldLevel) bool {
	data := fl.Field().String()

	match, _ := regexp.MatchString(`^\d{4}-(01|02)$`, data)
	if !match {
		return false
	}

	parts := strings.Split(data, "-")
	if len(parts) != 2 {
		return false
	}

	anoStr, semestreStr := parts[0], parts[1]
	ano, err := strconv.Atoi(anoStr)
	if err != nil {
		return false
	}

	if ano < 2021 {
		return false
	}

	if semestreStr != "01" && semestreStr != "02" {
		return false
	}

	return true
}
