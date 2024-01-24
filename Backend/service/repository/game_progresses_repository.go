package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type GameProgresses struct {
	Id             int `gorm:"primaryKey;auto_increment"`
	PlayerId       int
	GameId         int
	CardId         int
	Action         string
	TargetPlayerId int
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoCreateTime"`
	DeletedAt      gorm.DeletedAt
}

type GameProgressesRepository interface {
	CreateGameProgress(c context.Context, gameProgress *GameProgresses) (*GameProgresses, error)
	GetGameProgresses(c context.Context, targetPlayerId int, gameId int) (*GameProgresses, error)
	UpdateGameProgress(c context.Context, gameProgress *GameProgresses, next_playerId int) (*GameProgresses, error)
}
