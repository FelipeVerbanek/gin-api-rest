package database

import (
	"fmt"
	"gin-api-rest/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type InputConnectDB struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

func Connection(inputDB InputConnectDB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", inputDB.Host, inputDB.User, inputDB.Password, inputDB.Dbname, inputDB.Port)
	fmt.Println(dsn)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
