package enums

import (
	"github.com/Game-as-a-Service/The-Message/service/repository"
)

func GetCards(gameId int) []*repository.Card {

	IsDrawed := 0

	cards := []*repository.Card{}
	cards = append(cards, &repository.Card{
		Id:       1,
		Name:     "Lock On",
		Color:    "Red",
		IsDrawed: IsDrawed,
		GameId:   gameId,
	}, &repository.Card{
		Id:       2,
		Name:     "Lock On",
		Color:    "Blue",
		IsDrawed: IsDrawed,
		GameId:   gameId,
	}, &repository.Card{
		Id:       3,
		Name:     "Lure Away",
		Color:    "Red",
		IsDrawed: IsDrawed,
		GameId:   gameId,
	}, &repository.Card{
		Id:       4,
		Name:     "Lure Away",
		Color:    "Red",
		IsDrawed: IsDrawed,
		GameId:   gameId,
	})

	return cards
}
