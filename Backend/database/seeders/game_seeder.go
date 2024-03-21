package seeders

import (
	"math/rand"

	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func SeederGameWithPlayers(db *gorm.DB) {
	// Get all cards
	var cards []*repository.Card
	_ = db.Find(&cards)

	// Fake game data
	game := &repository.Game{
		RoomID:  faker.UUIDDigit(),
		Status:  enums.ActionCardStage,
		Players: []repository.Player{},
	}

	// Fake player count random 1~3
	playerCount := rand.Intn(3) + 1

	// Fake players data
	for i := 0; i < playerCount; i++ {
		player := repository.Player{
			UserID:      faker.UUIDDigit(),
			Name:        faker.FirstName(),
			PlayerCards: []repository.PlayerCard{},
		}

		// Each player gets 3 cards
		for j := 0; j < 3; j++ {
			player.PlayerCards = append(player.PlayerCards, repository.PlayerCard{
				CardID: cards[i*3+j].ID,
				Type:   "hand",
			})
		}
		game.Players = append(game.Players, player)
	}

	_ = db.Create(&game)
}
