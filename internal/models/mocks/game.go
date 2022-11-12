package mocks

import (
	"github.com/felipedavid/sushi_roll/internal/models"
	"time"
)

var mockGame = &models.Game{
	Id:          1,
	Title:       "Call of Duty",
	Description: "Best game of all time!!!",
	Release:     time.Now(),
	CreatedAt:   time.Now(),
}

type GameModel struct{}

// Insert irá inserir um novo jogo no banco de dados
func (m *GameModel) Insert(title, description, release string) (int64, error) {
	return 2, nil
}

// Get irá retonar um jogo específico baseado no "id"
func (m *GameModel) Get(id int64) (*models.Game, error) {
	switch id {
	case 1:
		return mockGame, nil
	default:
		return nil, models.ErrNoRecord
	}
}

// Latest irá retonar os 10 últimos jogos inseridos
// IMPORTANTE: Essa função será substituida futuramente por algo mais dinâmico e útil,
// então deve ser usada apenas para prototipação de funcionalidades
func (m *GameModel) Latest() ([]*models.Game, error) {
	return []*models.Game{mockGame}, nil
}

func (m *GameModel) Delete(id int64) error {
	return nil
}
