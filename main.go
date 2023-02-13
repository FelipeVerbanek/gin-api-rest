package main

import (
	"gin-api-rest/db"
	"gin-api-rest/models"
	"gin-api-rest/routes"
)

func main() {
	db.Connection()
	models.Alunos = []models.Aluno{
		{Nome: "Felipe", CPF: "1231231231", RG: "12312312312"},
		{Nome: "Ana", CPF: "1231235261", RG: "12312392312"},
	}

	routes.HandleRequest()
}
