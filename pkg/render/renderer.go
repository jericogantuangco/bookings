package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/jericogantuangco/bookings/pkg/config"
	"github.com/jericogantuangco/bookings/pkg/models"
)

var functions = template.FuncMap{}
var configuration *config.AppConfig

func Newtemplates(cfg *config.AppConfig) {
	configuration = cfg
}

func AddDefaultData(td models.TemplateData) models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td models.TemplateData) {

	tc := configuration.TemplateCache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser")
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	if err := parsedTemplate.Execute(w, nil); err != nil {
		fmt.Println("Error parsing template", err)
		return
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tmplCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tmplCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tmplCache, err
			}
		}
		tmplCache[name] = ts
	}

	return tmplCache, nil

}
