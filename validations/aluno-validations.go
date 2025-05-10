package validations

import (
	"github.com/gin-gonic/gin"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

func AlunoValido(aluno *models.Aluno, ctx *gin.Context) bool {
	return utils.BindAndValidate(ctx, aluno)
}
