package config

import "html/template"

// AppConfig holds the application configg
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
