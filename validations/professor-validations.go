package validations

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
	"time"
)

func ProfessorValido(professor *models.Professor, ctx *gin.Context) bool {
	if !utils.BindAndValidate(professor, ctx) {
		return false
	}

	if err := confirmaSenhasIguais(professor); err != nil {
		response := utils.NewAppMessage("Erro de Validação", http.StatusBadRequest, nil, []map[string]interface{}{
			{
				"expected": "Senhas iguais",
				"message":  err.Error(),
			},
		})
		ctx.JSON(http.StatusBadRequest, response)
		return false
	}
	return true
}

func SenhaForte(fl validator.FieldLevel) bool {
	senha := fl.Field().String()

	if len(senha) < 8 {
		return false
	}

	temMinuscula, _ := regexp.MatchString(`[a-z]`, senha)
	if !temMinuscula {
		return false
	}

	temMaiuscula, _ := regexp.MatchString(`[A-Z]`, senha)
	if !temMaiuscula {
		return false
	}

	temNumero, _ := regexp.Match(`[0-9]`, []byte(senha))
	if !temNumero {
		return false
	}

	temSimbolo, _ := regexp.MatchString(`[!@#$%^&*()_+\-=[\]{};:'",.<>/?\\|]`, senha)
	if !temSimbolo {
		return false
	}

	return true
}

func DataValida(fl validator.FieldLevel) bool {
	data := fl.Field().String()
	match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, data)
	if !match {
		return false
	}

	_, err := time.Parse("2006-01-02", data)
	if err != nil {
		return false
	}

	return true
}

func confirmaSenhasIguais(p *models.Professor) error {
	if p.Senha != p.ConfirmarSenha {
		return errors.New("as duas senhas devem ser iguais")
	}
	return nil
}

func LoginValido(login *models.Login, ctx *gin.Context) bool {
	if err := ctx.ShouldBindJSON(&login); err != nil {
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
