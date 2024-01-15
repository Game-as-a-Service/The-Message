package e2e

import (
	"context"
	"encoding/json"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
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
	game, _ = suite.gameServ.GetGameById(context.TODO(), game.Id)
	suite.gameServ.UpdateCurrentPlayer(context.TODO(), game, game.Players[0].Id)
	cards, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), 2)
	playCardId := cards.PlayerCards[0].CardId

	// when
	api := "/api/v1/players/" + strconv.Itoa(game.Players[0].Id) + "/player-cards"
	playCardRequest := PlayCardRequest{CardId: playCardId}
	jsonBody, err := json.Marshal(playCardRequest)
	if err != nil {
		suite.T().Fatalf("Failed to marshal JSON: %v", err)
	}
	resp := suite.requestJson(api, jsonBody, http.MethodPost)

	// then
	assert.Equal(suite.T(), 200, resp.StatusCode)
	play, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), 2)
	assert.Equal(suite.T(), 2, len(play.PlayerCards))
}

type PlayCardRequest struct {
	CardId int `json:"card_id"`
}
