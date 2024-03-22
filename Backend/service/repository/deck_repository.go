package repository

import (
	"context"

	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	GameID uint
	Cards  []DeckCard
	Game   Game `gorm:"foreignKey:GameID"`
}

type DeckCard struct {
	gorm.Model
	CardID uint
	DeckID uint
	Deck   Deck `gorm:"foreignKey:DeckID"`
	Card   Card `gorm:"foreignKey:CardID"`
}

type DeckRepository interface {
	CreateDeck(ctx context.Context, deck *Deck) (*Deck, error)
	GetDecksByGameId(ctx context.Context, id uint) ([]*Deck, error)
	DeleteDeck(ctx context.Context, id uint) error
}
