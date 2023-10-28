package main

import (
	"github.com/Game-as-a-Service/The-Message/config"
	"github.com/gin-gonic/gin"

	http "github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"
	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
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

	http.NewGameHandler(engine, gameRepo, playerRepo, cardRepo, deckRepo, playerCardRepo)

	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
