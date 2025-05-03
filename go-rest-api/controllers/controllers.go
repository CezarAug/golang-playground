package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"study.co.jp/go-rest-api/database"
	"study.co.jp/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {
	var persons []models.Person

	database.DB.Find(&persons)
	json.NewEncoder(w).Encode(persons)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {

	// id := r.PathValue("id") (Plain go)
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person
	database.DB.First(&person, id)

	json.NewEncoder(w).Encode(person)

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var newPerson models.Person

	json.NewDecoder(r.Body).Decode(&newPerson)
	newPerson.Id = 0 //Avoiding any sent ID to be propagated to the DB.
	database.DB.Create(&newPerson)
	// Just a response
	json.NewEncoder(w).Encode(newPerson)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person
	database.DB.Delete(&person, id)

	// Create a better response, maybe just a status code
	json.NewEncoder(w).Encode(person)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person
	database.DB.First(&person, id)

	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Save(&person)

	json.NewEncoder(w).Encode(person)
}
