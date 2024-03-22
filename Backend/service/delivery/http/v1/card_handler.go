package http

import (
	"net/http"
	"strconv"

	"github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	cardService service.CardService
}

type CardHandlerOptions struct {
	Engine  *gin.Engine
	Service service.CardService
}

func RegisterCardHandler(opts *CardHandlerOptions) {
	handler := &CardHandler{
		cardService: opts.Service,
	}

	opts.Engine.GET("/api/v1/player/:playerId/player-cards/", handler.GetPlayerCards)
}

// GetPlayerCards godoc
// @Summary GetPlayerCards
// @Description GetPlayerCardsByPlayerId
// @Tags player_cards
// @Produce json
// @Param id path int true "Player ID"
// @Success 200 {object} request.PlayerCardsResponse
// @Router /api/v1/player/{id}/player-cards/ [get]
func (p *CardHandler) GetPlayerCards(c *gin.Context) {
	reqPlayerId, _ := strconv.Atoi(c.Param("playerId"))
	playerId := uint(reqPlayerId)

	playerCards, err := p.cardService.GetPlayerCardsByPlayerId(c, playerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var playerCardsInfo []map[string]interface{}

	for _, card := range playerCards {
		dict := map[string]interface{}{
			"id":    card.ID,
			"name":  card.Name,
			"color": card.Color,
		}
		playerCardsInfo = append(playerCardsInfo, dict)
	}

	c.JSON(http.StatusOK, gin.H{"player_cards": playerCardsInfo})
}
