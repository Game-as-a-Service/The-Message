package main

import (
	_ "github.com/Game-as-a-Service/The-Message/cmd/app/docs"
	"github.com/Game-as-a-Service/The-Message/config"
	"github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"
	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			The Message API
// @description	This is an online version of the "The Message" board game backend API
// @host			127.0.0.1:8080
func main() {
	db := config.NewDatabase()

	engine := gin.Default()
	sse := http.NewSSEServer()

	gameRepo := mysqlRepo.NewGameRepository(db)
	playerRepo := mysqlRepo.NewPlayerRepository(db)
	cardRepo := mysqlRepo.NewCardRepository(db)
	deckRepo := mysqlRepo.NewDeckRepository(db)
	playerCardRepo := mysqlRepo.NewPlayerCardRepository(db)
	gameProgressRepo := mysqlRepo.NewGameProgressRepository(db)

	cardService := service.NewCardService(&service.CardServiceOptions{
		CardRepo:       cardRepo,
		PlayerRepo:     playerRepo,
		PlayerCardRepo: playerCardRepo,
		GameRepo:       gameRepo,
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

	http.RegisterGameHandler(
		&http.GameHandlerOptions{
			Engine:  engine,
			Service: gameService,
			SSE:     sse,
		},
	)

	// Register the heartbeat handler
	http.RegisterHeartbeatHandler(
		&http.HeartbeatHandler{
			Engine: engine,
		})

	http.RegisterCardHandler(
		&http.CardHandlerOptions{
			Engine:  engine,
			Service: cardService,
		},
	)

	http.RegisterPlayerHandler(
		&http.PlayerHandlerOptions{
			Engine:      engine,
			Service:     playerService,
			GameService: gameService,
			SSE:         sse,
		},
	)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
