package models

import (
	"database/sql"
	"errors"
	"time"
)

// Game representa uma linha da tabela "games"
type User struct {
	Id          int64
	Name      	string
	Email 		string
	Birth     	time.Time
	CreatedAt   time.Time
}

// GameModel é um objeto que representa as ações que podem ser realizadas contra a tabela "games"
type UserModel struct {
	DB *sql.DB
}

// Insert irá inserir um novo jogo no banco de dados
func (m *UserModel) Insert(name, email, birth string) (int64, error) {
	stmt := `INSERT INTO user VALUES(DEFAULT, $1, $2, $3)`

	res, err := m.DB.Exec(stmt, name, email, birth)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}

// Get irá retonar um jogo específico baseado no "id"
func (m *UserModel) Get(id int64) (*User, error) {
	stmt := `SELECT id, name, email, birth, created_at FROM user WHERE id = $1`

	row := m.DB.QueryRow(stmt, id)

	u := &User{}

	if err := row.Scan(&u.Id, &u.Name, &u.email, &u.Birth, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return u, nil
}

func (m *UserModel) Delete(id int64) error {
	stmt := `DELETE FROM user WHERE id = $1`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		// TODO: Return persionalized error if game does not exist
		return err
	}

	return nil
}
