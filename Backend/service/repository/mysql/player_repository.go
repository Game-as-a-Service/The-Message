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
	err := p.db.Create(&player).Error
	return player, err
}

func (p *PlayerRepository) GetPlayer(ctx context.Context, playerId int) (*repository.Player, error) {
	player := new(repository.Player)

	result := p.db.First(&player, "id = ?", playerId)

	if result.Error != nil {
		return nil, result.Error
	}

	return player, nil
}

func (p *PlayerRepository) GetPlayersByGameId(ctx context.Context, id int) ([]*repository.Player, error) {
	var players []*repository.Player

	result := p.db.Find(&players, "game_id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return players, nil
}

func (p *PlayerRepository) GetPlayerWithPlayerCards(ctx context.Context, playerId int) (*repository.Player, error) {
	var player repository.Player
	if err := p.db.Preload("PlayerCards").Preload("PlayerCards.Card").First(&player, playerId).Error; err != nil {
		return nil, err
	}
	return &player, nil
}
