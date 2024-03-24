package repository

import (
	"context"

	"gorm.io/gorm"
)

type PlayerCard struct {
	gorm.Model
	PlayerID uint
	CardID   uint
	Type     string
	Card     Card   `gorm:"foreignKey:CardID"`
	Player   Player `gorm:"foreignKey:PlayerID"`
}

type PlayerCardRepository interface {
	GetPlayerCardById(ctx context.Context, id uint) (*PlayerCard, error)
	GetPlayerCardsByGameId(ctx context.Context, id uint) ([]*PlayerCard, error)
	CreatePlayerCard(ctx context.Context, card *PlayerCard) (*PlayerCard, error)
	DeletePlayerCard(ctx context.Context, id uint) error
	DeletePlayerCardByPlayerIdAndCardId(ctx context.Context, playerId uint, cardId uint) (bool, error)
	ExistPlayerCardByPlayerIdAndCardId(ctx context.Context, playerId uint, cardId uint) (bool, error)
	GetPlayerCards(ctx context.Context, playerCard *PlayerCard) (*[]PlayerCard, error)
}
