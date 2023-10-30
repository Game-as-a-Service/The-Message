package service

import (
	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
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

func (g *GameService) InitialDeck(gameId int, cards []*repository.Card) []*repository.Card {
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

func (g *GameService) GetGameById(c *gin.Context, id int) (*repository.Game, error) {
	game, err := g.GameRepo.GetGameById(c, id)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *GameService) CreateGame(c *gin.Context, game *repository.Game) (*repository.Game, error) {
	game, err := g.GameRepo.CreateGame(c, game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *GameService) CreatePlayer(c *gin.Context, player *repository.Player) (*repository.Player, error) {
	player, err := g.PlayerRepo.CreatePlayer(c, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (g *GameService) GetCards(c *gin.Context) ([]*repository.Card, error) {
	cards, err := g.CardRepo.GetCards(c)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (g *GameService) CreateDeck(c *gin.Context, deck *repository.Deck) (*repository.Deck, error) {
	deck, err := g.DeckRepo.CreateDeck(c, deck)
	if err != nil {
		return nil, err
	}
	return deck, nil
}

func (g *GameService) GetPlayersByGameId(c *gin.Context, id int) ([]*repository.Player, error) {
	players, err := g.PlayerRepo.GetPlayersByGameId(c, id)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (g *GameService) GetDecksByGameId(c *gin.Context, id int) ([]*repository.Deck, error) {
	decks, err := g.DeckRepo.GetDecksByGameId(c, id)
	if err != nil {
		return nil, err
	}
	return decks, nil

}

func (g *GameService) CreatePlayerCard(c *gin.Context, cards *repository.PlayerCard) (*repository.PlayerCard, error) {
	cards, err := g.PlayerCardRepo.CreatePlayerCard(c, cards)
	if err != nil {
		return nil, err
	}
	return cards, nil

}

func (g *GameService) DeleteDeck(c *gin.Context, id int) error {
	err := g.DeckRepo.DeleteDeck(c, id)
	if err != nil {
		return err
	}
	return nil
}

func (g *GameService) DeleteGame(c *gin.Context, id int) error {
	err := g.GameRepo.DeleteGame(c, id)
	if err != nil {
		return err
	}
	return nil

}
