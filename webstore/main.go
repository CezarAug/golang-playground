package main

import (
	"net/http"

	_ "github.com/lib/pq" //Runtime required import
	"study.co.jp/webstore/routes"
)

func main() {

	routes.LoadRoutes()
	//TODO: Set a configuration/environment for the port
	http.ListenAndServe("localhost:8000", nil)
}
