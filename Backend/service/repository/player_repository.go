package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Id           int `gorm:"primaryKey;auto_increment"`
	Name         string
	GameId       int `gorm:"foreignKey:GameId;references:Id"`
	IdentityCard string
	Status       string
	OrderNumber  int
	PlayerCards  []PlayerCard
	Game         *Game     `gorm:"foreignKey:GameId;references:Id"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoCreateTime"`
	DeletedAt    gorm.DeletedAt
}

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, player *Player) (*Player, error)
	GetPlayerById(ctx context.Context, playerId int) (*Player, error)
	GetPlayersByGameId(ctx context.Context, id int) ([]*Player, error)
	GetPlayerWithPlayerCards(ctx context.Context, playerId int) (*Player, error)
	GetPlayerWithGame(ctx context.Context, playerId int) (*Player, error)
	GetPlayerWithGamePlayersAndPlayerCardsCard(ctx context.Context, playerId int) (*Player, error)
}
