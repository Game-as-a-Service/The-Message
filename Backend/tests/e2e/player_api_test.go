package e2e

import (
	"context"
	"encoding/json"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func (suite *IntegrationTestSuite) TestPlayCardE2E() {
	// Given
	api := "/api/v1/players/{player_id}/player-cards"
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

	playerId := rand.Intn(playerCount) + 1

	// Get player's card
	cards, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), playerId)

	// Random card id
	num := rand.Intn(len(cards.PlayerCards))
	cardId := cards.PlayerCards[num].CardId

	// Set player to current player
	suite.gameServ.UpdateCurrentPlayer(context.TODO(), game, playerId)

	url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
	req := PlayCardRequest{CardId: cardId}
	reqBody, _ := json.Marshal(req)

	res := suite.requestJson(url, reqBody, http.MethodPost)

	// Convert response body from json to map
	resBodyAsByteArray, _ := io.ReadAll(res.Body)
	resBody := make(map[string]interface{})
	_ = json.Unmarshal(resBodyAsByteArray, &resBody)

	// Then
	assert.Equal(suite.T(), 200, res.StatusCode)

	player, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), playerId)
	assert.Equal(suite.T(), len(cards.PlayerCards)-1, len(player.PlayerCards))
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

	suite.T().Run("it can validate card id", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1

		// Request only intelligence type
		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{}
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

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "Player not found", resBody["message"])
	})

	suite.T().Run("it can fail when player card not found", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1
		cardId := math.MaxInt32

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId}
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
		req := PlayCardRequest{CardId: cardId}
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

		// Get player's card
		cards, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), playerId)

		// Random card id
		num := rand.Intn(len(cards.PlayerCards))
		cardId := cards.PlayerCards[num].CardId

		// Set other player to current player
		suite.gameServ.UpdateCurrentPlayer(context.TODO(), game, playerId-1)

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "尚未輪到你出牌", resBody["message"])
	})

	suite.T().Run("it can success when valid card id", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1

		// Get player's card
		cards, _ := suite.playerRepo.GetPlayerWithPlayerCards(context.TODO(), playerId)

		// Random card id
		num := rand.Intn(len(cards.PlayerCards))
		cardId := cards.PlayerCards[num].CardId

		// Set player to current player
		suite.gameServ.UpdateCurrentPlayer(context.TODO(), game, playerId)

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := PlayCardRequest{CardId: cardId}
		reqBody, _ := json.Marshal(req)

		res := suite.requestJson(url, reqBody, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(res.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, resBody["result"])
	})
}

type PlayCardRequest struct {
	CardId int `json:"card_id"`
}
