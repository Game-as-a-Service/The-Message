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

//func (p *PlayerRepository) GetRandomCard(ctx context.Context, gameId int) (*repository.Card, error) {
//	card := new(repository.Card)
//
//	result := p.db.Order("RAND()").Where(&repository.Card{GameId: gameId}).First(card)
//
//	if result.Error != nil {
//		return nil, result.Error
//	}
//
//	return card, nil
//}

//func (p *PlayerRepository) PlayerDrawCard(ctx context.Context, playerId int) (*repository.PlayerCards, error) {
//
//	player, _ := p.GetPlayer(ctx, playerId)
//
//	card, _ := p.GetRandomCard(ctx, player.GameId)
//
//	playerCard := &repository.PlayerCards{
//		PlayerId: player.Id,
//		CardId:   card.Id,
//		Type:     "hands",
//	}
//
//	result := p.db.Table("PlayerCards").Create(&playerCard)
//
//	if result.Error != nil {
//		return nil, result.Error
//	}
//
//	return playerCard, nil
//}
