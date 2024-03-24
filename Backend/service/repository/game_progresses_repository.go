package repository

import (
	"context"

	"gorm.io/gorm"
)

type GameProgresses struct {
	gorm.Model
	PlayerID       uint
	GameID         uint
	CardID         uint
	Action         string
	TargetPlayerID uint
}

type GameProgressesRepository interface {
	CreateGameProgress(c context.Context, gameProgress *GameProgresses) (*GameProgresses, error)
	GetGameProgresses(c context.Context, targetPlayerId uint, gameId uint) (*GameProgresses, error)
	UpdateGameProgress(c context.Context, gameProgress *GameProgresses, nextPlayerId uint) (*GameProgresses, error)
}
