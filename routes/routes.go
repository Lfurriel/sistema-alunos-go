package routes

import (
	"github.com/gin-gonic/gin"
	"sistema-alunos-go/controllers"
	middleware "sistema-alunos-go/middlewares"
)

func RegistraRotas(router *gin.Engine) {
	api := router.Group("")

	{
		aluno := api.Group("/aluno")
		aluno.POST("/", middleware.IsAuthenticated, controllers.CadastrarAluno)
		aluno.GET("/desativar/:id", middleware.IsAuthenticated, controllers.DesativarAluno)
		aluno.GET("/reativar/:id", middleware.IsAuthenticated, controllers.ReativarAluno)
		aluno.DELETE("/:id", middleware.IsAuthenticated, controllers.RemoverAluno)
	}

	{
		aula := api.Group("/aula")
		aula.POST("/:disciplinaId", middleware.IsAuthenticated, controllers.CadastrarAula)
		aula.GET("/disciplina/:disciplinaId", middleware.IsAuthenticated, controllers.ListarAulasDisciplina)
		aula.GET("/:id", middleware.IsAuthenticated, controllers.GetAula)
	}

	{
		disciplina := api.Group("disciplina")
		disciplina.POST("/", middleware.IsAuthenticated, controllers.CadastrarDisciplina)
		disciplina.POST("/matricular", middleware.IsAuthenticated, controllers.MatricularAluno)
		disciplina.POST("/avaliacao/:disciplinaId", middleware.IsAuthenticated, controllers.AdicionarAvaliacao)
		disciplina.POST("/avaliacao/:disciplinaId/nota/:avaliacaoId", middleware.IsAuthenticated, controllers.AdicionarNotaAvaliacao)
		disciplina.GET("/", middleware.IsAuthenticated, controllers.ListarDisciplinas)
		disciplina.GET("/fechar-semestre/:disciplinaId", middleware.IsAuthenticated, controllers.FecharSemestre)
	}

	{
		professor := api.Group("/professor")
		professor.POST("/", controllers.CadastrarProfessor)
		professor.POST("/login", controllers.Login)
		professor.DELETE("/:id", controllers.RemoverProfessor)
	}
}
