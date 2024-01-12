package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

var ErrNoRecord = errors.New("no matching record found")

//fungsi memanggil movie berdasarkan id

func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at FROM movies WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)
	var movie Movie
	var releaseDateStr string
	var createdAtStr, updatedAtStr string

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&releaseDateStr,
		&movie.Runtime,
		&movie.Rating,
		&movie.MPAARating,
		&createdAtStr,
		&updatedAtStr,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, fmt.Errorf("error scanning database row: %w", err)
	}

	query = `SELECT mg.id, mg.movie_id, mg.genre_id, g.genre_name FROM movies_genres mg LEFT JOIN genres g on (g.id = mg.genre_id) WHERE mg.movie_id = ?`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	genres := make(map[int]string)

	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.GenreID] = mg.Genre.GenreName
	}

	movie.ReleaseDate, err = time.Parse("2006-01-02", releaseDateStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing release date: %w", err)
	}
	movie.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing created at date: %w", err)
	}
	movie.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing updated at date: %w", err)
	}
	movie.MovieGenre = genres
	return &movie, nil
}

//end section

// fungsi memanggil movie semuanya tanpa membatasi id
func (m *DBModel) All() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at FROM movies ORDER BY title`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		var releaseDateStr, createdAtStr, updatedAtStr string

		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&releaseDateStr,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&createdAtStr,
			&updatedAtStr,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning database row: %w", err)
		}

		movie.ReleaseDate, err = time.Parse("2006-01-02", releaseDateStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing release date: %w", err)
		}
		movie.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing created at date: %w", err)
		}
		movie.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing updated at date: %w", err)
		}

		genreQuery := `SELECT mg.id, mg.movie_id, mg.genre_id, g.genre_name FROM movies_genres mg LEFT JOIN genres g on (g.id = mg.genre_id) WHERE mg.movie_id = ?`
		genreRows, err := m.DB.QueryContext(ctx, genreQuery, movie.ID)
		if err != nil {
			return nil, fmt.Errorf("error querying database: %w", err)
		}

		genres := make(map[int]string)
		for genreRows.Next() {
			var mg MovieGenre
			err := genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)
			if err != nil {
				genreRows.Close()
				return nil, err
			}
			genres[mg.GenreID] = mg.Genre.GenreName
		}
		genreRows.Close()

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}

//end section
