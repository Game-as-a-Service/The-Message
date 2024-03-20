package repository

import (
	"context"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	GameID       uint `gorm:"foreignKey:GameID;references:ID"`
	UserID       string
	Name         string
	IdentityCard string
	Status       string
	Priority     int
	PlayerCards  []PlayerCard
	Game         *Game `gorm:"foreignKey:GameID;references:ID"`
}

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, player *Player) (*Player, error)
	GetPlayerById(ctx context.Context, playerId uint) (*Player, error)
	GetPlayersByGameId(ctx context.Context, id uint) ([]*Player, error)
	GetPlayerWithPlayerCards(ctx context.Context, playerId uint) (*Player, error)
	GetPlayerWithGame(ctx context.Context, playerId uint) (*Player, error)
	GetPlayerWithGamePlayersAndPlayerCardsCard(ctx context.Context, playerId uint) (*Player, error)
}
