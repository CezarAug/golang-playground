package routes

import (
	"net/http"

	"study.co.jp/webstore/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
