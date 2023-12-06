package e2e

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPlayerCards(t *testing.T) {
	player := &PlayerCard{
		ID:       "1",
		PlayerID: "player1",
		GameID:   "game1",
		CardID:   "card1",
		Type:     "type1",
	}

	// Create a new Gin router
	router := gin.Default()

	// Create a new PlayerHandler instance
	playerHandler := &v1.PlayerHandler{}

	// Define the request
	playerID := player.ID // Replace with a valid player ID
	req, _ := http.NewRequest(http.MethodGet, "/players/"+playerID+"/cards", nil)
	router.GET("/players/:playerId/cards", playerHandler.GetPlayerCards)

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body
	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	// Assert the expected response
	expectedResponse := map[string]interface{}{
		"player_cards": []map[string]interface{}{},
	}
	assert.Equal(t, expectedResponse, responseBody)
}

type PlayerCard struct {
	ID       string `json:"id"`
	PlayerID string `json:"player_id"`
	GameID   string `json:"game_id"`
	CardID   string `json:"card_id"`
	Type     string `json:"type"`
}
