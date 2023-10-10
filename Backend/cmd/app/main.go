package main

import (
	"github.com/Game-as-a-Service/The-Message/database"
	"github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"

	http "github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"
	repository "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.InitDB()

	engine := gin.Default()

	gameRepo := repository.NewGameRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	gameServ := service.NewGameService(gameRepo, playerRepo)

	http.NewGameHandler(engine, gameServ)

	engine.Run(":8080")
}
