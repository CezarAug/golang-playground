package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"study.co.jp/go-rest-gin-gorm/models"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	// Autmatically creating the database
	DB.AutoMigrate(&models.Student{})
}
