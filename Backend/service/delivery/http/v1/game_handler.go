package http

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"github.com/Game-as-a-Service/The-Message/service/request"
	service "github.com/Game-as-a-Service/The-Message/service/service"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	GameRepo       repository.GameRepository
	PlayerRepo     repository.PlayerRepository
	CardRepo       repository.CardRepository
	DeckRepo       repository.DeckRepository
	PlayerCardRepo repository.PlayerCardRepository
}

func NewGameHandler(
	engine *gin.Engine,
	gameRepo *mysqlRepo.GameRepository,
	playerRepo *mysqlRepo.PlayerRepository,
	cardRepo *mysqlRepo.CardRepository,
	deckRepo *mysqlRepo.DeckRepository,
	playerCardRepo *mysqlRepo.PlayerCardRepository,
) *GameHandler {
	handler := &GameHandler{
		GameRepo:       gameRepo,
		PlayerRepo:     playerRepo,
		CardRepo:       cardRepo,
		DeckRepo:       deckRepo,
		PlayerCardRepo: playerCardRepo,
	}
	engine.POST("/api/v1/games", handler.StartGame)
	engine.Static("/swagger", "./web/swagger-ui")
	return handler
}

func (g *GameHandler) GetGame(c *gin.Context) {
	gameId, _ := strconv.Atoi(c.Param("gameId"))

	game, err := g.GameRepo.GetGameById(c, gameId)

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

	game, err := g.GameRepo.CreateGame(c, game)

	identityCards := service.InitIdentityCards(len(req.Players))
	for i, reqPlayer := range req.Players {
		player := new(repository.Player)
		player.Name = reqPlayer.Name
		player.GameId = game.Id
		player.IdentityCard = identityCards[i]
		_, err = g.PlayerRepo.CreatePlayer(c, player)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	cards, err := g.CardRepo.GetCards(c)
	cards = service.InitialDeck(game.Id, cards)
	for _, card := range cards {
		deck := new(repository.Deck)
		deck.GameId = game.Id
		deck.CardId = card.Id
		_, err := g.DeckRepo.CreateDeck(c, deck)
		if err != nil {
			return
		}
	}

	players, err := g.PlayerRepo.GetPlayersByGameId(c, game.Id)
	for _, player := range players {
		drawCards, _ := g.DeckRepo.GetDecksByGameId(c, game.Id)
		for i := 0; i < 3; i++ {
			playerCards := new(repository.PlayerCard)
			playerCards.GameId = game.Id
			playerCards.PlayerId = player.Id
			playerCards.CardId = drawCards[i].CardId
			playerCards.Type = "hand"
			_, err := g.PlayerCardRepo.CreatePlayerCard(c, playerCards)
			if err != nil {
				return
			}
			// delete deck
			err = g.DeckRepo.DeleteDeck(c, drawCards[i].Id)
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

	err := g.GameRepo.DeleteGame(c, gameId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted"})
}
