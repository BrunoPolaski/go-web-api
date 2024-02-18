package main

import (
	"net/http"

	"github.com/BrunoPolaski/go-web-api/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
