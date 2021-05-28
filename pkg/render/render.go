package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/roadluck/web_app/pkg/config"
)

var functions = template.FuncMap{}

func NewTemplates(a *config.AppConfig)

// RenderTemplate renderiza los templates utilizando el paquete html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Toma el cache del template desde app config.

	tc, err := CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to bowser", err)
	}

}

//CreateTemplateCache Busca el cache y lo carga al template.
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
