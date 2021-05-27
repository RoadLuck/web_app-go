package main

import (
	"fmt"
	"net/http"

	"github.com/roadluck/web_app/pkg/handlers"
)

var portNumber = ":8080"

// main es la raiz de la aplicacion
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Servidot en el puerto:", portNumber)
	http.ListenAndServe(portNumber, nil)

}
