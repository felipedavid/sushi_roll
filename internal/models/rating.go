package models

import (
	"database/sql"
	"errors"
	"time"
)

type Rating struct {
	Id        int64
	value      int
	CreatedAt time.Time
}

type RatingModel struct {
	DB *sql.DB
}

func (m *RatingModel) Insert(value int) (int64, error) {
	stmt := `INSERT INTO rating VALUES(DEFAULT, $1)`

	res, err := m.DB.Exec(stmt, value)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}

func (m *RaitngModel) Get(id int64) (*Rating, error) {
	stmt := `SELECT id, value, created_at FROM rating WHERE id = $1`

	row := m.DB.QueryRow(stmt, id)

	r := &Rating{}

	if err := row.Scan(&r.Id, &r.value, &r.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return r, nil
}

func (m *RatingModel) Delete(id int64) error {
	stmt := `DELETE FROM rating WHERE id = $1`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
