package handlers

import (
	"net/http"

	"github.com/jericogantuangco/bookings/pkg/config"
	"github.com/jericogantuangco/bookings/pkg/models"
	"github.com/jericogantuangco/bookings/pkg/render"
)

var App *Application

type Application struct {
	Config *config.AppConfig
}

func NewRepo(cfg *config.AppConfig) *Application {
	return &Application{
		Config: cfg,
	}
}

func NewHandlers(app *Application) {
	App = app
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	app.Config.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", models.TemplateData{})
}

func (app *Application) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := app.Config.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", models.TemplateData{
		StringMap: stringMap,
	})
}
