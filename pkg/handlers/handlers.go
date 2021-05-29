package handlers

import (
	"net/http"

	"github.com/roadluck/web_app/pkg/config"
	"github.com/roadluck/web_app/pkg/models"
	"github.com/roadluck/web_app/pkg/render"
)

var Repo *Repository

// Repository es el tipo repositorio
type Repository struct {
	App *config.AppConfig
}

//NewRepo crea un nuevo repositorio
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sete el repositorio de handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About es la pagina hacerca de nosostros
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//Logica
	stringMap := make(map[string]string)
	stringMap["testing"] = "Hola desde el render"

	//Enviar infotmacion al template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
