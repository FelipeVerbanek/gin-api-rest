package routes

import (
	"fmt"
	"gin-api-rest/controllers"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.POST("/api/alunos", controllers.CriaAluno)
	r.GET("/api/alunos", controllers.ListarAlunos)
	r.GET("/api/alunos/:id", controllers.BuscaAlunoID)
	r.GET("/api/alunos/cpf/:cpf", controllers.BuscaALunoCPF)
	r.DELETE("/api/alunos/:id", controllers.DeleteAluno)
	r.PATCH("/api/alunos/:id", controllers.EditarAluno)

	r.GET("/", controllers.ExibePaginaIndex)

	r.NoRoute(controllers.RouteNotFound)

	port := os.Getenv("POST")

	if strings.TrimSpace(port) == "" {
		port = "3000"
	}

	r.Run(fmt.Sprintf(":%s", port))
}
