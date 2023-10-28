package repository

import (
	"context"
	"gorm.io/gorm"
	"time"
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
	GetDeckById(ctx context.Context, id int) (*Deck, error)
	Get(ctx context.Context) ([]*Deck, error)
	CreateDeck(ctx context.Context, deck *Deck) (*Deck, error)
}
