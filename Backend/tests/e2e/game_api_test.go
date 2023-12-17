package e2e

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func (suite *IntegrationTestSuite) TestStartGameE2E() {
	players := []Player{
		{ID: "6497f6f226b40d440b9a90cc", Name: "A"},
		{ID: "6498112b26b40d440b9a90ce", Name: "B"},
		{ID: "6499df157fed0c21a4fd0425", Name: "C"},
	}
	requestBody := StartGameRequest{Players: players}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		suite.T().Fatalf("Failed to marshal JSON: %v", err)
	}

	api := "/api/v1/games"
	resp := suite.requestJson(api, jsonBody, http.MethodPost)

	assert.Equal(suite.T(), 200, resp.StatusCode)

	responseJson := suite.responseJson(resp)

	assert.NotNil(suite.T(), responseJson["Token"], "JSON response should contain a 'Token' field")
	assert.NotNil(suite.T(), responseJson["Id"], "JSON response should contain a 'Id' field")

	// 驗證Game內的玩家都持有identity
	game, _ := suite.gameRepo.GetGameWithPlayers(context.TODO(), int(responseJson["Id"].(float64)))

	assert.NotEmpty(suite.T(), game.Players[0].IdentityCard)
	assert.NotEmpty(suite.T(), game.Players[1].IdentityCard)
	assert.NotEmpty(suite.T(), game.Players[2].IdentityCard)

	for _, player := range game.Players {
		playerCards, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), player.Id)
		assert.NotEmpty(suite.T(), playerCards.PlayerCards)
	}
}

type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StartGameRequest struct {
	Players []Player `json:"players"`
}
