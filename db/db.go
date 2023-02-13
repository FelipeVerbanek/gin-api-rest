package database

import (
	"fmt"
	"gin-api-rest/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	host := os.Getenv("HOST_DB")
	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	dbname := os.Getenv("NAME_DB")
	port := os.Getenv("PORT_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	fmt.Println(dsn)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
