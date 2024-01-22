package e2e

import (
	"context"
	"encoding/json"
	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"
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

func (suite *IntegrationTestSuite) TestTransmitIntelligenceE2E() {
	api := "/api/v1/player/{player_id}/transmit-intelligence"
	game, _ := suite.gameServ.InitGame(context.TODO())

	// Fake player count random 1~3
	playerCount := rand.Intn(3) + 1

	// Fake players data
	var players []request.PlayerInfo
	for i := 0; i < playerCount; i++ {
		player := request.PlayerInfo{
			ID:   faker.UUIDDigit(),
			Name: faker.FirstName(),
		}
		players = append(players, player)
	}

	// Fake game data
	createGameRequest := request.CreateGameRequest{
		Players: players,
	}

	_ = suite.playerServ.InitPlayers(context.TODO(), game, createGameRequest)
	_ = suite.gameServ.InitDeck(context.TODO(), game)
	_ = suite.gameServ.DrawCardsForAllPlayers(context.TODO(), game)

	suite.T().Run("it can validate intelligence type", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1
		cardId := rand.Intn(playerCount)

		// Request only card id
		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, "Invalid intelligence type", resBody["message"])
	})

	suite.T().Run("it can validate card id", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1
		intelligenceType := rand.Intn(3) + 1

		// Request only intelligence type
		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{IntelligenceType: intelligenceType}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "Card not found", resBody["message"])
	})

	suite.T().Run("it can fail when player not found", func(t *testing.T) {
		playerId := math.MaxInt32
		cardId := rand.Intn(playerCount)
		intelligenceType := rand.Intn(3) + 1

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId, IntelligenceType: intelligenceType}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "Player not found", resBody["message"])
	})

	suite.T().Run("it can fail when intelligence type is not valid", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1
		cardId := rand.Intn(playerCount)
		intelligenceType := math.MaxInt32

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId, IntelligenceType: intelligenceType}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, "Invalid intelligence type", resBody["message"])

	})

	suite.T().Run("it can fail when player card not found", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1
		cardId := math.MaxInt32
		intelligenceType := rand.Intn(3) + 1

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId, IntelligenceType: intelligenceType}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "Card not found", resBody["message"])
	})

	suite.T().Run("it can fail when game is end", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1
		intelligenceType := rand.Intn(3) + 1

		// Get player's card
		cards, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), playerId)

		// Random card id
		num := rand.Intn(len(cards.PlayerCards))
		cardId := cards.PlayerCards[num].CardId

		// Set player to current player
		suite.gameServ.UpdateCurrentPlayer(context.TODO(), game, playerId)

		// Set game status to end
		suite.gameServ.UpdateStatus(context.TODO(), game, enums.GameEnd)

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId, IntelligenceType: intelligenceType}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "遊戲已結束", resBody["message"])

		// Recover game status to start
		suite.gameServ.UpdateStatus(context.TODO(), game, enums.GameStart)
	})

	suite.T().Run("it can fail when not player's turn", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1
		intelligenceType := rand.Intn(3) + 1

		// Get player's card
		cards, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), playerId)

		// Random card id
		num := rand.Intn(len(cards.PlayerCards))
		cardId := cards.PlayerCards[num].CardId

		// Set other player to current player
		suite.gameServ.UpdateCurrentPlayer(context.TODO(), game, playerId-1)

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId, IntelligenceType: intelligenceType}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "尚未輪到你出牌", resBody["message"])
	})

	suite.T().Run("it can success when valid card id and intelligence type", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1
		intelligenceType := rand.Intn(3) + 1

		// Get player's card
		cards, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), playerId)

		// Random card id
		num := rand.Intn(len(cards.PlayerCards))
		cardId := cards.PlayerCards[num].CardId

		// Set player to current player
		suite.gameServ.UpdateCurrentPlayer(context.TODO(), game, playerId)

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId, IntelligenceType: intelligenceType}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		msg := enums.ToString(intelligenceType) + " intelligence transmitted"
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, msg, resBody["message"])
		assert.Equal(t, true, resBody["result"])
	})
}

type PlayCardRequest struct {
	CardId           int `json:"card_id"`
	IntelligenceType int `json:"intelligence_type"`
}
