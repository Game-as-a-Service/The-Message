package service

import (
	"context"
	"errors"
	"math/rand"
	"time"

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

	if len(cards) == 0 {
		return nil, errors.New("no cards found")
	}

	return cards, nil
}

func (c *CardService) GetPlayerCardsByPlayerId(ctx context.Context, id uint) ([]*repository.Card, error) {
	player, err := c.PlayerRepo.GetPlayerById(ctx, id)
	if err != nil {
		return nil, err
	}

	cards, err := c.CardRepo.GetPlayerCardsByPlayerIdWithGameId(ctx, player.ID)

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardService) ShuffleCards(ctx context.Context, cards []*repository.Card) []*repository.Card {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards
}
