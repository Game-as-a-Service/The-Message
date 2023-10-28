package mysql

import (
	"context"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

func (p *PlayerRepository) CreatePlayer(ctx context.Context, player *repository.Player) (*repository.Player, error) {
	err := p.db.Table("players").Create(player).Error
	return player, err
}

func (p *PlayerRepository) GetPlayer(ctx context.Context, playerId int) (*repository.Player, error) {
	player := new(repository.Player)

	result := p.db.Table("players").First(player, "id = ?", playerId)

	if result.Error != nil {
		return nil, result.Error
	}

	return player, nil
}

func (p *PlayerRepository) GetPlayerById(ctx context.Context, id int) (*repository.Player, error) {
	player := new(repository.Player)

	result := p.db.Table("players").First(player, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return player, nil
}

func (p *PlayerRepository) GetPlayerWithCardById(ctx context.Context, id int) (*repository.Player, error) {
	var player repository.Player
	if err := p.db.Preload("Cards").First(&player, id).Error; err != nil {
		return nil, err
	}
	return &player, nil
}
