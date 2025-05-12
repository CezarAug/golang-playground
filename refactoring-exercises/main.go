package main

import (
	"log"
	"myapi/internal/config"
	"myapi/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDatabase()

	r := mux.NewRouter()

	// Endpoints para Itens
	r.HandleFunc("/api/itens", handlers.ListItens).Methods(http.MethodGet)                       // GET para listar todos os itens
	r.HandleFunc("/api/itens/{id}", handlers.GetItem).Methods(http.MethodGet)                    // GET para buscar um item (espera id via query: ?id=1)
	r.HandleFunc("/api/itens/codigo/{codigo}", handlers.GetItemByCodigo).Methods(http.MethodGet) // get-code?codigo=TEC001
	r.HandleFunc("/api/itens", handlers.CreateItem).Methods(http.MethodPost)                     // POST para criar um item
	r.HandleFunc("/api/itens", handlers.UpdateItem).Methods(http.MethodPut)                      // PUT para atualizar um item (JSON com id)
	r.HandleFunc("/api/itens", handlers.DeleteItem).Methods(http.MethodDelete)                   // DELETE para deletar um item (espera id via query: ?id=1)

	// Endpoints para Categorias
	const baseCategoriasApi = "/api/categorias"
	r.HandleFunc(baseCategoriasApi, handlers.ListCategorias).Methods(http.MethodGet)       // GET para listar todas as categorias
	r.HandleFunc(baseCategoriasApi+"/{id}", handlers.GetCategoria).Methods(http.MethodGet) // GET para buscar uma categoria (espera id via query)
	r.HandleFunc(baseCategoriasApi, handlers.CreateCategoria).Methods(http.MethodPost)     // POST para criar uma categoria
	r.HandleFunc(baseCategoriasApi, handlers.UpdateCategoria).Methods(http.MethodPut)      // PUT para atualizar uma categoria (JSON com id)
	r.HandleFunc(baseCategoriasApi, handlers.DeleteCategoria).Methods(http.MethodDelete)   // DELETE para deletar uma categoria (espera id via query)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
