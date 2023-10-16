package mysql

import (
	"context"
	"time"

	"github.com/Game-as-a-Service/The-Message/domain"
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

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) domain.PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

func (p *PlayerRepository) CreatePlayer(ctx context.Context, player *domain.Player) (*domain.Player, error) {
	err := p.db.Table("players").Create(player).Error
	return player, err
}
