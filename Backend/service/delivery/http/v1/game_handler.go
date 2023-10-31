package http

import (
	"net/http"
	"strconv"

	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gameService service.GameService
}

type GameHandlerOptions struct {
	Engine  *gin.Engine
	Service service.GameService
}

func RegisterGameHandler(opts *GameHandlerOptions) {
	handler := &GameHandler{
		gameService: opts.Service,
	}

	opts.Engine.POST("/api/v1/games", handler.StartGame)
	opts.Engine.Static("/swagger", "./web/swagger-ui")
}

func (g *GameHandler) GetGame(c *gin.Context) {
	gameId, _ := strconv.Atoi(c.Param("gameId"))

	game, err := g.gameService.GetGameById(c, gameId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Id":    game.Id,
		"Token": game.Token,
	})
}

func (g *GameHandler) StartGame(c *gin.Context) {
	var req request.CreateGameRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, err := g.gameService.InitGame(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := g.gameService.PlayerService.InitPlayers(c, game, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := g.gameService.InitDeck(c, game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := g.gameService.DrawCardsForPlayers(c, game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Id":    game.Id,
		"Token": game.Token,
	})
}

func (g *GameHandler) DeleteGame(c *gin.Context) {
	gameId, _ := strconv.Atoi(c.Param("gameId"))

	err := g.gameService.DeleteGame(c, gameId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted"})
}
