package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds th application config
type AppConfig struct {
	UeCache       bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
