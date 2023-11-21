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

	opts.Engine.GET("/api/v1/player_cards/:playerId", handler.GetPlayerCards)
}

func (p *PlayerHandler) GetPlayerCards(c *gin.Context) {
	playerId, _ := strconv.Atoi(c.Param("playerId"))
	cards, err := p.playerService.GetPlayerCardsByPlayerId(c, playerId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Id":    playerId,
		"Cards": cards,
	})
}
