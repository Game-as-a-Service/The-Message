package repository

import (
	"context"

	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	GameId uint
	Cards  []DeckCard
	Game   Game `gorm:"foreignKey:GameId"`
}

type DeckCard struct {
	gorm.Model
	CardId uint
	DeckId uint
	Deck   Deck `gorm:"foreignKey:DeckId"`
	Card   Card `gorm:"foreignKey:CardId"`
}

type DeckRepository interface {
	CreateDeck(ctx context.Context, deck *Deck) (*Deck, error)
	GetDecksByGameId(ctx context.Context, id uint) ([]*Deck, error)
	DeleteDeck(ctx context.Context, id uint) error
}
