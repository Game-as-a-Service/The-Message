package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Id        int `gorm:"primaryKey;auto_increment"`
	Name      string
	Color     string
	IsDrawed  int
	GameId    int       `gorm:"foreignKey:GameId;references:Id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

type CardRepository interface {
	InitialCard(ctx context.Context, gameId int) ([]*Card, error)
}
