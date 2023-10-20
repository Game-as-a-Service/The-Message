package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Game-as-a-Service/The-Message/config"
	handler "github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"

	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var serverURL string
var gameRepo *mysqlRepo.GameRepository

func TestMain(m *testing.M) {
	testDB := config.InitTestDB()

	engine := gin.Default()

	gameRepo = mysqlRepo.NewGameRepository(testDB)
	playerRepo := mysqlRepo.NewPlayerRepository(testDB)
	// gameServ := service.NewGameService(gameRepo, playerRepo)

	handler.NewGameHandler(engine, gameRepo, playerRepo)

	server := httptest.NewServer(engine)
	serverURL = server.URL

	code := m.Run()
	defer server.Close()
	os.Exit(code)
}

func TestStartGameE2E(t *testing.T) {
	players := []Player{
		{ID: "6497f6f226b40d440b9a90cc", Name: "A"},
		{ID: "6498112b26b40d440b9a90ce", Name: "B"},
		{ID: "6499df157fed0c21a4fd0425", Name: "C"},
	}
	requestBody := Request{Players: players}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	api := "/api/v1/games"
	resp := requestJson(t, api, jsonBody)

	assert.Equal(t, 200, resp.StatusCode)

	responseJson := responseJson(t, resp)

	assert.NotNil(t, responseJson["Token"], "JSON response should contain a 'Token' field")
	assert.NotNil(t, responseJson["Id"], "JSON response should contain a 'Id' field")

	// 驗證Game內的玩家都持有identity
	game, _ := gameRepo.GetGameWithPlayers(context.TODO(), int(responseJson["Id"].(float64)))

	assert.NotEmpty(t, game.Players[0].IdentityCard)
	assert.NotEmpty(t, game.Players[1].IdentityCard)
	assert.NotEmpty(t, game.Players[2].IdentityCard)
}

// Helper functions
type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Request struct {
	Players []Player `json:"players"`
}

func responseJson(t *testing.T, resp *http.Response) map[string]interface{} {
	var responseMap map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&responseMap)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}
	return responseMap
}

func requestJson(t *testing.T, api string, jsonBody []byte) *http.Response {
	req, err := http.NewRequest(http.MethodPost, serverURL+api, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
