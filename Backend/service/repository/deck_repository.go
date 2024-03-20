package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
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
	GetDecksByGameId(ctx context.Context, id int) ([]*Deck, error)
	DeleteDeck(ctx context.Context, id int) error
}
