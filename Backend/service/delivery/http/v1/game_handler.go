package http

import (
	"net/http"
	"strconv"

	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"github.com/gin-gonic/gin"
)

type Game struct {
	GameRepo *mysqlRepo.GameRepository
}

func (g *Game) GetGameById(c *gin.Context) {
	gameId, _ := strconv.Atoi(c.Param("gameId"))

	game, err := g.GameRepo.GetGameById(c, gameId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mysqlRepo.Game{
		Id:   game.Id,
		Name: game.Name,
	})
}

func (g *Game) CreateGame(c *gin.Context) {
	game := &mysqlRepo.Game{
		Name: "Game",
	}

	game, err := g.GameRepo.CreateGame(c, game)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mysqlRepo.Game{
		Id:   game.Id,
		Name: game.Name,
	})
}
