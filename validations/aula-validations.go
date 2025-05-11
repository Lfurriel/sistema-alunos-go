package validations

import (
	"github.com/gin-gonic/gin"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

func AulaValida(aula *models.Aula, ctx *gin.Context) bool {
	return utils.BindAndValidate(aula, ctx)
}

func NotaValida(alunoNota *[]models.AlunoAvaliacao, ctx *gin.Context) bool {
	return utils.BindAndValidate(alunoNota, ctx) //
}
