package service

import (
	"context"
	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssignIdentityCards(t *testing.T) {
	gameService := NewGameService(&GameServiceOptions{})

	playCountMin := 3
	playCountMax := 9
	for playCount := playCountMin; playCount <= playCountMax; playCount++ {
		cards, _ := gameService.AssignIdentityCards(context.Background(), playCount)
		assert.Equal(t, playCount, len(cards))

		result := []string{}
		if playCount == 3 {
			result = []string{enums.UndercoverFront, enums.MilitaryAgency, enums.Bystander}
		} else if playCount == 4 {
			result = []string{enums.UndercoverFront, enums.MilitaryAgency, enums.Bystander, enums.Bystander}
		} else if playCount == 5 {
			result = []string{enums.UndercoverFront, enums.UndercoverFront, enums.MilitaryAgency, enums.MilitaryAgency, enums.Bystander}
		} else if playCount == 6 {
			result = []string{enums.UndercoverFront, enums.UndercoverFront, enums.MilitaryAgency, enums.MilitaryAgency, enums.Bystander, enums.Bystander}
		} else if playCount == 7 {
			result = []string{enums.UndercoverFront, enums.UndercoverFront, enums.UndercoverFront, enums.MilitaryAgency, enums.MilitaryAgency, enums.MilitaryAgency, enums.Bystander}
		} else if playCount == 8 {
			result = []string{enums.UndercoverFront, enums.UndercoverFront, enums.UndercoverFront, enums.MilitaryAgency, enums.MilitaryAgency, enums.MilitaryAgency, enums.Bystander, enums.Bystander}
		} else if playCount == 9 {
			result = []string{enums.UndercoverFront, enums.UndercoverFront, enums.UndercoverFront, enums.MilitaryAgency, enums.MilitaryAgency, enums.MilitaryAgency, enums.Bystander, enums.Bystander, enums.Bystander}
		}

		assert.ElementsMatch(t, result, cards)
	}
}
