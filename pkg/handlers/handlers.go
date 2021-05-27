package handlers

import (
	"net/http"

	"github.com/roadluck/web_app/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home_page.html")
}

// About es la pagina hacerca de nosostros
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about_page.html")
}
