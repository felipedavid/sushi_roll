package mocks

import (
	"github.com/felipedavid/sushi_roll/internal/models"
)

var mockCategory = &models.Category{
	Id:    1,
	Title: "Who cares?",
}

type CategoryModel struct{}

// Insert irá inserir um novo jogo no banco de dados
func (m *CategoryModel) Insert(title string) (int64, error) {
	return 2, nil
}

// Get irá retonar um jogo específico baseado no "id"
func (m *CategoryModel) Get(id int64) (*models.Category, error) {
	switch id {
	case 1:
		return mockCategory, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *CategoryModel) Delete(id int64) error {
	return nil
}
