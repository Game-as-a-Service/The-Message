package mysql

import (
	"context"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Id   int    `gorm:"primaryKey"`
	Name string `json:"name"`
}

type GameRepository struct {
	db *sql.DB
}

func NewGameRepositoryRepository(db *sql.DB) *GameRepository {
	return &GameRepository{
		db: db,
	}
}

func (p *GameRepository) GetPetById(ctx context.Context, id string) (*Game, error) {
	row := p.db.QueryRow("SELECT * FROM pet WHERE id = " + id)
	var pet Game
	err := row.Scan(&pet.Id, &pet.Name)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}