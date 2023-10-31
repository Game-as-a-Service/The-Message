package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type GameService struct {
	GameRepo      repository.GameRepository
	PlayerService PlayerService
	DeckRepo      repository.DeckRepository
	CardService   CardService
}

type GameServiceOptions struct {
	GameRepo      repository.GameRepository
	PlayerService PlayerService
	DeckRepo      repository.DeckRepository
	CardService   CardService
}

func NewGameService(opts *GameServiceOptions) GameService {
	return GameService{
		GameRepo:      opts.GameRepo,
		PlayerService: opts.PlayerService,
		DeckRepo:      opts.DeckRepo,
		CardService:   opts.CardService,
	}
}

func (g *GameService) ShuffleDeck(cards []*repository.Card) []*repository.Card {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func (g *GameService) GetGameById(c context.Context, id int) (*repository.Game, error) {
	game, err := g.GameRepo.GetGameById(c, id)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *GameService) CreateGame(c context.Context, game *repository.Game) (*repository.Game, error) {
	game, err := g.GameRepo.CreateGame(c, game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *GameService) CreateDeck(c context.Context, deck *repository.Deck) (*repository.Deck, error) {
	deck, err := g.DeckRepo.CreateDeck(c, deck)
	if err != nil {
		return nil, err
	}
	return deck, nil
}

func (g *GameService) GetDecksByGameId(c context.Context, id int) ([]*repository.Deck, error) {
	decks, err := g.DeckRepo.GetDecksByGameId(c, id)
	if err != nil {
		return nil, err
	}
	return decks, nil

}

func (g *GameService) DeleteDeckFromGame(c context.Context, id int) error {
	err := g.DeckRepo.DeleteDeck(c, id)
	if err != nil {
		return err
	}
	return nil
}

func (g *GameService) DeleteGame(c context.Context, id int) error {
	err := g.GameRepo.DeleteGame(c, id)
	if err != nil {
		return err
	}
	return nil

}

func (g *GameService) InitGame(c *gin.Context) (*repository.Game, error) {
	game := new(repository.Game)
	jwtToken := "the-message" // 先亂寫Token
	jwtBytes := []byte(jwtToken)
	hash := sha256.Sum256(jwtBytes)
	hashString := hex.EncodeToString(hash[:])
	game.Token = hashString

	game, err := g.CreateGame(c, game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *GameService) InitDeck(c *gin.Context, game *repository.Game) error {
	cards, err := g.CardService.GetCards(c)
	if err != nil {
		return err
	}

	cards = g.ShuffleDeck(cards)

	var deck []*repository.Deck

	for _, card := range cards {
		card, err := g.DeckRepo.CreateDeck(c, &repository.Deck{
			GameId: game.Id,
			CardId: card.Id,
		})
		if err != nil {
			return err
		}
		deck = append(deck, card)
	}

	return nil
}

func (g *GameService) DrawCard(c *gin.Context, game *repository.Game, player *repository.Player, drawCards []*repository.Deck, count int) error {
	for i := 0; i < count; i++ {
		//card := new(repository.PlayerCard)
		//card.GameId = game.Id
		//card.PlayerId = player.Id
		//card.CardId = drawCards[i].CardId
		//card.Type = "hand"
		card := &repository.PlayerCard{
			GameId:   game.Id,
			PlayerId: player.Id,
			CardId:   drawCards[i].CardId,
			Type:     "hand",
		}
		err := g.PlayerService.CreatePlayerCard(c, card)
		if err != nil {
			return err
		}
		// delete deck
		err = g.DeleteDeckFromGame(c, drawCards[i].Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GameService) DrawCardsForPlayers(c *gin.Context, game *repository.Game) error {
	players, err := g.PlayerService.GetPlayersByGameId(c, game.Id)
	if err != nil {
		return err
	}
	for _, player := range players {
		drawCards, _ := g.GetDecksByGameId(c, game.Id)
		err2 := g.DrawCard(c, game, player, drawCards, 3)
		if err2 != nil {
			return err2
		}

	}
	return nil
}
