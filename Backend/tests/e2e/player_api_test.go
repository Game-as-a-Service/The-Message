package e2e

import (
	"context"
	"encoding/json"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/stretchr/testify/assert"
)

func (suite *IntegrationTestSuite) TestPlayCardE2E() {
	// given
	game, _ := suite.gameServ.InitGame(context.TODO())
	createGameRequest := request.CreateGameRequest{
		Players: []request.PlayerInfo{
			{ID: "6497f6f226b40d440b9a90cc", Name: "A"},
			{ID: "6498112b26b40d440b9a90ce", Name: "B"},
			{ID: "6499df157fed0c21a4fd0425", Name: "C"},
		},
	}
	_ = suite.playerServ.InitPlayers(context.TODO(), game, createGameRequest)
	_ = suite.gameServ.InitDeck(context.TODO(), game)
	_ = suite.gameServ.DrawCardsForAllPlayers(context.TODO(), game)
	cards, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), 1)
	playCardId := cards.PlayerCards[0].CardId

	// when
	api := "/api/v1/players/1/player-cards"
	playCardRequest := PlayCardRequest{CardId: playCardId}
	jsonBody, err := json.Marshal(playCardRequest)
	if err != nil {
		suite.T().Fatalf("Failed to marshal JSON: %v", err)
	}
	resp := suite.requestJson(api, jsonBody)

	// then
	assert.Equal(suite.T(), 200, resp.StatusCode)
	assert.Equal(suite.T(), 200, resp.StatusCode)
	cards, _ = suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), 1)
	assert.Equal(suite.T(), 2, len(cards.PlayerCards))
}

type PlayCardRequest struct {
	CardId int `json:"card_id"`
}
