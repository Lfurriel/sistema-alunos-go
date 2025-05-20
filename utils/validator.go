package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

// ValidationError representa um erro de validação de campo em uma requisição.
//
// Contém informações sobre o campo inválido, o valor esperado e uma mensagem de erro amigável.
type ValidationError struct {
	Expected string `json:"expected,omitempty"`
	Path     string `json:"path,omitempty"`
	Message  string `json:"message,omitempty"`
}

// MapValidationError converte um erro de validação do pacote validator para o formato personalizado ValidationError
//
// Analisa a tag de validação (`err.Tag()`) e retorna uma estrutura contendo mensagens de erro e valores esperados
// mais legíveis para o cliente
func MapValidationError(err validator.FieldError) ValidationError {
	var validationError ValidationError
	validationError.Message = "Required"
	validationError.Path = err.Field()

	switch err.Tag() {
	case "required":
		validationError.Expected = "valor obrigatório"
	case "email":
		validationError.Expected = "email válido"
		validationError.Message = "O campo deve conter um email válido."
	case "len":
		validationError.Expected = err.Param() + " caracteres"
		validationError.Message = "O campo deve ter no exatos " + err.Param() + " caracteres."
	case "min":
		validationError.Expected = "mínimo de " + err.Param() + " caracteres"
		validationError.Message = "O campo deve ter no mínimo " + err.Param() + " caracteres."
	case "max":
		validationError.Expected = "máximo de " + err.Param() + " caracteres"
		validationError.Message = "O campo deve ter no máximo " + err.Param() + " caracteres."
	case "gte":
		validationError.Expected = "número maior que " + err.Param()
		validationError.Message = "O número deve ser maior que " + err.Param()
	case "lte":
		validationError.Expected = "número menor que " + err.Param()
		validationError.Message = "O número deve ser menor que " + err.Param()
	case "number":
		validationError.Expected = "numero"
		validationError.Message = "O campo deve ser um numero"
	case "numeric":
		validationError.Expected = "numero"
		validationError.Message = "O campo deve ser uma string numérica"
	case "oneof":
		validationError.Expected = strings.ReplaceAll(err.Param(), " ", " | ")
		validationError.Message = "O valor deve ser um dos seguintes: " + validationError.Expected
	case "senha_forte":
		validationError.Expected = "mínimo de 8 caracteres, uma letra maiúscula, uma minúscula e um símbolo"
		validationError.Message = "A senha deve ter pelo menos 8 caracteres, incluindo uma letra maiúscula, uma minúscula, um número e um símbolo."
	case "data_valida":
		validationError.Expected = "Data de nascimento válida"
		validationError.Message = "A data de nascimento deve estar no padrão yyyy-MM-dd"
	case "ano_semestre":
		validationError.Expected = "Ano-Semestre válido"
		validationError.Message = "O campo ano-semestre deve estar no formato yyyy-01 ou yyyy-02"
	}

	return validationError
}

// BindAndValidate realiza o bind dos dados do corpo da requisição JSON para a struct fornecida e
// valida os campos com base nas tags de validação
//
// Em caso de erro de validação, envia uma resposta JSON com a lista de erros formatados
// Em caso de erro genérico de bind, retorna uma mensagem genérica
//
// Retorna true se os dados forem válidos; caso contrário, false.
func BindAndValidate[T any](obj *T, ctx *gin.Context) bool {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			var errorsList []ValidationError
			for _, e := range validationErrors {
				errorsList = append(errorsList, MapValidationError(e))
			}

			ctx.JSON(http.StatusBadRequest, NewAppMessage("Erro de Validação", http.StatusBadRequest, nil, errorsList))
			return false
		}

		ctx.JSON(http.StatusBadRequest, NewAppMessage("Dados inválidos", http.StatusBadRequest, nil, err.Error()))
		return false
	}
	return true
}
