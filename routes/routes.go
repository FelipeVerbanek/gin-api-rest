package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.POST("/alunos", controllers.CriaAluno)

	r.GET("/:nome", controllers.NotFound)

	r.Run(":3000")
}
