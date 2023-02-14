package main

import (
	database "gin-api-rest/db"
	"gin-api-rest/routes"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func main() {
	godotenv.Load()
	inputDb := loadEnvsDB()
	database.Connection(inputDb)

	routes.HandleRequest()
}

func loadEnvsDB() database.InputConnectDB {
	host := os.Getenv("HOST_DB")
	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	dbname := os.Getenv("NAME_DB")
	port := os.Getenv("PORT_DB")

	return database.InputConnectDB{
		Host:     host,
		User:     user,
		Password: password,
		Dbname:   dbname,
		Port:     port,
	}
}
