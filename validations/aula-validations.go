package validations

import (
	"github.com/gin-gonic/gin"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

func AulaValida(aula *models.Aula, ctx *gin.Context) bool {
	return utils.BindAndValidate(aula, ctx)
}
