package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"study.co.jp/webstore/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	//TODO: Is there a safer way than just writing post directly? (Like a Method enum)
	// This needs waaay more validations (e.g.: Negative values are allowed)
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := convertFloat(r.FormValue("price"), "price")
		quantity := convertInt(r.FormValue("quantity"), "quantity")

		models.CreateProduct(name, description, price, quantity)
	}

	http.Redirect(w, r, "/", 301)
}

func convertFloat(value string, attributeName string) float64 {
	convertedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Println("Failed to convert attribute", attributeName, err)
	}

	return convertedValue
}

func convertInt(value string, attributeName string) int {
	convertedValue, err := strconv.Atoi(value)
	if err != nil {
		log.Println("Failed to convert attribute", attributeName, err)
	}

	return convertedValue
}
