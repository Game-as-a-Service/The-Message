package service

import (
	"context"
	"math/rand"

	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/service/request"
)

type PlayerService struct {
	PlayerRepo     repository.PlayerRepository
	PlayerCardRepo repository.PlayerCardRepository
	// GameRepo       repository.GameRepository
	// CardRepo       repository.CardRepository
}

type PlayerServiceOptions struct {
	PlayerRepo     repository.PlayerRepository
	PlayerCardRepo repository.PlayerCardRepository
	// GameRepo       repository.GameRepository
	// CardRepo       repository.CardRepository
}

func NewPlayerService(opts *PlayerServiceOptions) PlayerService {
	return PlayerService{
		PlayerRepo:     opts.PlayerRepo,
		PlayerCardRepo: opts.PlayerCardRepo,
		// 	GameRepo:       opts.GameRepo,
		// 	CardRepo:       opts.CardRepo,
	}
}

func (p *PlayerService) InitPlayers(c context.Context, game *repository.Game, req request.CreateGameRequest) error {
	identityCards := p.InitIdentityCards(len(req.Players))
	for i, reqPlayer := range req.Players {
		_, err := p.CreatePlayer(c, &repository.Player{
			Name:         reqPlayer.Name,
			GameId:       game.Id,
			IdentityCard: identityCards[i],
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PlayerService) InitIdentityCards(playersCount int) []string {
	identityCards := make([]string, playersCount)

	if playersCount == 3 {
		identityCards[0] = enums.UndercoverFront
		identityCards[1] = enums.MilitaryAgency
		identityCards[2] = enums.Bystander
	}
	identityCards = p.ShuffleIdentityCards(identityCards)
	return identityCards
}

func (p *PlayerService) ShuffleIdentityCards(cards []string) []string {
	shuffledCards := make([]string, len(cards))
	for i, j := range rand.Perm(len(cards)) {
		shuffledCards[i] = cards[j]
	}
	return shuffledCards
}

func (p *PlayerService) CreatePlayer(c context.Context, player *repository.Player) (*repository.Player, error) {
	player, err := p.PlayerRepo.CreatePlayer(c, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (p *PlayerService) GetPlayersByGameId(c context.Context, id int) ([]*repository.Player, error) {
	players, err := p.PlayerRepo.GetPlayersByGameId(c, id)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (p *PlayerService) CreatePlayerCard(c context.Context, card *repository.PlayerCard) error {
	_, err := p.PlayerCardRepo.CreatePlayerCard(c, card)
	if err != nil {
		return err
	}
	return nil
}
