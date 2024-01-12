// nama package
package main

//Package
import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// inisialiasi routing
func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	//menambah routing movie
	router.HandlerFunc(http.MethodGet, "/movies/:id", app.getOneMovie) // Perbaikan path di sini
	router.HandlerFunc(http.MethodGet, "/movies", app.getAllMovies)
	//menambah routing middleware
	return app.enableCORS(router)
}
