package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PlayerCard struct {
	gorm.Model
	Id        int `gorm:"primaryKey;auto_increment"`
	PlayerId  int
	GameId    int
	CardId    int
	Type      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
	Card      Card   `gorm:"foreignKey:CardId"`
	Player    Player `gorm:"foreignKey:PlayerId"`
}

type PlayerCardRepository interface {
	GetPlayerCardById(ctx context.Context, id int) (*PlayerCard, error)
	GetPlayerCardsByGameId(ctx context.Context, id int) ([]*PlayerCard, error)
	CreatePlayerCard(ctx context.Context, card *PlayerCard) (*PlayerCard, error)
	DeletePlayerCard(ctx context.Context, id int) error
	DeletePlayerCardByPlayerIdAndCardId(ctx context.Context, playerId int, gameId int, cardId int) (bool, error)
	ExistPlayerCardByPlayerIdAndCardId(ctx context.Context, playerId int, gameId int, cardId int) (bool, error)
	GetPlayerCards(ctx context.Context, playerCard *PlayerCard) (*[]PlayerCard, error)
}
