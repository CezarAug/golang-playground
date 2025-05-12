package routes

import (
	"myapi/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.JsonContentType)

	SetupItemRoutes(r)
	SetupCategoriaRoutes(r)

	return r
}
