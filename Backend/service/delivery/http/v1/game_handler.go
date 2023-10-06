package http

import (
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/Game-as-a-Service/The-Message/service/service"
	"net/http"
	"strconv"

	repository "github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	GameRepo repository.GameRepository
	GameServ *service.GameService
}

func NewGameHandler(engine *gin.Engine, gameServ *service.GameService) *GameHandler {
	gameRepo := gameServ.GameRepo
	handler := &GameHandler{
		GameRepo: gameRepo,
		GameServ: gameServ,
	}
	engine.POST("/api/v1/games", handler.StartGame)
	engine.Static("/swagger", "./web/swagger-ui")
	return handler
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

func (g *GameHandler) StartGame(c *gin.Context) {
	var req request.CreateGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, err := g.GameServ.StartGame(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
