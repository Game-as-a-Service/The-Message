package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Game-as-a-Service/The-Message/service/request"

	"github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	playerService service.PlayerService
	gameService   service.GameService
	SSE           *Event
}

type PlayerHandlerOptions struct {
	Engine      *gin.Engine
	Service     service.PlayerService
	GameService service.GameService
	SSE         *Event
}

func RegisterPlayerHandler(opts *PlayerHandlerOptions) {
	handler := &PlayerHandler{
		playerService: opts.Service,
		gameService:   opts.GameService,
		SSE:           opts.SSE,
	}

	opts.Engine.POST("/api/v1/players/:playerId/player-cards", handler.PlayCard)
	// opts.Engine.POST("/api/v1/players/:playerId/accept", handler.AcceptCard)
}

// PlayCard godoc
// @Summary Play a card
// @Description Play a card
// @Tags players
// @Accept json
// @Produce json
// @Param playerId path int true "Player ID"
// @Param card_id body request.PlayCardRequest true "Card ID"
// @Success 200 {object} request.PlayCardResponse
// @Router /api/v1/players/{playerId}/player-cards [post]
func (p *PlayerHandler) PlayCard(c *gin.Context) {
	playerId, _ := strconv.Atoi(c.Param("playerId"))
	var req request.PlayCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, card, err := p.playerService.PlayCard(c, playerId, req.CardID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// TODO to Service
	p.SSE.Message <- gin.H{
		"game_id":     game.Id,
		"status":      game.Status,
		"message":     fmt.Sprintf("玩家: %d 已出牌", playerId),
		"card":        card.Name,
		"next_player": game.CurrentPlayerId,
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}

// func (p *PlayerHandler) AcceptCard(c *gin.Context) {
// 	playerId, _ := strconv.Atoi(c.Param("playerId"))

// 	result, err := p.playerService.AcceptCard(c, playerId)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "success",
// 	})
// }
