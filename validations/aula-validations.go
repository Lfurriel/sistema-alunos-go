package validations

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

func AulaValida(aula *models.Aula, ctx *gin.Context) bool {
	if err := ctx.ShouldBindJSON(&aula); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			var errorsList []utils.ValidationError
			for _, e := range validationErrors {
				errorsList = append(errorsList, utils.MapValidationError(e))
			}

			response := utils.NewAppMessage("Erro de Validação", http.StatusBadRequest, nil, errorsList)
			ctx.JSON(http.StatusBadRequest, response)
			return false
		}

		response := utils.NewAppMessage("Dados inválidos", http.StatusBadRequest, nil, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return false
	}

	return true
}
