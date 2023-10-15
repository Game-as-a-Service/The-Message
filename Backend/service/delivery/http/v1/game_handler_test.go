package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Game-as-a-Service/The-Message/config"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/service/service"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var serverURL string
var gameRepo repository.GameRepository

func TestMain(m *testing.M) {
	testDB := config.InitDB()

	engine := gin.Default()

	gameRepo = mysqlRepo.NewGameRepository(testDB)
	playerRepo := mysqlRepo.NewPlayerRepository(testDB)
	gameServ := service.NewGameService(gameRepo, playerRepo)

	NewGameHandler(engine, gameServ)

	server := httptest.NewServer(engine)
	serverURL = server.URL

	code := m.Run()
	defer server.Close()
	os.Exit(code)
}

//func TestGetGameByIdE2E(t *testing.T) {
//	// Create a virtual MySQL database connection
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("Error creating mock database: %v", err)
//	}
//	defer db.Close()
//
//	// Replace the GORM database connection with the virtual one
//	gdb, err := gorm.Open(mysql.New(mysql.Config{
//		Conn:                      db,
//		SkipInitializeWithVersion: true,
//	}), &gorm.Config{})
//	if err != nil {
//		t.Fatalf("Error opening GORM database: %v", err)
//	}
//
//	// Set up expected mock database queries and operations
//	mock.ExpectQuery(regexp.QuoteMeta("SELECT `id`,`name` FROM `games` WHERE id = ? ORDER BY `games`.`id` LIMIT 1")).
//		WithArgs(1).
//		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Test Game"))
//
//	// Create a Gin router and HTTP server
//	router := gin.Default()
//	gameHandler := &GameHandler{
//		GameRepo: mysqlRepo.NewGameRepository(gdb),
//	}
//	router.GET("/api/v1/game/:gameId", gameHandler.GetGame)
//
//	// Prepare an HTTP GET request
//	req, _ := http.NewRequest("GET", "/api/v1/game/1", nil)
//	recorder := httptest.NewRecorder()
//
//	// Execute the HTTP request
//	router.ServeHTTP(recorder, req)
//
//	// Check the HTTP response
//	assert.Equal(t, http.StatusOK, recorder.Code)
//
//	// Parse the HTTP response
//	var response mysqlRepo.Game
//	err = json.Unmarshal(recorder.Body.Bytes(), &response)
//	assert.NoError(t, err)
//
//	// Check if the mock database expectations were met
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("Unfulfilled expectations: %s", err)
//	}
//
//	// Check if the database operations were correct; this is achieved through mocking
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("Unfulfilled expectations: %s", err)
//	}
//}

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

//func TestDeleteGameE2E(t *testing.T) {
//	// Create a virtual MySQL database connection
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("Error creating mock database: %v", err)
//	}
//	defer db.Close()
//
//	// Replace the GORM database connection with the virtual one
//	gdb, err := gorm.Open(mysql.New(mysql.Config{
//		Conn:                      db,
//		SkipInitializeWithVersion: true,
//	}), &gorm.Config{})
//	if err != nil {
//		t.Fatalf("Error opening GORM database: %v", err)
//	}
//
//	mock.ExpectBegin() // Expect a transaction Begin
//	mock.ExpectExec("INSERT INTO `games`").WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectCommit() // Expect a transaction Commit
//
//	// Set up expected mock database queries and operations for DeleteGame
//	gameId := 1
//	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `games` WHERE `id` = ?")).
//		WithArgs(gameId).
//		WillReturnResult(sqlmock.NewResult(0, 1)) // Assumes the game with ID=1 is deleted successfully.
//
//	// Create a Gin router and HTTP handler
//	router := gin.Default()
//	gameHandler := &GameHandler{
//		GameRepo: mysqlRepo.NewGameRepository(gdb),
//	}
//	router.DELETE("/api/v1/game/:gameId", gameHandler.DeleteGame)
//
//	// Prepare an HTTP DELETE request
//	req, _ := http.NewRequest("DELETE", "/api/v1/game/1", nil)
//	recorder := httptest.NewRecorder()
//
//	// Execute the HTTP request
//	router.ServeHTTP(recorder, req)
//
//	// Check the HTTP response status code
//	assert.Equal(t, http.StatusOK, recorder.Code)
//
//	// Optional: Check the HTTP response body
//	var response mysqlRepo.Game
//	err = json.Unmarshal(recorder.Body.Bytes(), &response)
//	assert.NoError(t, err)
//
//	// Check if the mock database expectations were met
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("Unfulfilled expectations: %s", err)
//	}
//}

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
