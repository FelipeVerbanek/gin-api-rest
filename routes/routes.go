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

	r.POST("/alunos", controllers.CriaAluno)

	r.GET("/:nome", controllers.NotFound)

	port := os.Getenv("POST")

	if strings.TrimSpace(port) == "" {
		port = "3000"
	}

	r.Run(fmt.Sprintf(":%s", port))
}
