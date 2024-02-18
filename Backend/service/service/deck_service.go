package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/Game-as-a-Service/The-Message/service/repository"
)

type DeckService struct {
	CardService CardService
	DeckRepo    repository.DeckRepository
}

type DeckServiceOptions struct {
	CardService CardService
	DeckRepo    repository.DeckRepository
}

func NewDeckService(opts *DeckServiceOptions) DeckService {
	return DeckService{
		CardService: opts.CardService,
		DeckRepo:    opts.DeckRepo,
	}
}

func (d *DeckService) InitDeck(c context.Context, game *repository.Game) error {
	cards, err := d.CardService.GetCards(c)
	if err != nil {
		return err
	}

	cards = d.ShuffleDeck(cards)

	var deck []*repository.Deck

	for _, card := range cards {
		card, err := d.DeckRepo.CreateDeck(c, &repository.Deck{
			GameId: game.Id,
			CardId: card.Id,
		})
		if err != nil {
			return err
		}
		deck = append(deck, card)
	}

	return nil
}

func (d *DeckService) ShuffleDeck(cards []*repository.Card) []*repository.Card {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func (d *DeckService) CreateDeck(c context.Context, deck *repository.Deck) (*repository.Deck, error) {
	deck, err := d.DeckRepo.CreateDeck(c, deck)
	if err != nil {
		return nil, err
	}
	return deck, nil
}

func (d *DeckService) GetDecksByGameId(c context.Context, id int) ([]*repository.Deck, error) {
	decks, err := d.DeckRepo.GetDecksByGameId(c, id)
	if err != nil {
		return nil, err
	}
	return decks, nil

}

func (d *DeckService) DeleteDeckFromGame(c context.Context, id int) error {
	err := d.DeckRepo.DeleteDeck(c, id)
	if err != nil {
		return err
	}
	return nil
}
