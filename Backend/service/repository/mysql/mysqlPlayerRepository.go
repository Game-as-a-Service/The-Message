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

func (p *PlayerRepository) PlayerDrawCard(ctx context.Context, playerId int) (*repository.Player, error) {

	// Need to split.
	player := new(repository.Player)

	playerResult := p.db.Table("players").First(player, "id = ?", playerId)

	if playerResult.Error != nil {
		return nil, playerResult.Error
	}

	card := new(repository.Card)

	cardResult := p.db.Order("RAND()").Where(&repository.Card{GameId: player.GameId}).First(card)

	if cardResult.Error != nil {
		return nil, cardResult.Error
	}

	return player, cardResult.Error

	// err := p.db.Table("players").Create(player).Error
	// return player, err
}
