package config

import "html/template"

// AppConfig holds the application configg
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
