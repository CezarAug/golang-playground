package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	//TODO: This has to be better protected
	connnectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connnectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Error connecting to the database", DB)
	}
	log.Println("Database connected!")
}
