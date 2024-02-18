package routes

import (
	"net/http"

	"github.com/BrunoPolaski/go-web-api/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}