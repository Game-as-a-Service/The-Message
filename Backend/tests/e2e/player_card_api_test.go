package e2e

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/stretchr/testify/assert"
)

func (suite *IntegrationTestSuite) TestGetPlayerCards() {

	// given
	game := repository.Game{}
	_, err := suite.gameRepo.CreateGame(context.TODO(), &game)
	player := repository.Player{
		Name:         "player1",
		GameId:       1,
		IdentityCard: "醬油",
	}
	_, err = suite.playerRepo.CreatePlayer(context.TODO(), &player)
	if err != nil {
		panic(err)
	}

	_, err = suite.playerCardRepo.CreatePlayerCard(context.TODO(), &repository.PlayerCard{
		PlayerId: 1,
		GameId:   1,
		CardId:   1,
		Type:     "hand",
	})
	if err != nil {
		panic(err)
	}

	// when
	api := "/api/v1/player/1/player-cards/"
	resp := suite.requestJson(api, nil, http.MethodGet)
	response := suite.responseTest(resp)

	// then
	assert.Equal(suite.T(), 200, resp.StatusCode)

	jsonStr1 := `{
		"player_cards": [
			{
				"color": "",
				"id": 1,
				"name": ""
			}
		]
	}`

	playerCard := map[string]interface{}{
		"color": "",
		"id":    1,
		"name":  "",
	}

	playerCards := map[string]interface{}{
		"player_cards": []interface{}{playerCard},
	}

	err = json.Unmarshal([]byte(jsonStr1), &playerCards)
	if err != nil {
		panic(err)
	}

	assert.Equal(suite.T(), response, playerCards)
}

type PlayerCard struct {
	ID       string `json:"id"`
	PlayerID string `json:"player_id"`
	GameID   string `json:"game_id"`
	CardID   string `json:"card_id"`
	Type     string `json:"type"`
}

type Request_p struct {
	Player Player `json:"player"`
}
