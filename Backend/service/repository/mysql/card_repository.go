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

func (c *CardRepository) GetCardById(ctx context.Context, id int) (*repository.Card, error) {
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

func (c *CardRepository) GetPlayerCardsByPlayerId(ctx context.Context, id int, gid int) ([]*repository.Card, error) {

	var playerCards []*repository.PlayerCard
	var cards []*repository.Card
	result := c.db.Find(&playerCards, "player_id = ? AND game_id = ?", id, gid)

	if result.Error != nil {
		return nil, result.Error
	}
	var cardIDs []int
	for _, pc := range playerCards {
		cardIDs = append(cardIDs, pc.CardId)
	}

	result = c.db.Find(&cards, "id IN ?", cardIDs)
	if result.Error != nil {
		return nil, result.Error
	}

	return cards, nil

}
