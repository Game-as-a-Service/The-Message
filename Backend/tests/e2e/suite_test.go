package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Game-as-a-Service/The-Message/config"
	"github.com/Game-as-a-Service/The-Message/database/seeders"
	v1 "github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type IntegrationTestSuite struct {
	suite.Suite
	db             *gorm.DB
	tx             *gorm.DB
	server         *httptest.Server
	gameRepo       repository.GameRepository
	playerRepo     repository.PlayerRepository
	playerCardRepo repository.PlayerCardRepository
	gameServ       *service.GameService
	playerServ     *service.PlayerService
}

func (suite *IntegrationTestSuite) SetupSuite() {
	sourceURL := config.GetSourceURL()

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := config.BaseTestDSN()
	val := url.Values{}
	val.Add("multiStatements", "true")
	dsn = fmt.Sprintf("%s?%s", dsn, val.Encode())

	m, err := config.NewMigration(dsn, sourceURL)
	if err != nil {
		panic(err)
	}

	err = m.Down()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("no change")
		} else {
			panic(err)
		}
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("no change")
		} else {
			panic(err)
		}
	}
	db := config.NewDatabase()

	seeders.SeederCards(db)

	engine := gin.Default()
	sse := v1.NewSSEServer()

	gameRepo := mysqlRepo.NewGameRepository(db)
	playerRepo := mysqlRepo.NewPlayerRepository(db)
	cardRepo := mysqlRepo.NewCardRepository(db)
	deckRepo := mysqlRepo.NewDeckRepository(db)
	playerCardRepo := mysqlRepo.NewPlayerCardRepository(db)
	gameProgressRepo := mysqlRepo.NewGameProgressRepository(db)

	cardService := service.NewCardService(&service.CardServiceOptions{
		CardRepo:       cardRepo,
		GameRepo:       gameRepo,
		PlayerRepo:     playerRepo,
		PlayerCardRepo: playerCardRepo,
	})

	deckService := service.NewDeckService(&service.DeckServiceOptions{
		DeckRepo:    deckRepo,
		CardService: cardService,
	})

	playerService := service.NewPlayerService(&service.PlayerServiceOptions{
		PlayerRepo:       playerRepo,
		PlayerCardRepo:   playerCardRepo,
		GameRepo:         gameRepo,
		GameProgressRepo: gameProgressRepo,
	})

	gameService := service.NewGameService(
		&service.GameServiceOptions{
			GameRepo:      gameRepo,
			PlayerService: playerService,
			CardService:   cardService,
			DeckService:   deckService,
		},
	)
	playerService.GameServ = &gameService

	v1.RegisterGameHandler(
		&v1.GameHandlerOptions{
			Engine:  engine,
			Service: gameService,
			SSE:     sse,
		},
	)

	v1.RegisterCardHandler(
		&v1.CardHandlerOptions{
			Engine:  engine,
			Service: cardService,
		},
	)

	v1.RegisterPlayerHandler(
		&v1.PlayerHandlerOptions{
			Engine:      engine,
			Service:     playerService,
			GameService: gameService,
			SSE:         sse,
		},
	)

	server := httptest.NewServer(engine)

	suite.db = db
	suite.server = server
	suite.gameRepo = gameRepo
	suite.playerRepo = playerRepo
	suite.gameServ = &gameService
	suite.playerServ = &playerService
	suite.playerCardRepo = playerCardRepo
}

func (suite *IntegrationTestSuite) TearDownSuite() {
	sqlDB, _ := suite.db.DB()
	err := sqlDB.Close()
	if err != nil {
		return
	}

	suite.server.Close()
}

func (suite *IntegrationTestSuite) SetupTest() {
	suite.tx = suite.db.Begin()

	//Fixme Run db refresh and seeders
	config.RunRefresh()
	db := config.NewDatabase()
	seeders.Run(db)
}

func (suite *IntegrationTestSuite) TearDownTest() {
	suite.tx.Rollback()
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) responseJson(resp *http.Response) map[string]interface{} {
	var responseMap map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&responseMap)
	if err != nil {
		suite.T().Fatalf("Failed to decode JSON: %v", err)
	}
	return responseMap
}

func (suite *IntegrationTestSuite) requestJson(api string, jsonBody []byte, method string) *http.Response {
	req, err := http.NewRequest(method, suite.server.URL+api, bytes.NewBuffer(jsonBody))
	if err != nil {
		suite.T().Fatalf("Failed to send request: %v", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func (suite *IntegrationTestSuite) responseTest(resp *http.Response) interface{} {
	var responseMap interface{}
	err := json.NewDecoder(resp.Body).Decode(&responseMap)
	if err != nil {
		suite.T().Fatalf("Failed to decode JSON: %v", err)
	}
	return responseMap
}
