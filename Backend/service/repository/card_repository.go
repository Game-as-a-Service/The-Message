package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Id    int `gorm:"primaryKey;auto_increment"`
	Name  string
	Color string
	//IsDrawn   int
	//GameId    int       `gorm:"foreignKey:GameId;references:Id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

type CardRepository interface {
	GetCardById(ctx context.Context, id int) (*Card, error)
	CreateCard(ctx context.Context, card *Card) (*Card, error)
}
