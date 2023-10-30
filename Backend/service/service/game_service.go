package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type GameService struct {
	GameRepo       repository.GameRepository
	PlayerRepo     repository.PlayerRepository
	CardRepo       repository.CardRepository
	DeckRepo       repository.DeckRepository
	PlayerCardRepo repository.PlayerCardRepository
}

type GameServiceOptions struct {
	GameRepo       repository.GameRepository
	PlayerRepo     repository.PlayerRepository
	CardRepo       repository.CardRepository
	DeckRepo       repository.DeckRepository
	PlayerCardRepo repository.PlayerCardRepository
}

func NewGameService(opts *GameServiceOptions) GameService {
	return GameService{
		GameRepo:       opts.GameRepo,
		PlayerRepo:     opts.PlayerRepo,
		CardRepo:       opts.CardRepo,
		DeckRepo:       opts.DeckRepo,
		PlayerCardRepo: opts.PlayerCardRepo,
	}
}

func (g *GameService) ShuffleDeck(cards []*repository.Card) []*repository.Card {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func (g *GameService) InitIdentityCards(count int) []string {
	identityCards := make([]string, count)

	if count == 3 {
		identityCards[0] = enums.UndercoverFront
		identityCards[1] = enums.MilitaryAgency
		identityCards[2] = enums.Bystander
	}
	identityCards = g.shuffle(identityCards)
	return identityCards
}

func (g *GameService) shuffle(cards []string) []string {
	shuffledCards := make([]string, len(cards))
	for i, j := range rand.Perm(len(cards)) {
		shuffledCards[i] = cards[j]
	}
	return shuffledCards
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

func (g *GameService) CreatePlayer(c context.Context, player *repository.Player) (*repository.Player, error) {
	player, err := g.PlayerRepo.CreatePlayer(c, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (g *GameService) GetCards(c context.Context) ([]*repository.Card, error) {
	cards, err := g.CardRepo.GetCards(c)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (g *GameService) CreateDeck(c context.Context, deck *repository.Deck) (*repository.Deck, error) {
	deck, err := g.DeckRepo.CreateDeck(c, deck)
	if err != nil {
		return nil, err
	}
	return deck, nil
}

func (g *GameService) GetPlayersByGameId(c context.Context, id int) ([]*repository.Player, error) {
	players, err := g.PlayerRepo.GetPlayersByGameId(c, id)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (g *GameService) GetDecksByGameId(c context.Context, id int) ([]*repository.Deck, error) {
	decks, err := g.DeckRepo.GetDecksByGameId(c, id)
	if err != nil {
		return nil, err
	}
	return decks, nil

}

func (g *GameService) CreatePlayerCard(c context.Context, cards *repository.PlayerCard) (*repository.PlayerCard, error) {
	cards, err := g.PlayerCardRepo.CreatePlayerCard(c, cards)
	if err != nil {
		return nil, err
	}
	return cards, nil

}

func (g *GameService) DeleteDeck(c context.Context, id int) error {
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

func (g *GameService) InitPlayers(c *gin.Context, game *repository.Game, req request.CreateGameRequest) error {
	identityCards := g.InitIdentityCards(len(req.Players))
	for i, reqPlayer := range req.Players {
		player := new(repository.Player)
		player.Name = reqPlayer.Name
		player.GameId = game.Id
		player.IdentityCard = identityCards[i]
		_, err := g.CreatePlayer(c, player)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GameService) InitDeck(c *gin.Context, game *repository.Game) error {
	cards, err := g.GetCards(c)
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
		playerCards := new(repository.PlayerCard)
		playerCards.GameId = game.Id
		playerCards.PlayerId = player.Id
		playerCards.CardId = drawCards[i].CardId
		playerCards.Type = "hand"
		_, err := g.CreatePlayerCard(c, playerCards)
		if err != nil {
			return err
		}
		// delete deck
		err = g.DeleteDeck(c, drawCards[i].Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GameService) DrawCardsForPlayers(c *gin.Context, game *repository.Game) error {
	players, err := g.GetPlayersByGameId(c, game.Id)
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
