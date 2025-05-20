package validations

import (
	"github.com/gin-gonic/gin"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

// AvaliacaoValida valida os campos de um objeto Avaliacao com base nas regras definidas, retornando true para dados válidos.
func AvaliacaoValida(avaliacao *models.Avaliacao, ctx *gin.Context) bool {
	return utils.BindAndValidate(avaliacao, ctx)
}

// NotaValida valida os campos de um objeto AlunoAvaliacao com base nas regras definidas, retornando true para dados válidos.
func NotaValida(alunoNota *[]models.AlunoAvaliacao, ctx *gin.Context) bool {
	return utils.BindAndValidate(alunoNota, ctx)
}
