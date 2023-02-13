package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "teste",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/alunos", ExibeAlunos)

	r.Run(":3000")
}
