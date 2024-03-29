package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/BrunoPolaski/go-web-api/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SelectProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			panic(err.Error())
		}
		amount, err := strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			panic(err.Error())
		}
		models.InsertProduct(name, description, price, amount)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.SelectProductById(id)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			panic(err.Error())
		}
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			panic(err.Error())
		}
		amount, err := strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			panic(err.Error())
		}
		models.UpdateProduct(id, amount, name, description, price)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
