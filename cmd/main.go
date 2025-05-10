package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"sistema-alunos-go/configs"
	"sistema-alunos-go/database"
	middleware "sistema-alunos-go/middlewares"
	"sistema-alunos-go/routes"
)

func init() {
	configs.LoadEnv()
	database.ConectaBD()
	configs.BindingValidator()
}

func main() {
	r := gin.Default()
	r.Use(middleware.ErrorHandlingMiddleware())

	routes.RegistraRotas(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		panic("Erro ao iniciar o servidor: " + err.Error())
	}
	log.Printf("Servidor rodando na porta %s...", port)
}
