package controllers

import "github.com/gin-gonic/gin"

func ExibeAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "teste",
	})
}
