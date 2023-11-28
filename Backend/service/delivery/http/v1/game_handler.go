package http

import (
	"io"
	"net/http"
	"strconv"

	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gameService service.GameService
	SSE         *Event
}

type GameHandlerOptions struct {
	Engine  *gin.Engine
	Service service.GameService
	SSE     *Event
}

func RegisterGameHandler(opts *GameHandlerOptions) {
	handler := &GameHandler{
		gameService: opts.Service,
		SSE:         opts.SSE,
	}

	opts.Engine.POST("/api/v1/games", handler.StartGame)
	opts.Engine.GET("/api/v1/games/:gameId/events", HeadersMiddleware(), opts.SSE.serveHTTP(), handler.GameEvent)
}

// StartGame godoc
// @Summary Start a new game
// @Description Start a new game
// @Tags games
// @Accept json
// @Produce json
// @Param players body request.CreateGameRequest true "Players"
// @Success 200 {object} request.CreateGameResponse
// @Router /api/v1/games [post]
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

	if err := g.gameService.DrawCardsForAllPlayers(c, game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	g.SSE.Message <- gin.H{
		"message": "Game started",
		"game":    game,
	}

	c.JSON(http.StatusOK, gin.H{
		"Id":    game.Id,
		"Token": game.Token,
	})
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

func (g *GameHandler) DeleteGame(c *gin.Context) {
	gameId, _ := strconv.Atoi(c.Param("gameId"))

	err := g.gameService.DeleteGame(c, gameId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted"})
}

func (g *GameHandler) GameEvent(c *gin.Context) {
	v, ok := c.Get("clientChan")
	if !ok {
		return
	}
	clientChan, ok := v.(ClientChan)
	if !ok {
		return
	}
	c.Stream(func(w io.Writer) bool {
		// Stream message to client from message channel
		if msg, ok := <-clientChan; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}
