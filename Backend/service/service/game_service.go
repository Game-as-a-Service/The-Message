package service

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"log"
	"math/rand"
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
		Token:  token,
		Status: enums.GameStart,
	})
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (g *GameService) AssignIdentityCards(c context.Context, playCount int) ([]string, error) {

	if playCount < 3 || playCount > 9 {
		return nil, errors.New("player count must be between 3 and 9")
	}

	var result []string

	if playCount <= 6 {
		y := playCount / 3
		for i := 0; i < y; i++ {
			result = append(result, enums.UndercoverFront)
			result = append(result, enums.MilitaryAgency)
			result = append(result, enums.Bystander)
		}

		x := playCount % 3
		if x == 1 {
			result = append(result, enums.Bystander)
		} else if x == 2 {
			result = append(result, enums.UndercoverFront)
			result = append(result, enums.MilitaryAgency)
		}

	} else if playCount > 6 {
		y := 3
		for i := 0; i < y; i++ {
			result = append(result, enums.UndercoverFront)
			result = append(result, enums.MilitaryAgency)
		}
		x := playCount % 3
		if x == 0 {
			x = 3
		}
		for i := 0; i < x; i++ {
			result = append(result, enums.Bystander)
		}
	}

	log.Printf("result: %+v", result)
	log.Printf("count: %+v", len(result))

	// 創建一個新的隨機數生成器實例
	//src := rand.NewSource(time.Now().UnixNano())
	//r := rand.New(src)

	// 使用創建的隨機數生成器實例來執行洗牌
	//r.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })

	return result, nil
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
	game, err := g.GameRepo.GetGameWithPlayers(c, id)
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

func (g *GameService) UpdateCurrentPlayer(c context.Context, game *repository.Game, playerId int) {
	game.CurrentPlayerId = playerId
	err := g.GameRepo.UpdateGame(c, game)
	if err != nil {
		panic(err)
	}
}

func (g *GameService) NextPlayer(c context.Context, player *repository.Player) (*repository.Game, error) {
	players := player.Game.Players

	currentPlayerId := player.Id

	var currentPlayerIndex int
	for index, gPlayer := range players {
		if gPlayer.Id == currentPlayerId {
			currentPlayerIndex = index
			break
		}
	}

	if currentPlayerIndex+1 >= len(players) {
		player.Game.CurrentPlayerId = players[0].Id
		player.Game.Status = enums.TransmitIntelligenceStage
	} else {
		player.Game.CurrentPlayerId = players[currentPlayerIndex+1].Id
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
