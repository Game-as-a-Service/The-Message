package http

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gameService service.GameService
}

type GameHandlerOptions struct {
	Engine  *gin.Engine
	Service service.GameService
}

func RegisterGameHandler(opts *GameHandlerOptions) {
	handler := &GameHandler{
		gameService: opts.Service,
	}

	opts.Engine.POST("/api/v1/games", handler.StartGame)
	opts.Engine.Static("/swagger", "./web/swagger-ui")
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

func (g *GameHandler) StartGame(c *gin.Context) {
	var req request.CreateGameRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game := new(repository.Game)
	jwtToken := "the-message" // 先亂寫Token
	jwtBytes := []byte(jwtToken)
	hash := sha256.Sum256(jwtBytes)
	hashString := hex.EncodeToString(hash[:])
	game.Token = hashString

	game, err := g.gameService.CreateGame(c, game)

	identityCards := g.gameService.InitIdentityCards(len(req.Players))
	for i, reqPlayer := range req.Players {
		player := new(repository.Player)
		player.Name = reqPlayer.Name
		player.GameId = game.Id
		player.IdentityCard = identityCards[i]
		_, err = g.gameService.CreatePlayer(c, player)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}

	cards, err := g.gameService.GetCards(c)
	cards = g.gameService.InitialDeck(game.Id, cards)
	for _, card := range cards {
		deck := new(repository.Deck)
		deck.GameId = game.Id
		deck.CardId = card.Id
		_, err := g.gameService.CreateDeck(c, deck)
		if err != nil {
			return
		}
	}

	players, err := g.gameService.GetPlayersByGameId(c, game.Id)
	for _, player := range players {
		drawCards, _ := g.gameService.GetDecksByGameId(c, game.Id)
		for i := 0; i < 3; i++ {
			playerCards := new(repository.PlayerCard)
			playerCards.GameId = game.Id
			playerCards.PlayerId = player.Id
			playerCards.CardId = drawCards[i].CardId
			playerCards.Type = "hand"
			_, err := g.gameService.CreatePlayerCard(c, playerCards)
			if err != nil {
				return
			}
			// delete deck
			err = g.gameService.DeleteDeck(c, drawCards[i].Id)
			if err != nil {
				return
			}
		}

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
