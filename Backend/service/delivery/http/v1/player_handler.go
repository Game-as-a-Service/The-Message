package http

import (
	"fmt"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"net/http"
	"strconv"

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

	canPlay, err := p.playerService.CanPlayCard(c, playerId, req.CardID)
	if !canPlay {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := p.playerService.PlayCard(c, playerId, req.CardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// TODO to Service
	if result {
		p.gameService.NextPlayer(c, playerId)

		player, err := p.playerService.PlayerRepo.GetPlayer(c, playerId)
		game, err := p.gameService.GameRepo.GetGameWithPlayers(c, player.GameId)
		if err != nil {
			return
		}

		var currentPlayerIndex int
		for index, gPlayer := range game.Players {
			if gPlayer.Id == playerId {
				currentPlayerIndex = index
				break
			}
		}

		maxLen := len(game.Players)
		if currentPlayerIndex+1 >= maxLen {
			p.SSE.Message <- gin.H{
				"message":     "傳遞",
				"status":      "傳遞",
				"game_id":     game.Id,
				"next_player": game.Players[0].Id,
			}
		} else {
			nextId := game.Players[currentPlayerIndex+1].Id

			message := fmt.Sprintf("玩家: %d 已出牌", nextId)

			p.SSE.Message <- gin.H{
				"message":     message,
				"status":      "功能",
				"game_id":     game.Id,
				"next_player": nextId,
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
