package mysql

import (
	"context"
	"time"

	"github.com/Game-as-a-Service/The-Message/domain"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Id        int `gorm:"primaryKey;auto_increment"`
	Token     string
	Players   []Player
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) domain.GameRepository {
	return &GameRepository{
		db: db,
	}
}

func (p *GameRepository) GetGameById(ctx context.Context, id int) (*domain.Game, error) {
	game := new(domain.Game)

	result := p.db.Table("games").First(game, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}

func (p *GameRepository) CreateGame(ctx context.Context, game *domain.Game) (*domain.Game, error) {

	result := p.db.Table("games").Create(game)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}

func (p *GameRepository) DeleteGame(ctx context.Context, id int) error {
	game := new(domain.Game)

	result := p.db.Table("games").Delete(game, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (g *GameRepository) GetGameWithPlayers(ctx context.Context, id int) (*domain.Game, error) {
	var game domain.Game
	if err := g.db.Preload("Players").First(&game, id).Error; err != nil {
		return nil, err
	}
	return &game, nil
}
