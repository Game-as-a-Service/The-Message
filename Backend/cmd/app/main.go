package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"

	http "github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"
	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("../../.env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %v\n", err)
	}
}

func main() {
	dbHost := viper.GetString("DB_HOST")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("DB_USER")
	dbPwd := viper.GetString("DB_PASSWORD")
	dbPort := viper.GetString("DB_PORT")

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbHost, dbPort, dbDatabase)), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	gameRepo := mysqlRepo.NewGameRepositoryRepository(db)

	gameHandler := &http.Game{
		GameRepo: gameRepo,
	}
	db.Table("games").AutoMigrate(&mysqlRepo.Game{})

	engine := gin.Default()

	engine.POST("/api/v1/game", gameHandler.CreateGame)
	engine.GET("/api/v1/game/:gameId", gameHandler.GetGameById)

	engine.Run(":8080")
}
