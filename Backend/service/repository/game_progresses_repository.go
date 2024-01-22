package repository

import (
	"context"
	"gorm.io/gorm"
	"time"
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
}
