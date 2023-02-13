package controllers

import (
	database "gin-api-rest/db"
	"gin-api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriaAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func NotFound(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(404, gin.H{
		"message": "Pagina " + nome + " não encontrada!",
	})
}
