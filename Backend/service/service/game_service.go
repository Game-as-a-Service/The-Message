package service

import (
	"context"
	"fmt"
	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"math/rand"
	"time"
)

func InitialDeck(gameId int, cards []*repository.Card, deckRepo repository.DeckRepository) {
	// 初始化隨機數生成器
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 隨機排序
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	dd(cards)

	for _, card := range cards {
		deck := new(repository.Deck)
		deck.GameId = gameId
		deck.CardId = card.Id
		_, err := deckRepo.CreateDeck(context.TODO(), deck)
		if err != nil {
			return
		}
	}
}

func dd(x interface{}) {
	fmt.Printf("%+v\n", x)
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
