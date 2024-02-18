package mysql

import (
	"context"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{
		db: db,
	}
}

func (g *GameRepository) GetGameById(ctx context.Context, id int) (*repository.Game, error) {
	game := new(repository.Game)

	result := g.db.First(&game, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}

func (g *GameRepository) CreateGame(ctx context.Context, game *repository.Game) (*repository.Game, error) {

	result := g.db.Create(&game)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}

func (g *GameRepository) DeleteGame(ctx context.Context, id int) error {
	game := new(repository.Game)

	result := g.db.Delete(&game, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (g *GameRepository) GetGameWithPlayers(ctx context.Context, id int) (*repository.Game, error) {
	var game repository.Game
	if err := g.db.Preload("Players").First(&game, id).Error; err != nil {
		return nil, err
	}
	return &game, nil
}

func (g *GameRepository) UpdateGame(ctx context.Context, game *repository.Game) error {
	result := g.db.Save(&game)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
