package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// get movie berdasarkan id pada database
func (app *application) getOneMovie(rw http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		// app.logger.Print(errors.New("Invalid Request"))

		app.errorJSON(rw, err)

		return

	}
	//end function movie

	// app.logger.Println("the id is", id)

	//dummy data movie

	movie, err := app.models.DB.Get(id)
	if err != nil {
		// Handle error here, misalnya kirim error JSON response
		app.errorJSON(rw, err)
		return
	}
	// Lanjutkan jika tidak ada error
	err = app.writeJSON(rw, http.StatusOK, movie, "movie")
}

//end section

// get semua movie dari database
func (app *application) getAllMovies(rw http.ResponseWriter, r *http.Request) {

	//dummy data movie
	movies, err := app.models.DB.All()
	if err != nil {
		// Handle error here, misalnya kirim error JSON response
		app.errorJSON(rw, err)
		return
	}
	// Lanjutkan jika tidak ada error
	err = app.writeJSON(rw, http.StatusOK, movies, "movies")
	if err != nil {
		// Handle error here, misalnya kirim error JSON response
		app.errorJSON(rw, err)
		return
	}
}

//end section
