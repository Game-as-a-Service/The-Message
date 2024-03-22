package repository

import (
	"context"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	RoomID          string
	Status          string
	CurrentPlayerID uint
	Players         []Player
	Deck            *Deck
}

type GameRepository interface {
	GetGameById(ctx context.Context, id uint) (*Game, error)
	CreateGame(ctx context.Context, game *Game) (*Game, error)
	DeleteGame(ctx context.Context, id uint) error
	GetGameWithPlayers(ctx context.Context, id uint) (*Game, error)
	UpdateGame(ctx context.Context, game *Game) error
	CreateGameWithPlayers(c context.Context, game *Game) (*Game, error)
}
