package mysql

import (
	"context"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

type CardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) *CardRepository {
	return &CardRepository{
		db: db,
	}
}

func (c *CardRepository) GetCardById(ctx context.Context, id uint) (*repository.Card, error) {
	card := new(repository.Card)

	result := c.db.First(&card, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return card, nil
}

func (c *CardRepository) CreateCard(ctx context.Context, card *repository.Card) (*repository.Card, error) {

	result := c.db.Create(&card)

	if result.Error != nil {
		return nil, result.Error
	}

	return card, nil
}

func (c *CardRepository) GetCards(ctx context.Context) ([]*repository.Card, error) {
	var cards []*repository.Card

	result := c.db.Find(&cards)

	if result.Error != nil {
		return nil, result.Error
	}

	return cards, nil
}

func (c *CardRepository) GetPlayerCardsByPlayerIdWithGameId(ctx context.Context, playerId uint) ([]*repository.Card, error) {

	var playerCards []*repository.PlayerCard
	var cards []*repository.Card
	result := c.db.Find(&playerCards, "player_id = ?", playerId)

	if result.Error != nil {
		return nil, result.Error
	}
	var cardIDs []uint
	for _, pc := range playerCards {
		cardIDs = append(cardIDs, pc.CardID)
	}

	result = c.db.Find(&cards, "ID IN ?", cardIDs)
	if result.Error != nil {
		return nil, result.Error
	}

	return cards, nil

}
