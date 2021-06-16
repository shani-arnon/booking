package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole is a middleware to log message when page is hit
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf is a middleware to protect all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad saves the session and loads on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
