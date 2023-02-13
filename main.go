package main

import (
	database "gin-api-rest/db"
	"gin-api-rest/routes"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func main() {
	godotenv.Load()
	database.Connection()

	routes.HandleRequest()
}
