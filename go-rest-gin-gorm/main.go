package main

import (
	"study.co.jp/go-rest-gin-gorm/database"
	"study.co.jp/go-rest-gin-gorm/routes"
)

func main() {

	database.Connect()
	routes.HandleRequests()
}
