package validations

import (
	"github.com/gin-gonic/gin"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

// AulaValida valida os campos de um objeto Aula com base nas regras definidas, retornando true para dados válidos.
func AulaValida(aula *models.Aula, ctx *gin.Context) bool {
	return utils.BindAndValidate(aula, ctx)
}
