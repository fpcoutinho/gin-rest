package routes

import (
	"github.com/fpcoutinho/gin-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/api/alunos", controllers.All)
	r.GET("/api/alunos/curso/:curso", controllers.FindAlunosPorCurso)
	r.POST("/api/alunos", controllers.CriaAluno)
	r.GET("/api/alunos/:matricula", controllers.FindAluno)
	r.DELETE("/api/alunos/:matricula", controllers.DeleteAluno)
	r.PUT("/api/alunos/:matricula", controllers.UpdateAluno)
	r.Run("127.0.0.1:8080")
}
