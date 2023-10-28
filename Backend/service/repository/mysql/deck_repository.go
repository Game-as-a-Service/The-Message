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

func (d *DeckRepository) GetDeckById(ctx context.Context, id int) (*repository.Deck, error) {
	deck := new(repository.Deck)

	result := d.db.Table("decks").First(deck, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return deck, nil
}

func (d *DeckRepository) Get(ctx context.Context) ([]*repository.Deck, error) {
	var decks []*repository.Deck

	result := d.db.Table("decks").Find(&decks)

	if result.Error != nil {
		return nil, result.Error
	}

	return decks, nil
}

func (d *DeckRepository) CreateDeck(ctx context.Context, deck *repository.Deck) (*repository.Deck, error) {
	result := d.db.Table("decks").Create(deck)

	if result.Error != nil {
		return nil, result.Error
	}

	return deck, nil
}
