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

// ProfessorValido valida os campos de um objeto Aluno com base nas regras definidas, além de confirmar se a senhas em
// 'senha' e 'confirmar_senha' são iguais, retornando true para dados válidos.
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

// LoginValido valida os campos de um objeto Login com base nas regras definidas, retornando true para dados válidos.
func LoginValido(login *models.Login, ctx *gin.Context) bool {
	return utils.BindAndValidate(login, ctx)
}

// SenhaForte verifica se uma senha atende aos critérios de força: mínimo 8 caracteres, letras maiúsculas e minúsculas, números e símbolos.
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

// DataValida verifica se uma string representa uma data válida no formato "YYYY-MM-DD".
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

// confirmaSenhasIguais verifica se os campos 'Senha' e 'ConfirmarSenha' de um Professor são iguais, retornando erro se não forem.
func confirmaSenhasIguais(p *models.Professor) error {
	if p.Senha != p.ConfirmarSenha {
		return errors.New("as duas senhas devem ser iguais")
	}
	return nil
}
