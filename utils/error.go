package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

// RestErr representa uma estrutura padronizada para lidar com erros da API RESTful
// Inclui um código de status HTTP, uma mensagem amigável e um erro interno opcional
type RestErr struct {
	Code int
	Msg  string
	Err  error
}

// NewRestErr cria uma nova instância de RestErr contendo código HTTP, mensagem e erro interno
//
// Usada para representar erros personalizados em respostas JSON estruturadas
func NewRestErr(code int, msg string, err error) *RestErr {
	return &RestErr{Code: code, Msg: msg, Err: err}
}

// RespondRestErr envia uma resposta de erro JSON padronizada para o cliente.
//
// Caso um erro interno esteja presente, ele é logado no terminal com destaque.
// Em seguida, a resposta JSON é enviada com o código e mensagem definidos em RestErr.
// Se RestErr for nulo, retorna um erro 500 genérico.
//
// A execução do contexto do Gin é encerrada com ctx.Abort().
func RespondRestErr(err *RestErr, ctx *gin.Context) {
	if err != nil && err.Err != nil {
		log.Printf("\033[31m[ERRO INTERNO] panic recuperado: %v\033[0m\n", err)
	}

	if err != nil {
		ctx.JSON(err.Code, NewAppMessage(
			err.Msg,
			err.Code,
			nil,
		))
	} else {
		ctx.JSON(500, NewAppMessage(
			"Erro interno",
			500,
			nil,
		))
	}

	ctx.Abort()
}
