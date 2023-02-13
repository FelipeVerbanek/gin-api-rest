package main

import (
	"gin-api-rest/db"
	"gin-api-rest/routes"

	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func main() {
	db.Connection()

	routes.HandleRequest()
}
