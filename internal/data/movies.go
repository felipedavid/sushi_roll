package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/felipedavid/sushi_roll/internal/validator"
	"github.com/lib/pq"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty,string"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"-"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contains more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate genres")
}

type MovieModel struct {
	DB *sql.DB
}

func (m *MovieModel) Insert(movie *Movie) error {
	stmt := `INSERT INTO movies (title, year, runtime, genres)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version`

	// Timeout queries that take longer then 3 seconds to run
	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelCtx()

	row := m.DB.QueryRowContext(ctx, stmt, movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres))
	return row.Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m *MovieModel) Get(id int64) (*Movie, error) {
	stmt := `SELECT id, created_at, title, year, runtime, genres, version FROM movies WHERE id = $1`

	// Timeout queries that take longer then 3 seconds to run
	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelCtx()

	var movie Movie
	err := m.DB.QueryRowContext(ctx, stmt, id).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		pq.Array(&movie.Genres),
		&movie.Version,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &movie, nil
}

func (m *MovieModel) Update(movie *Movie) error {
	stmt := `UPDATE movies 
		SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1 
		WHERE id = $5 AND version = $6
		RETURNING version`

	// Timeout queries that take longer then 3 seconds to run
	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelCtx()

	row := m.DB.QueryRowContext(ctx, stmt, movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres), movie.ID, movie.Version)
	if err := row.Scan(&movie.Version); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrEditConflict
		}
		return err
	}

	return nil
}

func (m *MovieModel) Delete(id int64) error {
	stmt := `DELETE FROM movies WHERE id = $1`

	// Timeout queries that take longer then 3 seconds to run
	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelCtx()

	result, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	if nRows, _ := result.RowsAffected(); nRows == 0 {
		return ErrRecordNotFound
	}

	return nil
}
