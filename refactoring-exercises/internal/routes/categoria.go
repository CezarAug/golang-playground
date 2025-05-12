package routes

import (
	"myapi/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupCategoriaRoutes(r *mux.Router) {
	// Endpoints para Categorias
	const baseCategoriasApi = "/api/categorias"
	r.HandleFunc(baseCategoriasApi, handlers.ListCategorias).Methods(http.MethodGet)             // GET para listar todas as categorias
	r.HandleFunc(baseCategoriasApi+"/{id}", handlers.GetCategoria).Methods(http.MethodGet)       // GET para buscar uma categoria (espera id via query)
	r.HandleFunc(baseCategoriasApi, handlers.CreateCategoria).Methods(http.MethodPost)           // POST para criar uma categoria
	r.HandleFunc(baseCategoriasApi, handlers.UpdateCategoria).Methods(http.MethodPut)            // PUT para atualizar uma categoria (JSON com id)
	r.HandleFunc(baseCategoriasApi+"/{id}", handlers.DeleteCategoria).Methods(http.MethodDelete) // DELETE para deletar uma categoria (espera id via query)

}
