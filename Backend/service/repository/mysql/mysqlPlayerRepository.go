package mysql

import (
	"context"
	domain "github.com/Game-as-a-Service/The-Message/service/repository"
	repository "github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
	"time"
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

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) repository.PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

func (p *PlayerRepository) CreatePlayer(cxt context.Context, player *domain.Player) (*domain.Player, error) {
	err := p.db.Table("players").Create(player).Error
	return player, err
}
