package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/roadluck/web_app/pkg/config"
	"github.com/roadluck/web_app/pkg/handlers"
	"github.com/roadluck/web_app/pkg/render"
)

var portNumber = ":8080"

// main es la raiz de la aplicacion
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("No se puede cargar el cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Servidot en el puerto:", portNumber)
	http.ListenAndServe(portNumber, nil)

}
