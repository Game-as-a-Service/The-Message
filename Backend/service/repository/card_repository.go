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
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

type CardRepository interface {
	GetCardById(ctx context.Context, id int) (*Card, error)
	CreateCard(ctx context.Context, card *Card) (*Card, error)
	GetCards(ctx context.Context) ([]*Card, error)
}
