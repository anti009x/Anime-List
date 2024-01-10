package main

import (
	"net/http"
	"server/models"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// create function movie
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
	movie := models.Movie{
		ID:          id,
		Title:       "Title",
		Description: "Deskripsi",
		Year:        2024,
		ReleaseDate: time.Date(2014, 12, 4, 0, 0, 0, 0, time.Local),
		Runtime:     120,
		Rating:      9,
		MPAARating:  "Good",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = app.writeJSON(rw, http.StatusOK, movie, "movie")
}

//end dummy data movie
