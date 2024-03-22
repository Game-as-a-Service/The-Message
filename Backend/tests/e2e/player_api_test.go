package e2e

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/utils"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func (suite *IntegrationTestSuite) TestPlayCardE2E() {
	// Given
	api := "/api/v1/players/{player_id}/player-cards"

	// Get all cards
	var cards []*repository.Card
	_ = suite.db.Find(&cards)

	// Fake game data
	game := &repository.Game{
		RoomID:  faker.UUIDDigit(),
		Status:  enums.GameStart,
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

	_ = suite.db.Create(&game)

	id := rand.Intn(playerCount)
	player := game.Players[id]

	// Random card id
	num := rand.Intn(len(player.PlayerCards))
	cardId := player.PlayerCards[num].CardID

	// Set player to current player
	game.CurrentPlayerID = player.ID
	_ = suite.db.Save(&game)

	url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(int(player.ID)))

	req := utils.RequestToJsonBody(PlayCardRequest{CardID: cardId})

	res := suite.requestJson(url, req, http.MethodPost)

	// Then
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	var count int64
	playerCard := &repository.PlayerCard{}
	_ = suite.db.Model(&playerCard).Where("player_id = ?", player.ID).Count(&count)

	PlayerCards := int64(len(player.PlayerCards) - 1)
	assert.Equal(suite.T(), PlayerCards, count)
}

func (suite *IntegrationTestSuite) TestTransmitIntelligenceE2E() {
	api := "/api/v1/player/{player_id}/transmit-intelligence"

	// Get all cards
	var cards []*repository.Card
	_ = suite.db.Find(&cards)

	// Fake game data
	game := &repository.Game{
		RoomID:  faker.UUIDDigit(),
		Status:  enums.GameStart,
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

	_ = suite.db.Create(&game)

	suite.T().Run("it can validate card id", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1

		// Request only intelligence type
		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := utils.RequestToJsonBody(PlayCardRequest{})

		res := suite.requestJson(url, req, http.MethodPost)

		// Convert response body from json to map
		resBody := utils.JsonBodyToMap(res)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "Card not found", resBody["message"])
	})

	suite.T().Run("it can fail when player not found", func(t *testing.T) {
		playerId := math.MaxInt32
		cardId := uint(rand.Intn(playerCount))

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := utils.RequestToJsonBody(PlayCardRequest{CardID: cardId})

		res := suite.requestJson(url, req, http.MethodPost)

		// Convert response body from json to map
		resBody := utils.JsonBodyToMap(res)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "Player not found", resBody["message"])
	})

	suite.T().Run("it can fail when player card not found", func(t *testing.T) {
		playerId := rand.Intn(playerCount) + 1
		cardId := uint(math.MaxInt32)

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(playerId))
		req := utils.RequestToJsonBody(PlayCardRequest{CardID: cardId})

		res := suite.requestJson(url, req, http.MethodPost)

		// Convert response body from json to map
		resBody := utils.JsonBodyToMap(res)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "Card not found", resBody["message"])
	})

	suite.T().Run("it can fail when game is end", func(t *testing.T) {
		id := rand.Intn(playerCount)
		player := game.Players[id]

		// Random card id
		num := rand.Intn(len(player.PlayerCards))
		cardId := player.PlayerCards[num].CardID

		// Set player to current player
		game.CurrentPlayerID = player.ID
		game.Status = enums.GameEnd
		_ = suite.db.Save(&game)

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(int(player.ID)))

		req := utils.RequestToJsonBody(PlayCardRequest{CardID: cardId})

		res := suite.requestJson(url, req, http.MethodPost)

		// Convert response body from json to map
		resBody := utils.JsonBodyToMap(res)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "遊戲已結束", resBody["message"])

		// Recover game status to start
		game.Status = enums.GameStart
		_ = suite.db.Save(&game)
	})

	suite.T().Run("it can fail when not player's turn", func(t *testing.T) {
		id := rand.Intn(playerCount)
		player := game.Players[id]

		// Random card id
		num := rand.Intn(len(player.PlayerCards))
		cardId := player.PlayerCards[num].CardID

		// Set other player to current player
		game.CurrentPlayerID = player.ID + 1
		_ = suite.db.Save(&game)

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(int(player.ID)))
		req := utils.RequestToJsonBody(PlayCardRequest{CardID: cardId})

		res := suite.requestJson(url, req, http.MethodPost)

		// Convert response body from json to map
		resBody := utils.JsonBodyToMap(res)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, "尚未輪到你出牌", resBody["message"])
	})

	suite.T().Run("it can success when valid card id", func(t *testing.T) {
		id := rand.Intn(playerCount)
		player := game.Players[id]

		// Random card id
		num := rand.Intn(len(player.PlayerCards))
		cardId := player.PlayerCards[num].CardID

		// Set player to current player
		game.CurrentPlayerID = player.ID
		_ = suite.db.Save(&game)

		url := strings.ReplaceAll(api, "{player_id}", strconv.Itoa(int(player.ID)))

		req := utils.RequestToJsonBody(PlayCardRequest{CardID: cardId})

		res := suite.requestJson(url, req, http.MethodPost)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}

type PlayCardRequest struct {
	CardID uint `json:"card_id"`
}
