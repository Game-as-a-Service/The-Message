package service

import (
	"context"
	"errors"
	"math/rand"

	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/service/request"
)

type PlayerService struct {
	PlayerRepo       repository.PlayerRepository
	PlayerCardRepo   repository.PlayerCardRepository
	GameRepo         repository.GameRepository
	GameServ         *GameService
	GameProgressRepo repository.GameProgressesRepository
}

type PlayerServiceOptions struct {
	PlayerRepo       repository.PlayerRepository
	PlayerCardRepo   repository.PlayerCardRepository
	GameRepo         repository.GameRepository
	GameServ         *GameService
	GameProgressRepo repository.GameProgressesRepository
}

func NewPlayerService(opts *PlayerServiceOptions) PlayerService {
	return PlayerService{
		PlayerRepo:       opts.PlayerRepo,
		PlayerCardRepo:   opts.PlayerCardRepo,
		GameRepo:         opts.GameRepo,
		GameServ:         opts.GameServ,
		GameProgressRepo: opts.GameProgressRepo,
	}
}

func (p *PlayerService) InitPlayers(c context.Context, game *repository.Game, req request.CreateGameRequest) error {
	identityCards := p.InitIdentityCards(len(req.Players))
	for i, reqPlayer := range req.Players {
		_, err := p.CreatePlayer(c, &repository.Player{
			Name:         reqPlayer.Name,
			GameID:       game.ID,
			IdentityCard: identityCards[i],
			Priority:     i + 1,
			Status:       enums.PlayerStatusAlive,
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

func (p *PlayerService) CanPlayCard(c context.Context, player *repository.Player) (bool, error) {
	if player.Game.Status == enums.GameEnd {
		return false, errors.New("遊戲已結束")
	}

	if player.Status == enums.PlayerStatusDead {
		return false, errors.New("你已死亡")
	}

	if player.Game.CurrentPlayerID != player.ID {
		return false, errors.New("尚未輪到你出牌")
	}

	return true, nil
}

func (p *PlayerService) CheckPlayerCardExist(c context.Context, playerId uint, cardId uint) (bool, error) {
	exist, err := p.PlayerCardRepo.ExistPlayerCardByPlayerIdAndCardId(c, playerId, cardId)

	if err != nil {
		return false, err
	}

	return exist, nil
}

func (p *PlayerService) CreatePlayer(c context.Context, player *repository.Player) (*repository.Player, error) {
	player, err := p.PlayerRepo.CreatePlayer(c, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (p *PlayerService) CreatePlayerCard(c context.Context, card *repository.PlayerCard) error {
	_, err := p.PlayerCardRepo.CreatePlayerCard(c, card)
	if err != nil {
		return err
	}
	return nil
}

func (p *PlayerService) GetPlayerById(c context.Context, id uint) (*repository.Player, error) {
	player, err := p.PlayerRepo.GetPlayerById(c, id)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (p *PlayerService) GetPlayersByGameId(c context.Context, id uint) ([]*repository.Player, error) {
	players, err := p.PlayerRepo.GetPlayersByGameId(c, id)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (p *PlayerService) GetHandCardId(player *repository.Player, cardId uint) (*repository.PlayerCard, error) {
	for _, card := range player.PlayerCards {
		if card.CardID == cardId && card.Type == "hand" {
			return &card, nil
		}
	}
	return nil, errors.New("找不到手牌")
}

func (p *PlayerService) PlayCard(c context.Context, playerId uint, cardId uint) (*repository.Game, *repository.Card, error) {
	player, err := p.PlayerRepo.GetPlayerWithGamePlayersAndPlayerCardsCard(c, playerId)
	if err != nil {
		return nil, nil, err
	}

	result, err := p.CanPlayCard(c, player)
	if !result || err != nil {
		return nil, nil, err
	}

	handCard, err := p.GetHandCardId(player, cardId)
	if err != nil {
		return nil, nil, err
	}

	game, err := p.GameServ.NextPlayer(c, player)
	if err != nil {
		return nil, nil, err
	}

	err = p.PlayerCardRepo.DeletePlayerCard(c, handCard.ID)
	if err != nil {
		return nil, nil, err
	}

	err = p.GameRepo.UpdateGame(c, game)
	if err != nil {
		return nil, nil, err
	}

	return game, &handCard.Card, nil
}

func (p *PlayerService) TransmitIntelligenceCard(c context.Context, playerId uint, cardId uint) (bool, error) {
	player, err := p.PlayerRepo.GetPlayerWithGamePlayersAndPlayerCardsCard(c, playerId)
	if err != nil {
		return false, err
	}

	result, err := p.CanPlayCard(c, player)
	if !result || err != nil {
		return false, err
	}

	game, err := p.GameServ.NextPlayer(c, player)
	if err != nil {
		return false, err
	}

	ret, err := p.PlayerCardRepo.DeletePlayerCardByPlayerIdAndCardId(c, playerId, cardId)
	if err != nil {
		return false, err
	}

	err = p.GameRepo.UpdateGame(c, game)
	if err != nil {
		return false, err
	}

	_, err = p.GameProgressRepo.CreateGameProgress(c, &repository.GameProgresses{
		GameID:         game.ID,
		PlayerID:       playerId,
		CardID:         cardId,
		Action:         enums.TransmitIntelligence,
		TargetPlayerID: game.CurrentPlayerID,
	})

	if err != nil {
		return false, err
	}

	return ret, nil
}

func (p *PlayerService) AcceptCard(c context.Context, playerId uint, accept bool) (bool, error) {
	player, err := p.PlayerRepo.GetPlayerWithGamePlayersAndPlayerCardsCard(c, playerId)
	if err != nil {
		return false, err
	}

	result, err := p.CanPlayCard(c, player)
	if !result || err != nil {
		return false, err
	}

	game, err := p.GameServ.NextPlayer(c, player)
	if err != nil {
		return false, err
	}

	gameId := game.ID
	gameProgress, err := p.GameProgressRepo.GetGameProgresses(c, playerId, gameId)
	if err != nil {
		return false, err
	}
	cardId := gameProgress.CardID
	// assume the type is SecretTelegram
	res := accept
	if accept {
		_, err := p.PlayerCardRepo.CreatePlayerCard(c, &repository.PlayerCard{
			PlayerID: playerId,
			CardID:   cardId,
			Type:     "intelligence",
		})
		if err != nil {
			return false, err
		}
		p.GameServ.UpdateStatus(c, game, enums.ActionCardStage)

	} else {
		_, err := p.GameProgressRepo.UpdateGameProgress(c, gameProgress, game.CurrentPlayerID)
		if err != nil {
			return false, err
		}

		err = p.GameRepo.UpdateGame(c, game)
		if err != nil {
			return false, err
		}
	}

	return res, nil
}

func (p *PlayerService) CheckWin(c context.Context, playerId uint) (*repository.Player, error) {
	player, err := p.PlayerRepo.GetPlayerWithGamePlayersAndPlayerCardsCard(c, playerId)
	if err != nil {
		return nil, err
	}

	win := 0
	var winPlayer *repository.Player
	for _, player := range player.Game.Players {
		win = 0
		for _, card := range player.PlayerCards {
			if card.Type == enums.Intelligence && player.IdentityCard == enums.MilitaryAgency && card.Card.Color == enums.Red {
				win++
				if win == 3 {
					winPlayer = &player
					break
				}
			}

			if card.Type == enums.Intelligence && player.IdentityCard == enums.UndercoverFront && card.Card.Color == enums.Blue {
				win++
				if win == 3 {
					winPlayer = &player
					break
				}
			}

			if card.Type == enums.Intelligence && player.IdentityCard == enums.MilitaryAgency && card.Card.Color == enums.Red || card.Card.Color == enums.Blue {
				win++
				if win == 5 {
					winPlayer = &player
					break
				}
			}
		}
	}
	return winPlayer, nil
}
