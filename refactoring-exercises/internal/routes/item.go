package routes

import (
	"myapi/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupItemRoutes(r *mux.Router) {
	// Endpoints para Itens
	r.HandleFunc("/api/itens", handlers.ListItens).Methods(http.MethodGet)                       // GET para listar todos os itens
	r.HandleFunc("/api/itens/{id}", handlers.GetItem).Methods(http.MethodGet)                    // GET para buscar um item (espera id via query: ?id=1)
	r.HandleFunc("/api/itens/codigo/{codigo}", handlers.GetItemByCodigo).Methods(http.MethodGet) // get-code?codigo=TEC001
	r.HandleFunc("/api/itens", handlers.CreateItem).Methods(http.MethodPost)                     // POST para criar um item
	r.HandleFunc("/api/itens", handlers.UpdateItem).Methods(http.MethodPut)                      // PUT para atualizar um item (JSON com id)
	r.HandleFunc("/api/itens/{id}", handlers.DeleteItem).Methods(http.MethodDelete)              // DELETE para deletar um item (espera id via query: ?id=1)

}
