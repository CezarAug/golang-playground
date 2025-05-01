package controllers

import (
	"net/http"
	"text/template"

	"study.co.jp/webstore/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	templates.ExecuteTemplate(w, "index", products)
}
