package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jericogantuangco/bookings/pkg/config"
	"github.com/jericogantuangco/bookings/pkg/handlers"
	"github.com/jericogantuangco/bookings/pkg/render"
)

const portNumber = ":8080"

var cfg config.AppConfig
var session *scs.SessionManager

func main() {

	cfg.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = cfg.InProduction

	cfg.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	cfg.TemplateCache = tc

	repo := handlers.NewRepo(&cfg)
	handlers.NewHandlers(repo)

	render.Newtemplates(&cfg)

	fmt.Printf("Starting server at port %s\n", portNumber)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&cfg),
	}

	err = server.ListenAndServe()
	log.Fatal(err)

}
