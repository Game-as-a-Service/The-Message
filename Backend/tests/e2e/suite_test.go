package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	_ "github.com/mattes/migrate/source/file"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
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
	dir, _ := os.Getwd()
	sourceURL := "file://" + dir + "/../../database/migrations"

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

	gameRepo := mysqlRepo.NewGameRepository(db)
	playerRepo := mysqlRepo.NewPlayerRepository(db)
	cardRepo := mysqlRepo.NewCardRepository(db)
	deckRepo := mysqlRepo.NewDeckRepository(db)
	playerCardRepo := mysqlRepo.NewPlayerCardRepository(db)

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
		PlayerRepo:     playerRepo,
		PlayerCardRepo: playerCardRepo,
	})

	gameService := service.NewGameService(
		&service.GameServiceOptions{
			GameRepo:      gameRepo,
			PlayerService: playerService,
			CardService:   cardService,
			DeckService:   deckService,
		},
	)

	v1.RegisterGameHandler(
		&v1.GameHandlerOptions{
			Engine:  engine,
			Service: gameService,
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
			Engine:  engine,
			Service: playerService,
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
	sqlDB.Close()

	suite.server.Close()
}

func (suite *IntegrationTestSuite) SetupTest() {
	suite.tx = suite.db.Begin()
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
