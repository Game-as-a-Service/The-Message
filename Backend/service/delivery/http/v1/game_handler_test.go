package http

import (
	// "bytes"
	// "encoding/json"
	// "net/http"
	// "net/http/httptest"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-sqlmock"
	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetGameByIdE2E(t *testing.T) {
	// Create a virtual MySQL database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Replace the GORM database connection with the virtual one
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM database: %v", err)
	}

	// Set up expected mock database queries and operations
	mock.ExpectQuery(regexp.QuoteMeta("SELECT `id`,`name` FROM `games` WHERE id = ? AND `games`.`deleted_at` IS NULL ORDER BY `games`.`id` LIMIT 1")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Test Game"))

	// Create a Gin router and HTTP server
	router := gin.Default()
	gameHandler := &GameHandler{
		GameRepo: mysqlRepo.NewGameRepository(gdb),
	}
	router.GET("/api/v1/game/:gameId", gameHandler.GetGameById)

	// Prepare an HTTP GET request
	req, _ := http.NewRequest("GET", "/api/v1/game/1", nil)
	recorder := httptest.NewRecorder()

	// Execute the HTTP request
	router.ServeHTTP(recorder, req)

	// Check the HTTP response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the HTTP response
	var response mysqlRepo.Game
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Check if the mock database expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

	// Check if the database operations were correct; this is achieved through mocking
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestCreateGameE2E(t *testing.T) {
	// Create a virtual MySQL database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Replace the GORM database connection with the virtual one
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM database: %v", err)
	}

	// Set up expected mock database queries and operations for CreateGame
	mock.ExpectBegin() // Expect a transaction Begin
	mock.ExpectExec("INSERT INTO `games`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit() // Expect a transaction Commit

	// Create a Gin router and HTTP handler
	router := gin.Default()
	gameHandler := &GameHandler{
		GameRepo: mysqlRepo.NewGameRepository(gdb),
	}
	router.POST("/api/v1/game", gameHandler.CreateGame) // Register the CreateGame route

	// Prepare an HTTP POST request to create a game
	createGameRequest := `{"name": "Test Game"}`
	req, _ := http.NewRequest("POST", "/api/v1/game", strings.NewReader(createGameRequest))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Execute the HTTP request
	router.ServeHTTP(recorder, req)

	// Check the HTTP response status code
	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d. Response Body: %s", recorder.Code, recorder.Body.String())
	}

	// Check if the mock database expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
