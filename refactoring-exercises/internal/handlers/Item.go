package handlers

import (
	"encoding/json"
	"myapi/internal/config"
	"myapi/internal/models"
	"myapi/internal/repositories"
	"myapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ==================== HANDLERS PARA ITENS ====================

// Listar todos os itens
func ListItens(w http.ResponseWriter, r *http.Request) {

	repository := repositories.NewItemRepository()
	items, err := repository.ListAll()

	if err != nil {
		http.Error(w, "Error listing items", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

// Buscar um único item pelo id (via query string: ?id=1)
func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	item, err := repositories.NewItemRepository().FindById(id)
	if err != nil {
		http.Error(w, "Could not find item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

// Buscar um item pelo campo "codigo"
func GetItemByCodigo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cod := vars["codigo"]
	if cod == "" {
		http.Error(w, "Código não fornecido", http.StatusBadRequest)
		return
	}
	var item models.Item
	// Busca o item onde o campo "codigo" é igual ao valor fornecido
	if err := config.DB.Where("codigo = ?", cod).First(&item).Error; err != nil {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

// Criar um novo item (envie JSON via POST)
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Erro ao decodificar o item", http.StatusBadRequest)
		return
	}

	itemToSave, err := services.CreateItem(&item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := config.DB.Create(&itemToSave).Error; err != nil {
		http.Error(w, "Erro ao criar o item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(itemToSave)
}

// Atualizar um item (envie JSON via PUT, com o campo id preenchido)
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Erro ao decodificar o item", http.StatusBadRequest)
		return
	}
	if err := config.DB.Save(&item).Error; err != nil {
		http.Error(w, "Erro ao atualizar o item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

// Deletar um item (via query string: ?id=1)
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	if err := config.DB.Delete(&models.Item{}, id).Error; err != nil {
		http.Error(w, "Erro ao deletar o item", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Item deletado com sucesso"))
}
