package service

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/gin-gonic/gin"
)

type GameService struct {
	GameRepo      repository.GameRepository
	PlayerService PlayerService
	CardService   CardService
	DeckService   DeckService
}

type GameServiceOptions struct {
	GameRepo      repository.GameRepository
	PlayerService PlayerService
	CardService   CardService
	DeckService   DeckService
}

func NewGameService(opts *GameServiceOptions) GameService {
	return GameService{
		GameRepo:      opts.GameRepo,
		PlayerService: opts.PlayerService,
		CardService:   opts.CardService,
		DeckService:   opts.DeckService,
	}
}

func (g *GameService) InitGame(c context.Context, roomId string) (*repository.Game, error) {
	game, err := g.CreateGame(c, &repository.Game{
		RoomID: roomId,
		Status: enums.GameStart,
	})

	if err != nil {
		return nil, err
	}

	return game, nil
}

func (g *GameService) InitCards(c context.Context) ([]*repository.Card, error) {
	cards, err := g.CardService.GetCards(c)

	if err != nil {
		return nil, err
	}

	cards = g.CardService.ShuffleCards(c, cards)

	return cards, nil
}

func (g *GameService) DrawCard(c context.Context, player *repository.Player, drawCards []*repository.Deck, count int) error {
	for i := 0; i < count; i++ {
		card := &repository.PlayerCard{
			PlayerID: player.ID,
			CardID:   drawCards[i].ID,
			Type:     "hand",
		}
		err := g.PlayerService.CreatePlayerCard(c, card)
		if err != nil {
			return err
		}
		err = g.DeckService.DeleteDeckFromGame(c, drawCards[i].ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GameService) InitDeck(c context.Context, game *repository.Game) error {
	err := g.DeckService.InitDeck(c, game)
	if err != nil {
		return err
	}
	return nil
}

func (g *GameService) DrawCardsForAllPlayers(c context.Context, game *repository.Game) error {
	players, err := g.PlayerService.GetPlayersByGameId(c, game.ID)
	if err != nil {
		return err
	}
	for _, player := range players {
		drawCards, _ := g.DeckService.GetDecksByGameId(c, game.ID)
		err2 := g.DrawCard(c, player, drawCards, 3)
		if err2 != nil {
			return err2
		}

	}
	return nil
}

func (g *GameService) CreateGame(c context.Context, game *repository.Game) (*repository.Game, error) {
	game, err := g.GameRepo.CreateGame(c, game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *GameService) GetGameById(c context.Context, id uint) (*repository.Game, error) {
	game, err := g.GameRepo.GetGameWithPlayers(c, id)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *GameService) DeleteGame(c context.Context, id uint) error {
	err := g.GameRepo.DeleteGame(c, id)
	if err != nil {
		return err
	}
	return nil
}

func (g *GameService) UpdateCurrentPlayer(c context.Context, game *repository.Game, playerId uint) {
	game.CurrentPlayerID = playerId
	err := g.GameRepo.UpdateGame(c, game)
	if err != nil {
		panic(err)
	}
}

func (g *GameService) NextPlayer(c context.Context, player *repository.Player) (*repository.Game, error) {
	players := player.Game.Players

	currentPlayerId := player.ID

	var currentPlayerIndex int
	for index, gPlayer := range players {
		if gPlayer.ID == currentPlayerId {
			currentPlayerIndex = index
			break
		}
	}

	if currentPlayerIndex+1 >= len(players) {
		player.Game.CurrentPlayerID = players[0].ID
		player.Game.Status = enums.TransmitIntelligenceStage
	} else {
		player.Game.CurrentPlayerID = players[currentPlayerIndex+1].ID
	}
	return player.Game, nil
}

func (g *GameService) UpdateStatus(c context.Context, game *repository.Game, stage string) {
	game.Status = stage
	err := g.GameRepo.UpdateGame(c, game)
	if err != nil {
		panic(err)
	}
}

func (g *GameService) CreateGameWithPlayers(c *gin.Context, req request.CreateGameRequest, cards []*repository.Card) (*repository.Game, error) {
	// Get the number of players and assign identity cards
	length := len(req.Players)
	identityCards, err := g.AssignIdentityCards(c, length)

	if err != nil {
		return nil, err
	}

	// Create the players
	var players []repository.Player
	for i, playerReq := range req.Players {
		player := repository.Player{
			UserID:       playerReq.ID,
			Name:         playerReq.Name,
			Priority:     i,
			IdentityCard: identityCards[i],
			Status:       enums.PlayerStatusAlive,
			PlayerCards:  []repository.PlayerCard{},
		}

		// Each player gets 3 cards
		for j := 0; j < 3; j++ {
			player.PlayerCards = append(player.PlayerCards, repository.PlayerCard{
				CardID: cards[i*3+j].ID,
				Type:   "hand",
			})
		}

		players = append(players, player)
	}

	// Remove the player cards and create the deck
	deck := &repository.Deck{
		Cards: []repository.DeckCard{},
	}

	for _, card := range cards[length:] {
		deck.Cards = append(deck.Cards, repository.DeckCard{
			CardID: card.ID,
		})
	}

	// Create a game and players in one transaction
	game := &repository.Game{
		RoomID:  req.RoomID,
		Players: players,
		Status:  enums.ActionCardStage,
		Deck:    deck,
	}

	game, err = g.GameRepo.CreateGameWithPlayers(c, game)

	if err != nil {
		return nil, err
	}

	return game, nil
}

func (g *GameService) AssignIdentityCards(c context.Context, playCount int) ([]string, error) {
	var result []string

	if playCount == 3 {
		result = []string{enums.UndercoverFront, enums.MilitaryAgency, enums.Bystander}
	} else if playCount == 4 {
		result = []string{enums.UndercoverFront, enums.MilitaryAgency, enums.Bystander, enums.Bystander}
	} else if playCount == 5 {
		result = []string{enums.UndercoverFront, enums.UndercoverFront, enums.MilitaryAgency, enums.MilitaryAgency, enums.Bystander}
	} else if playCount == 6 {
		result = []string{enums.UndercoverFront, enums.UndercoverFront, enums.MilitaryAgency, enums.MilitaryAgency, enums.Bystander, enums.Bystander}
	} else if playCount == 7 {
		result = []string{enums.UndercoverFront, enums.UndercoverFront, enums.UndercoverFront, enums.MilitaryAgency, enums.MilitaryAgency, enums.MilitaryAgency, enums.Bystander}
	} else if playCount == 8 {
		result = []string{enums.UndercoverFront, enums.UndercoverFront, enums.UndercoverFront, enums.MilitaryAgency, enums.MilitaryAgency, enums.MilitaryAgency, enums.Bystander, enums.Bystander}
	} else if playCount == 9 {
		result = []string{enums.UndercoverFront, enums.UndercoverFront, enums.UndercoverFront, enums.UndercoverFront, enums.MilitaryAgency, enums.MilitaryAgency, enums.MilitaryAgency, enums.MilitaryAgency, enums.Bystander}
	} else {
		return nil, errors.New("invalid play count")
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result, nil
}
