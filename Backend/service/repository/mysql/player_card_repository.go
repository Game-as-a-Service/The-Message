package mysql

import (
	"context"
	"errors"

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

func (p PlayerCardRepository) DeletePlayerCardByPlayerIdAndCardId(ctx context.Context, playerId int, gameId int, cardId int) (bool, error) {
	card := new(repository.PlayerCard)

	result := p.db.Delete(&card, "player_id = ? AND game_id = ? AND card_id = ?", playerId, gameId, cardId)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (p *PlayerCardRepository) ExistPlayerCardByPlayerIdAndCardId(ctx context.Context, playerId int, gameId int, cardId int) (bool, error) {
	var card repository.PlayerCard
	result := p.db.First(&card, "player_id = ? AND game_id = ? AND card_id = ?", playerId, gameId, cardId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (p PlayerCardRepository) GetPlayerCardById(ctx context.Context, id int) (*repository.PlayerCard, error) {
	card := new(repository.PlayerCard)

	result := p.db.First(&card, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return card, nil
}

func (p *PlayerCardRepository) GetPlayerCards(ctx context.Context, playerCard *repository.PlayerCard) (*[]repository.PlayerCard, error) {
	var playerCards *[]repository.PlayerCard
	result := p.db.Model(&repository.PlayerCard{}).Preload("Card").Where(&playerCard).Find(&playerCards)
	if result.Error != nil {
		return nil, result.Error
	}

	return playerCards, nil
}

func (p PlayerCardRepository) GetPlayerCardsByGameId(ctx context.Context, id int) ([]*repository.PlayerCard, error) {
	var cards []*repository.PlayerCard

	result := p.db.Find(&cards, "game_id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return cards, nil
}
