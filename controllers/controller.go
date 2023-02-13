package controllers

import (
	"gin-api-rest/models"

	"github.com/gin-gonic/gin"
)

func ExibeAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func NotFound(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(404, gin.H{
		"message": "Pagina " + nome + " não encontrada!",
	})
}
