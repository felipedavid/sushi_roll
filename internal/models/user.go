package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	Name      string
	Email     string
	Birth     time.Time
	CreatedAt time.Time
}

type UserModel struct {
	DB *sql.DB
}

// Insert insere um novo usuário na tabela users
func (m *UserModel) Insert(name, email string, birth time.Time) error {
	stmt := `INSERT INTO user VALUES(DEFAULT, $1, $2, $3)`

	res, err := m.DB.Exec(stmt, name, email, birth)
	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return err
}

// Authenticate irá checar se existe um usuário na tabela users com o seguinte email e password.
// Se sim, retorna o id do usuário
func (m *UserModel) Authenticate(email, password string) (int64, error) {
	return 0, nil
}

// Exists checa se um usuário com seguinte 'id' existe
func (m *UserModel) Exists(id int64) error {
	return nil
}
