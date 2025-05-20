package configs

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"sistema-alunos-go/validations"
)

// BindingValidator registra funções de validação personalizadas para "senha_forte", "data_valida" e "ano_semestre"
func BindingValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("senha_forte", validations.SenhaForte); err != nil {
			log.Printf("Erro ao registrar validação 'senha_forte': %v", err)
		}
		if err := v.RegisterValidation("data_valida", validations.DataValida); err != nil {
			log.Printf("Erro ao registrar validação 'data_valida': %v", err)
		}
		if err := v.RegisterValidation("ano_semestre", validations.AnoSemestre); err != nil {
			log.Printf("Erro ao registrar validação 'ano_semestre': %v", err)
		}
	}
}
