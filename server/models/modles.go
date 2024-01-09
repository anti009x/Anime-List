// nama package
package models

//package
import (
	"time"
)

// create movie struktur
type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"releasedate"`
	Runtime     int          `json:"runtime"`
	Rating      int          `json:"rating"`
	MPAARating  string       `json:"mpaarating"`
	CreatedAt   time.Time    `json:"createdat"`
	UpdatedAt   time.Time    `json:"updatedat"`
	MovieGenre  []MovieGenre `json:"-"`
}

// create Genre struktur
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genrename"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

// create MovieGnere struktur
type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movieid"`
	GenreID   int       `json:"genreid"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
