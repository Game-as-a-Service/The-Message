package service

import (
	"context"
	"github.com/Game-as-a-Service/The-Message/service/repository"
)

type CardService struct {
	CardRepo repository.CardRepository
}

type CardServiceOptions struct {
	CardRepo repository.CardRepository
}

func NewCardService(opts *CardServiceOptions) CardService {
	return CardService{
		CardRepo: opts.CardRepo,
	}
}

func (c *CardService) GetCards(ctx context.Context) ([]*repository.Card, error) {
	cards, err := c.CardRepo.GetCards(ctx)
	if err != nil {
		return nil, err
	}
	return cards, nil
}
