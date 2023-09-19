package main

import (
	"fmt"
	"log"

	"os"

	"github.com/gin-gonic/gin"

	http "github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"
	repository "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbHost, dbPort, dbDatabase)), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	gameRepo := repository.NewGameRepository(db)

	gameHandler := &http.GameHandler{
		GameRepo: gameRepo,
	}
	db.Table("games").AutoMigrate(&repository.Game{})

	engine := gin.Default()

	engine.POST("/api/v1/game", gameHandler.CreateGame)
	engine.GET("/api/v1/game/:gameId", gameHandler.GetGameById)

	engine.Run(":8080")
}
