package service

import (
	"math/rand"

	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
)

type GameService struct {
	GameRepo   repository.GameRepository
	PlayerRepo repository.PlayerRepository
}

func InitIdentityCards(count int) []string {
	identityCards := make([]string, count)
	// if count == 3 had 1 UndercoverFront, 1 MilitaryIntel, 1 Bystander
	if count == 3 {
		identityCards[0] = enums.UndercoverFront
		identityCards[1] = enums.MilitaryIntel
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
