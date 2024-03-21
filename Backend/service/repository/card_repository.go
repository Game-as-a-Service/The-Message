package repository

import (
	"context"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Name             string
	Color            string
	IntelligenceType int
}

type CardRepository interface {
	GetCardById(ctx context.Context, id uint) (*Card, error)
	CreateCard(ctx context.Context, card *Card) (*Card, error)
	GetCards(ctx context.Context) ([]*Card, error)
	GetPlayerCardsByPlayerIdWithGameId(ctx context.Context, playerId uint) ([]*Card, error)
}
