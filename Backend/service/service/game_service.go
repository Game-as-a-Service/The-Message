package service

import (
	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"math/rand"
	"time"
)

func InitialDeck(gameId int, cards []*repository.Card) []*repository.Card {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func InitIdentityCards(count int) []string {
	identityCards := make([]string, count)

	if count == 3 {
		identityCards[0] = enums.UndercoverFront
		identityCards[1] = enums.MilitaryAgency
		identityCards[2] = enums.Bystander
	}
	identityCards = shuffle(identityCards)
	return identityCards
}

func shuffle(cards []string) []string {
	shuffledCards := make([]string, len(cards))
	for i, j := range rand.Perm(len(cards)) {
		shuffledCards[i] = cards[j]
	}
	return shuffledCards
}
