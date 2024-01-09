// nama package
package main

//Package
import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// inisialiasi routing
func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	return router

}
