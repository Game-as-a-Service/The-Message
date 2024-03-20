package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Id           int `gorm:"primaryKey;auto_increment"`
	GameId       uint `gorm:"foreignKey:GameId;references:ID"`
	UserId       string
	Name         string
	IdentityCard string
	Status       string
	Priority     int
	PlayerCards  []PlayerCard
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoCreateTime"`
	DeletedAt    gorm.DeletedAt
	Game         *Game `gorm:"foreignKey:GameId;references:ID"`
}

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, player *Player) (*Player, error)
	GetPlayerById(ctx context.Context, playerId int) (*Player, error)
	GetPlayersByGameId(ctx context.Context, id int) ([]*Player, error)
	GetPlayerWithPlayerCards(ctx context.Context, playerId int) (*Player, error)
	GetPlayerWithGame(ctx context.Context, playerId int) (*Player, error)
	GetPlayerWithGamePlayersAndPlayerCardsCard(ctx context.Context, playerId int) (*Player, error)
}
