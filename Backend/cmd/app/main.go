package main

import (
	"github.com/Game-as-a-Service/The-Message/config"
	"github.com/gin-gonic/gin"

	http "github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"
	repository "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := config.InitDB()

	engine := gin.Default()

	gameRepo := repository.NewGameRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	// gameServ := service.NewGameService(gameRepo, playerRepo)

	http.NewGameHandler(engine, gameRepo, playerRepo)

	engine.Run(":8080")
}
