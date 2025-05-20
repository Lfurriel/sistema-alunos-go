package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sistema-alunos-go/utils"
)

// ErrorHandlingMiddleware handles and recovers from panics within middleware or handlers, logging errors and responding with HTTP 500.
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("\033[31m[ERRO INTERNO] panic recuperado: %v\033[0m\n", err)
				restErr := utils.NewRestErr(http.StatusInternalServerError, "Erro interno", nil)
				utils.RespondRestErr(restErr, ctx)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
