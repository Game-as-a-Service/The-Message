package repository

import (
	"time"

	"gorm.io/gorm"
)

type playerCards struct {
	gorm.Model
	Player_id int `gorm:"primaryKey;auto_increment"`
	Type      string
	PlayerId  int       `gorm:"foreignKey:PlayerId;references:Id"`
	CardId    int       `gorm:"foreignKey:CardId;references:Id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

type playerCardsRepository interface {
	// InitialCard(ctx context.Context) (*Card, error)
}
