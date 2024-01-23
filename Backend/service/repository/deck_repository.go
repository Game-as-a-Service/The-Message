package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	Id        int `gorm:"primaryKey;auto_increment"`
	GameId    int
	CardId    int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

type DeckRepository interface {
	CreateDeck(ctx context.Context, deck *Deck) (*Deck, error)
	GetDecksByGameId(ctx context.Context, id int) ([]*Deck, error)
	DeleteDeck(ctx context.Context, id int) error
}
