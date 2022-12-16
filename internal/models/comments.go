package models

import (
	"database/sql"
	"errors"
	"time"
)

// Comment representa uma linha da tabela "comments"
type Comment struct {
	Id        int64
	UserID    int64
	GameID    int64
	Content   string
	CreatedAt time.Time
}

type CommentModelInterface interface {
	Insert(userID, movieID, content string) (int64, error)
}

// CommentModel é um objeto que representa as ações que podem ser realizadas contra a tabela "comments"
type CommentModel struct {
	DB *sql.DB
}

// Insert irá inserir um novo comentário no banco de dados
func (m *CommentModel) Insert(userID, gameID, content string) (int64, error) {
	stmt := `INSERT INTO comments VALUES(DEFAULT, $1, $2, $3, DEFAULT)`

	res, err := m.DB.Exec(stmt, userID, gameID, content)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}

// get irá retonar um comentário específico baseado no "id"
func (m *CommentModel) Get(id int64) (*Comment, error) {
	stmt := `SELECT * FROM comments WHERE id = $1`

	row := m.DB.QueryRow(stmt, id)

	c := &Comment{}

	if err := row.Scan(&c.Id, &c.UserID, &c.Content, &c.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return c, nil
}

func (m *CommentModel) Delete(id int64) error {
	stmt := `DELETE FROM comments WHERE id = $1`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		// TODO: Return persionalized error if comment does not exist
		return err
	}

	return nil
}
