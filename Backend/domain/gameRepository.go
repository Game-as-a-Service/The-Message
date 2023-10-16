package domain

import "context"

type Game struct {
	Id      int `gorm:"primaryKey;auto_increment"`
	Token   string
	Players []Player
}

type GameRepository interface {
	GetGameById(ctx context.Context, id int) (*Game, error)
	CreateGame(ctx context.Context, game *Game) (*Game, error)
	DeleteGame(ctx context.Context, id int) error
	GetGameWithPlayers(ctx context.Context, id int) (*Game, error)
}
