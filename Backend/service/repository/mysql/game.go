package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Id        int `gorm:"primaryKey;auto_increment"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{
		db: db,
	}
}

func (p *GameRepository) GetGameById(ctx context.Context, id int) (*Game, error) {
	game := &Game{}

	result := p.db.Table("games").Select("id", "Name").First(game, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}

func (p *GameRepository) CreateGame(ctx context.Context, game *Game) (*Game, error) {

	result := p.db.Table("games").Create(game)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}
