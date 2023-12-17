package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/Game-as-a-Service/The-Message/service/repository"
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

func (g *GameService) InitGame(c context.Context) (*repository.Game, error) {
	token, err := g.GenerateSecureToken(256)
	if err != nil {
		return nil, err
	}

	game, err := g.CreateGame(c, &repository.Game{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	// sent sse

	return game, nil
}

func (g *GameService) InitDeck(c context.Context, game *repository.Game) error {
	err := g.DeckService.InitDeck(c, game)
	if err != nil {
		return err
	}
	return nil
}

func (g *GameService) DrawCard(c context.Context, game *repository.Game, player *repository.Player, drawCards []*repository.Deck, count int) error {
	for i := 0; i < count; i++ {
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
		err = g.DeckService.DeleteDeckFromGame(c, drawCards[i].Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GameService) DrawCardsForAllPlayers(c context.Context, game *repository.Game) error {
	players, err := g.PlayerService.GetPlayersByGameId(c, game.Id)
	if err != nil {
		return err
	}
	for _, player := range players {
		drawCards, _ := g.DeckService.GetDecksByGameId(c, game.Id)
		err2 := g.DrawCard(c, game, player, drawCards, 3)
		if err2 != nil {
			return err2
		}

	}
	return nil
}

func (g *GameService) GenerateSecureToken(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (g *GameService) CreateGame(c context.Context, game *repository.Game) (*repository.Game, error) {
	game, err := g.GameRepo.CreateGame(c, game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *GameService) GetGameById(c context.Context, id int) (*repository.Game, error) {
	game, err := g.GameRepo.GetGameById(c, id)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *GameService) DeleteGame(c context.Context, id int) error {
	err := g.GameRepo.DeleteGame(c, id)
	if err != nil {
		return err
	}
	return nil

}
