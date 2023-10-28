package mysql

import (
	"context"
	"fmt"

	"github.com/Game-as-a-Service/The-Message/enums"
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

func (c *CardRepository) InitialCard(ctx context.Context, gameId int) ([]*repository.Card, error) {

	cards := enums.GetCards(gameId)

	var err error

	for _, card := range cards {
		err := c.db.Table("cards").Create(&card).Error
		if err != nil {
			fmt.Println("Error creating card:", err)
		}
	}

	return cards, err
}
