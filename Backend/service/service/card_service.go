package service

import (
	"context"

	"github.com/Game-as-a-Service/The-Message/service/repository"
)

type CardService struct {
	CardRepo       repository.CardRepository
	GameRepo       repository.GameRepository
	PlayerRepo     repository.PlayerRepository
	PlayerCardRepo repository.PlayerCardRepository
}

type CardServiceOptions struct {
	CardRepo       repository.CardRepository
	GameRepo       repository.GameRepository
	PlayerRepo     repository.PlayerRepository
	PlayerCardRepo repository.PlayerCardRepository
}

func NewCardService(opts *CardServiceOptions) CardService {
	return CardService{
		CardRepo:       opts.CardRepo,
		GameRepo:       opts.GameRepo,
		PlayerRepo:     opts.PlayerRepo,
		PlayerCardRepo: opts.PlayerCardRepo,
	}
}

func (c *CardService) GetCards(ctx context.Context) ([]*repository.Card, error) {
	cards, err := c.CardRepo.GetCards(ctx)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (c *CardService) GetPlayerCardsByPlayerId(ctx context.Context, id int) ([]*repository.Card, error) {
	player, err := c.PlayerRepo.GetPlayerById(ctx, id)
	game, err := c.GameRepo.GetGameWithPlayers(ctx, player.GameId)
	if err != nil {
		return nil, err
	}

	cards, err := c.CardRepo.GetPlayerCardsByPlayerId(ctx, player.Id, game.Id)

	if err != nil {
		return nil, err
	}
	return cards, nil
}
