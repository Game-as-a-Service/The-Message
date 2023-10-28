package mysql

import (
	"context"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

type DeckRepository struct {
	db *gorm.DB
}

func NewDeckRepository(db *gorm.DB) *DeckRepository {
	return &DeckRepository{
		db: db,
	}
}

func (d *DeckRepository) CreateDeck(ctx context.Context, deck *repository.Deck) (*repository.Deck, error) {
	result := d.db.Table("decks").Create(deck)

	if result.Error != nil {
		return nil, result.Error
	}

	return deck, nil
}

func (d *DeckRepository) GetDecksByGameId(ctx context.Context, id int) ([]*repository.Deck, error) {
	var decks []*repository.Deck

	result := d.db.Table("decks").Find(&decks, "game_id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return decks, nil
}

func (d *DeckRepository) DeleteDeck(ctx context.Context, id int) error {
	deck := new(repository.Deck)

	result := d.db.Table("decks").Delete(deck, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
