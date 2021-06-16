package handlers

import (
	"net/http"

	"github.com/shaniarnon/booking/pkg/config"
	"github.com/shaniarnon/booking/pkg/models"

	"github.com/shaniarnon/booking/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for th handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page
func (m *Repository) Home(w http.ResponseWriter, req *http.Request) {
	remoteIP := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//About is the about page
func (m *Repository) About(w http.ResponseWriter, req *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
