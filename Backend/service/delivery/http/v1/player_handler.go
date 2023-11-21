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

// GetPlayerCards godoc
// @Summary GetPlayerCards
// @Description GetPlayerCardsByPlayerId
// @Tags player_cards
// @Produce json
// @Param id path int true "Player ID"
// @Success 200 {object} request.PlayerCardsResponse
// @Router /api/v1/player_cards/{id} [get]
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
