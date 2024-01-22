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
