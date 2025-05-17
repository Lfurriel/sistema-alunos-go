package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

type RestErr struct {
	Code int
	Msg  string
	Err  error
}

func NewRestErr(code int, msg string, err error) *RestErr {
	return &RestErr{Code: code, Msg: msg, Err: err}
}

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
