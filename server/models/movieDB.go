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
	movie.ReleaseDate, err = time.Parse("2006-01-02", releaseDateStr)
	movie.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	movie.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtStr)
	return &movie, nil
}
