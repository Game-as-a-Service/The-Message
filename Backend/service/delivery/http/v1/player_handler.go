package http

import (
	"github.com/Game-as-a-Service/The-Message/service/request"
	"net/http"
	"strconv"

	"github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	playerService service.PlayerService
}

type PlayerHandlerOptions struct {
	Engine  *gin.Engine
	Service service.PlayerService
}

func RegisterPlayerHandler(opts *PlayerHandlerOptions) {
	handler := &PlayerHandler{
		playerService: opts.Service,
	}

	opts.Engine.POST("/api/v1/players/:playerId/player-cards", handler.PlayCard)
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

	result, err := p.playerService.PlayCard(c, playerId, req.CardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
