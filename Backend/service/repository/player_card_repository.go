package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PlayerCard struct {
	gorm.Model
	Id        int `gorm:"primaryKey;auto_increment"`
	PlayerId  int
	GameId    int
	CardId    int
	Type      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
	Card      Card
	Player    Player
}

type PlayerCardRepository interface {
	GetPlayerCardById(ctx context.Context, id int) (*PlayerCard, error)
	GetPlayerCardsByGameId(ctx context.Context, id int) ([]*PlayerCard, error)
	CreatePlayerCard(ctx context.Context, card *PlayerCard) (*PlayerCard, error)
	DeletePlayerCard(ctx context.Context, id int) error
	GetPlayerCardsByPlayerId(ctx context.Context, id int) ([]*PlayerCard, error)
}
