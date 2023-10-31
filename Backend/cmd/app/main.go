package main

import (
	"github.com/Game-as-a-Service/The-Message/config"
	http "github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"
	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := config.InitDB()

	engine := gin.Default()

	gameRepo := mysqlRepo.NewGameRepository(db)
	playerRepo := mysqlRepo.NewPlayerRepository(db)
	cardRepo := mysqlRepo.NewCardRepository(db)
	deckRepo := mysqlRepo.NewDeckRepository(db)
	playerCardRepo := mysqlRepo.NewPlayerCardRepository(db)

	cardService := service.NewCardService(&service.CardServiceOptions{
		CardRepo: cardRepo,
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

	http.RegisterGameHandler(
		&http.GameHandlerOptions{
			Engine:  engine,
			Service: gameService,
		},
	)

	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
