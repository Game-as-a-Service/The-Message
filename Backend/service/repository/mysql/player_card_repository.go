package mysql

import (
	"context"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

type PlayerCardRepository struct {
	db *gorm.DB
}

func NewPlayerCardRepository(db *gorm.DB) *PlayerCardRepository {
	return &PlayerCardRepository{
		db: db,
	}
}

func (p PlayerCardRepository) GetPlayerCardById(ctx context.Context, id int) (*repository.PlayerCard, error) {
	card := new(repository.PlayerCard)

	result := p.db.First(&card, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return card, nil
}

func (p PlayerCardRepository) GetPlayerCardsByGameId(ctx context.Context, id int) ([]*repository.PlayerCard, error) {
	var cards []*repository.PlayerCard

	result := p.db.Find(&cards, "game_id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return cards, nil
}

func (p PlayerCardRepository) CreatePlayerCard(ctx context.Context, card *repository.PlayerCard) (*repository.PlayerCard, error) {
	result := p.db.Create(&card)

	if result.Error != nil {
		return nil, result.Error
	}

	return card, nil
}

func (p PlayerCardRepository) DeletePlayerCard(ctx context.Context, id int) error {
	card := new(repository.PlayerCard)

	result := p.db.Delete(&card, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *PlayerCardRepository) GetPlayerCardsByPlayerId(ctx context.Context, id int, gameId int, cardType string, cardId int) ([]*repository.PlayerCard, error) {

	var cards []*repository.PlayerCard
	result := p.db.Preload("Card").Find(&cards, "player_id = ? and type = ? and game_id = ? and card_id = ?", id, cardType, gameId, cardId)

	if result.Error != nil {
		return nil, result.Error
	}

	return cards, nil
}
