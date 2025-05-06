package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistema-alunos-go/utils"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				restErr := utils.NewRestErr(http.StatusInternalServerError, "Erro interno", nil)
				utils.RespondRestErr(restErr, ctx)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
