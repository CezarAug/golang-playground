package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"study.co.jp/go-rest-api/controllers"
	"study.co.jp/go-rest-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()

	// Middleware for setting all content types
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/persons", controllers.AllPersons).Methods(http.MethodGet)
	r.HandleFunc("/api/persons", controllers.CreatePerson).Methods(http.MethodPost)
	r.HandleFunc("/api/persons/{id}", controllers.GetPerson).Methods(http.MethodGet)
	r.HandleFunc("/api/persons/{id}", controllers.UpdatePerson).Methods(http.MethodPut)
	r.HandleFunc("/api/persons/{id}", controllers.DeletePerson).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe("localhost:8000", handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowCredentials())(r)))
}
