package validations

import (
	"github.com/gin-gonic/gin"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

func AvaliacaoValida(avaliacao *models.Avaliacao, ctx *gin.Context) bool {
	return utils.BindAndValidate(avaliacao, ctx)
}

func NotaValida(alunoNota *[]models.AlunoAvaliacao, ctx *gin.Context) bool {
	return utils.BindAndValidate(alunoNota, ctx) //
}
