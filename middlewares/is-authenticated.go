package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"sistema-alunos-go/utils"
	"strings"
)

func IsAuthenticated(ctx *gin.Context) {
	secret := os.Getenv("JWT_SECRET")
	tokenValue := removeBearerPrefix(ctx.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, errors.New("token inválido")
	})

	if err != nil {
		restErr := utils.NewRestErr(http.StatusUnauthorized, "Token inválido", err)
		utils.RespondRestErr(restErr, ctx)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	professorData, ok := claims["professor"].(map[string]interface{})
	if !ok {
		restErr := utils.NewRestErr(http.StatusUnauthorized, "Formato inválido do token", nil)
		utils.RespondRestErr(restErr, ctx)
		return
	}

	professorId, exists := professorData["id"].(string)
	if !exists {
		restErr := utils.NewRestErr(http.StatusUnauthorized, "Token sem Id do professor", nil)
		utils.RespondRestErr(restErr, ctx)
		return
	}

	ctx.Set("professor", professorId)
}

func removeBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
