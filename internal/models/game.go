package models

import (
	"database/sql"
	"errors"
	"time"
)

// Game representa uma linha da tabela "games"
type Game struct {
	Id          int64
	Title       string
	Description string
	Release     time.Time
	CreatedAt   time.Time
}

// GameModel é um objeto que representa as ações que podem ser realizadas contra a tabela "games"
type GameModel struct {
	DB *sql.DB
}

// Insert irá inserir um novo jogo no banco de dados
func (m *GameModel) Insert(title, description, release string) (int64, error) {
	stmt := `INSERT INTO game VALUES(DEFAULT, $1, $2, $3)`

	res, err := m.DB.Exec(stmt, title, description, release)
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
func (m *GameModel) Get(id int64) (*Game, error) {
	stmt := `SELECT id, title, description, release_at, created_at FROM game WHERE id = $1`

	row := m.DB.QueryRow(stmt, id)

	g := &Game{}

	if err := row.Scan(&g.Id, &g.Title, &g.Description, &g.Release, &g.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return g, nil
}

// Latest irá retonar os 10 últimos jogos inseridos
// IMPORTANTE: Essa função será substituida futuramente por algo mais dinâmico e útil,
// então deve ser usada apenas para prototipação de funcionalidades
func (m *GameModel) Latest() ([]*Game, error) {
	stmt := `SELECT id, title, description, release_at, created_at FROM game LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	games := make([]*Game, 0, 10)

	for rows.Next() {
		g := &Game{}

		err = rows.Scan(&g.Id, &g.Title, &g.Description, &g.Release, &g.CreatedAt)
		if err != nil {
			return nil, err
		}

		games = append(games, g)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func (m *GameModel) Delete(id int64) error {
	stmt := `DELETE FROM game WHERE id = $1`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		// TODO: Return persionalized error if game does not exist
		return err
	}

	return nil
}
