package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Game-as-a-Service/The-Message/enums"
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

	// TODO 這邊可以優化 https://gorm.io/zh_CN/docs/associations.html
	if err := g.gameService.PlayerService.InitPlayers(c, game, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	game, _ = g.gameService.GetGameById(c, game.Id)
	g.gameService.UpdateCurrentPlayer(c, game, game.Players[0].Id)
	g.gameService.UpdateStatus(c, game, enums.ActionCardStage)

	if err := g.gameService.InitDeck(c, game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := g.gameService.DrawCardsForAllPlayers(c, game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	game, err = g.gameService.GetGameById(c, game.Id)
	if err != nil {
		return
	}

	g.SSE.Message <- gin.H{
		"message":     "Game started",
		"status":      "started",
		"game_id":     game.Id,
		"next_player": game.Players[0].Id,
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

// GameEvent godoc
// @Summary Get game events
// @Description Get game events
// @Tags games
// @Accept json
// @Produce json
// @Param gameId path int true "Game ID"
// @Success 200 {object} GameSSERequest
// @Router /api/v1/games/{gameId}/events [get]
func (g *GameHandler) GameEvent(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("gameId"))
	if err != nil {
		return
	}

	v, ok := c.Get("clientChan")
	if !ok {
		return
	}

	clientChan, ok := v.(ClientChan)
	if !ok {
		return
	}

	game, err := g.gameService.GetGameById(c, gameId)
	if err != nil {
		return
	}

	g.SSE.Message <- gin.H{
		"message":        game.Status,
		"status":         game.Status,
		"game_id":        gameId,
		"current_player": game.CurrentPlayerId,
	}

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-clientChan; ok {
			log.Printf("msg: %+v", msg)
			data := GameSSERequest{}
			err := json.Unmarshal([]byte(msg), &data)
			if err != nil {
				log.Fatalf(err.Error())
			}

			if data.GameId == gameId {
				c.SSEvent("message", msg)
			}
			return true
		}
		return false
	})
}

type GameSSERequest struct {
	GameId  int    `json:"game_id,int"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
