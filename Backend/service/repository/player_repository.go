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
	PlayerCards  []PlayerCard
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoCreateTime"`
	DeletedAt    gorm.DeletedAt
}

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, player *Player) (*Player, error)
	GetPlayer(ctx context.Context, playerId int) (*Player, error)
	GetPlayersByGameId(ctx context.Context, id int) ([]*Player, error)
	GetPlayerWithPlayerCards(ctx context.Context, playerId int) (*Player, error)
}
