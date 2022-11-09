package models

import (
	"database/sql"
	"errors"
)

//	Category representa uma linha da tabela "category"
type Category struct {
	Id       int64
	Title    string
	AgeGroup int
}

//	CategoryModel é um objeto que representa as ações que podem ser realizadas contra a tabela "category"
type CategoryModel struct {
	DB *sql.DB
}

//	Insert irá inserir uma nova categoria no banco de dados
func (m *CategoryModel) Insert(title string) (int64, error) {
	stmt := `INSERT INTO game VALUES(DEFAULT, $1, $2, $3)`

	res, err := m.DB.Exec(stmt, title)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}

// Get irá retonar uma categoria baseado no "id"
func (m *CategoryModel) Get(id int64) (*Category, error) {
	stmt := `SELECT id, title, agegroup FROM category WHERE id = $1`

	row := m.DB.QueryRow(stmt, id)

	c := &Category{}

	if err := row.Scan(&c.Id, &c.Title, &c.AgeGroup); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return c, nil
}

func (m *CategoryModel) Delete(id int64) error {
	stmt := `DELETE FROM category WHERE id = $1`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		// TODO: Return persionalized error if game does not exist
		return err
	}

	return nil
}
