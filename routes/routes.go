package routes

import (
	"github.com/gin-gonic/gin"
	"sistema-alunos-go/controllers"
	middleware "sistema-alunos-go/middlewares"
)

// RegistraRotas inicializa todas as rotas disponíveis na API
func RegistraRotas(router *gin.Engine) {
	api := router.Group("")

	{
		aluno := api.Group("/aluno")
		aluno.POST("/", middleware.Autenticado, controllers.CadastrarAluno)
		aluno.GET("/desativar/:id", middleware.Autenticado, controllers.DesativarAluno)
		aluno.GET("/reativar/:id", middleware.Autenticado, controllers.ReativarAluno)
		aluno.DELETE("/:id", middleware.Autenticado, controllers.RemoverAluno)
	}

	{
		aula := api.Group("/aula")
		aula.POST("/:disciplinaId", middleware.Autenticado, controllers.CadastrarAula)
		aula.GET("/disciplina/:disciplinaId", middleware.Autenticado, controllers.ListarAulasDisciplina)
		aula.GET("/:id", middleware.Autenticado, controllers.GetAula)
	}

	{
		disciplina := api.Group("disciplina")
		disciplina.POST("/", middleware.Autenticado, controllers.CadastrarDisciplina)
		disciplina.POST("/matricular", middleware.Autenticado, controllers.MatricularAluno)
		disciplina.POST("/avaliacao/:disciplinaId", middleware.Autenticado, controllers.AdicionarAvaliacao)
		disciplina.POST("/avaliacao/:disciplinaId/nota/:avaliacaoId", middleware.Autenticado, controllers.AdicionarNotaAvaliacao)
		disciplina.GET("/", middleware.Autenticado, controllers.ListarDisciplinas)
		disciplina.GET("/fechar-semestre/:disciplinaId", middleware.Autenticado, controllers.FecharSemestre)
	}

	{
		professor := api.Group("/professor")
		professor.POST("/", controllers.CadastrarProfessor)
		professor.POST("/login", controllers.Login)
		professor.DELETE("/:id", controllers.RemoverProfessor)
	}
}
