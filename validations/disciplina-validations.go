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

// DisciplinaValida valida os campos de um objeto Disciplina com base nas regras definidas, retornando true para dados válidos.
func DisciplinaValida(disciplina *models.Disciplina, ctx *gin.Context) bool {
	return utils.BindAndValidate(disciplina, ctx)
}

// AnoSemestre valida se uma string representa um formato ano-semestre válido (AAAA-01 ou AAAA-02) a partir de 2021.
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
