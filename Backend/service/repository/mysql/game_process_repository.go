package mysql

import (
	"context"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

type GameProgressRepository struct {
	db *gorm.DB
}

func NewGameProgressRepository(db *gorm.DB) *GameProgressRepository {
	return &GameProgressRepository{
		db: db,
	}
}

func (g *GameProgressRepository) CreateGameProgress(c context.Context, gameProgress *repository.GameProgresses) (*repository.GameProgresses, error) {
	result := g.db.Create(&gameProgress)

	if result.Error != nil {
		return nil, result.Error
	}

	return gameProgress, nil
}

func (g *GameProgressRepository) GetGameProgresses(c context.Context, targetPlayerId uint, gameId uint) (*repository.GameProgresses, error) {
	var gameProgress *repository.GameProgresses

	result := g.db.First(&gameProgress, "target_player_id = ? AND game_id = ?", targetPlayerId, gameId)
	if result.Error != nil {
		return nil, result.Error
	}

	return gameProgress, nil
}

func (g *GameProgressRepository) UpdateGameProgress(c context.Context, gameProgress *repository.GameProgresses, nextPlayerId uint) (*repository.GameProgresses, error) {
	result := g.db.First(&gameProgress)

	gameProgress.TargetPlayerID = nextPlayerId
	g.db.Save(&gameProgress)

	if result.Error != nil {
		return nil, result.Error
	}

	return gameProgress, nil
}
