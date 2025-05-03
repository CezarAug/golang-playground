package main

import (
	"fmt"

	"study.co.jp/go-rest-api/database"
	"study.co.jp/go-rest-api/routes"
)

func main() {
	database.DbConnect()

	fmt.Println("Started!!")
	routes.HandleRequest()
}
