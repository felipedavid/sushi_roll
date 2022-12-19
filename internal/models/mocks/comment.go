package mocks

import (
	"time"

	"github.com/felipedavid/sushi_roll/internal/models"
)

var mockComment = &models.Comment{
	Id:        1,
	Content:   "Best comment of all time!!!",
	CreatedAt: time.Now(),
}

type CommentModel struct{}

// Insert irá inserir um novo jogo no banco de dados
func (m *CommentModel) Insert(userID, gameID int64, release string) (int64, error) {
	return 2, nil
}

// Get irá retonar um jogo específico baseado no "id"
func (m *CommentModel) Get(id int64) (*models.Comment, error) {
	switch id {
	case 1:
		return mockComment, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *CommentModel) Delete(id int64) error {
	return nil
}
