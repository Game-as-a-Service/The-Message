package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Id              int `gorm:"primaryKey;auto_increment"`
	Token           string
	Status          string
	CurrentPlayerId int
	Players         []Player
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoCreateTime"`
	DeletedAt       gorm.DeletedAt
}

type GameRepository interface {
	GetGameById(ctx context.Context, id int) (*Game, error)
	CreateGame(ctx context.Context, game *Game) (*Game, error)
	DeleteGame(ctx context.Context, id int) error
	GetGameWithPlayers(ctx context.Context, id int) (*Game, error)
	UpdateGame(ctx context.Context, game *Game) error
}
