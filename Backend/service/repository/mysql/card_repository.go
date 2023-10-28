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

	result := c.db.Table("cards").First(card, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return card, nil
}

func (c *CardRepository) CreateCard(ctx context.Context, card *repository.Card) (*repository.Card, error) {

	result := c.db.Table("cards").Create(card)

	if result.Error != nil {
		return nil, result.Error
	}

	return card, nil
}

func (c *CardRepository) GetCards(ctx context.Context) ([]*repository.Card, error) {
	var cards []*repository.Card

	result := c.db.Table("cards").Find(&cards)

	if result.Error != nil {
		return nil, result.Error
	}

	return cards, nil
}
