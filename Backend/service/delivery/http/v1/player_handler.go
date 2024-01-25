package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Game-as-a-Service/The-Message/enums"
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
	opts.Engine.POST("/api/v1/player/:playerId/transmit-intelligence", handler.TransmitIntelligence)
	opts.Engine.POST("/api/v1/players/:playerId/accept", handler.AcceptCard)
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

// TransmitIntelligence godoc
// @Summary Transmit intelligence
// @Description Transmit an intelligence card
// @Tags players
// @Accept json
// @Produce json
// @Param playerId path int true "Player ID"
// @Param card_id body request.PlayCardRequest true "Card ID"
// @Param intelligence_type body request.PlayCardRequest true "Intelligence Type"
// @Success 200 {object} request.PlayCardResponse
// @Router /api/v1/player/{playerId}/transmit-intelligence [post]
func (p *PlayerHandler) TransmitIntelligence(c *gin.Context) {
	playerId, _ := strconv.Atoi(c.Param("playerId"))
	var req request.PlayCardRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	intelligenceType := enums.ToString(req.IntelligenceType)

	if intelligenceType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid intelligence type"})
		return
	}

	player, err := p.playerService.GetPlayerById(c, playerId)

	if err != nil || player == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Player not found"})
		return
	}

	// Check card_id exists in player_cards
	exist, err := p.playerService.CheckPlayerCardExist(c, playerId, player.GameId, req.CardID)
	if err != nil || !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Card not found"})
		return
	}

	ret, err := p.playerService.TransmitIntelligenceCard(c, playerId, player.GameId, req.CardID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  ret,
		"message": enums.ToString(req.IntelligenceType) + " intelligence transmitted",
	})
}

// AcceptCard godoc
// @Summary Accept Card
// @Description Decide accept card or not
// @Tags players
// @Accept json
// @Produce json
// @Param playerId path int true "Player ID"
// @Param accept body request.AcceptCardRequest true "Accept"
// @Success 200 {object} request.PlayCardResponse
// @Router /api/v1/players/{playerId}/accept [post]
func (p *PlayerHandler) AcceptCard(c *gin.Context) {
	playerId, _ := strconv.Atoi(c.Param("playerId"))
	var req request.AcceptCardRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := p.playerService.AcceptCard(c, playerId, req.Accept)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
