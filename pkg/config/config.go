package config

import (
	"html/template"
	"log"
)

// AppConfig configuracin de mi app
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
