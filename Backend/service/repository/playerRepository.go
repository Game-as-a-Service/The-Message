package repository

import (
	"context"
)

type Player struct {
	Id           int `gorm:"primaryKey;auto_increment"`
	Name         string
	GameId       int `gorm:"foreignKey:GameId;references:Id"`
	IdentityCard string
}

type PlayerRepository interface {
	CreatePlayer(cxt context.Context, player *Player) (*Player, error)
}
