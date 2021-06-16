package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/shaniarnon/booking/pkg/render"

	"github.com/shaniarnon/booking/pkg/config"
	"github.com/shaniarnon/booking/pkg/handlers"
)

const portNumber = ":8090"

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UeCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/headers", headers)

	fmt.Println(fmt.Sprintf("The server is listening on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routs(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
