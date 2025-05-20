package validations

import (
	"github.com/gin-gonic/gin"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

// AlunoValido valida os campos de um objeto Aluno com base nas regras definidas, retornando true para dados válidos.
func AlunoValido(aluno *models.Aluno, ctx *gin.Context) bool {
	return utils.BindAndValidate(aluno, ctx)
}
