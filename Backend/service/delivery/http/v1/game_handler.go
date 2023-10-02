package http

import (
	"net/http"
	"strconv"

	repository "github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	GameRepo repository.GameRepository
}

func (g *GameHandler) GetGame(c *gin.Context) {
	gameId, _ := strconv.Atoi(c.Param("gameId"))

	game, err := g.GameRepo.GetGameById(c, gameId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repository.Game{
		Id:    game.Id,
		Token: game.Token,
	})
}

func (g *GameHandler) CreateGame(c *gin.Context) {
	game := new(repository.Game)
	game.Token = "Game"

	game, err := g.GameRepo.CreateGame(c, game)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repository.Game{
		Id:    game.Id,
		Token: game.Token,
	})
}

func (g *GameHandler) DeleteGame(c *gin.Context) {
	gameId, _ := strconv.Atoi(c.Param("gameId"))

	err := g.GameRepo.DeleteGame(c, gameId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted"})
}
