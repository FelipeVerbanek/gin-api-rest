package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.NotFound)

	r.Run(":3000")
}
