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

func (p *CardService) GetPlayerCardsByPlayerId(c context.Context, id int) ([]*repository.Card, error) {
	player, err := p.PlayerRepo.GetPlayer(c, id)
	game, err := p.GameRepo.GetGameWithPlayers(c, id)
	if err != nil {
		return nil, err
	}

	cards, err := p.CardRepo.GetPlayerCardsByPlayerId(c, player.Id, game.Id)

	if err != nil {
		return nil, err
	}
	return cards, nil
}
