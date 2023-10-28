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
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoCreateTime"`
	DeletedAt    gorm.DeletedAt
}

type PlayerCards struct {
	gorm.Model
	Id        int `gorm:"primaryKey;auto_increment"`
	Type      string
	PlayerId  int       `gorm:"foreignKey:PlayerId;references:Id"`
	CardId    int       `gorm:"foreignKey:CardId;references:Id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, player *Player) (*Player, error)
	GetPlayer(ctx context.Context, playerId int) (*Player, error)
	// GetRandomCard(ctx context.Context, gameId int) (*Card, error)
	// PlayerDrawCard(ctx context.Context, playerId int) (*PlayerCards, error)
}
