package http

import (
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
// @Param card_id body int true "Card ID"
// @Success 200 {object} string
// @Router /api/v1/players/{playerId}/player-cards [post]
func (p *PlayerHandler) PlayCard(c *gin.Context) {
	playerId, _ := strconv.Atoi(c.Param("playerId"))
	json := make(map[string]int)
	err := c.BindJSON(&json)
	if err != nil {
		return
	}

	result, err := p.playerService.PlayCard(c, playerId, json["card_id"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
