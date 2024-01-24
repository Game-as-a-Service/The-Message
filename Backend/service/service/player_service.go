package service

import (
	"context"
	"math/rand"

	"errors"

	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/gin-gonic/gin"
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
			GameId:       game.Id,
			IdentityCard: identityCards[i],
			OrderNumber:  i + 1,
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

	if player.Game.CurrentPlayerId != player.Id {
		return false, errors.New("尚未輪到你出牌")
	}

	return true, nil
}

func (p *PlayerService) CheckPlayerCardExist(c context.Context, playerId int, gameId int, cardId int) (bool, error) {
	exist, err := p.PlayerCardRepo.ExistPlayerCardByPlayerIdAndCardId(c, playerId, gameId, cardId)

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

func (p *PlayerService) GetPlayerById(c context.Context, id int) (*repository.Player, error) {
	player, err := p.PlayerRepo.GetPlayer(c, id)
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

func (p *PlayerService) GetHandCardId(player *repository.Player, cardId int) (*repository.PlayerCard, error) {
	for _, card := range player.PlayerCards {
		if card.CardId == cardId && card.Type == "hand" {
			return &card, nil
		}
	}
	return nil, errors.New("找不到手牌")
}

func (p *PlayerService) PlayCard(c *gin.Context, playerId int, cardId int) (*repository.Game, *repository.Card, error) {
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

	err = p.PlayerCardRepo.DeletePlayerCard(c, handCard.Id)
	if err != nil {
		return nil, nil, err
	}

	err = p.GameRepo.UpdateGame(c, game)
	if err != nil {
		return nil, nil, err
	}

	return game, &handCard.Card, nil
}

func (p *PlayerService) TransmitIntelligenceCard(c *gin.Context, playerId int, gameId int, cardId int) (bool, error) {
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

	ret, err := p.PlayerCardRepo.DeletePlayerCardByPlayerIdAndCardId(c, playerId, gameId, cardId)
	if err != nil {
		return false, err
	}

	err = p.GameRepo.UpdateGame(c, game)
	if err != nil {
		return false, err
	}

	_, err = p.GameProgressRepo.CreateGameProgress(c, &repository.GameProgresses{
		GameId:         game.Id,
		PlayerId:       playerId,
		CardId:         cardId,
		Action:         enums.TransmitIntelligence,
		TargetPlayerId: game.CurrentPlayerId,
	})

	if err != nil {
		return false, err
	}

	return ret, nil
}

func (p *PlayerService) AcceptCard(c *gin.Context, playerId int) (bool, error) {
	player, err := p.PlayerRepo.GetPlayerWithGamePlayersAndPlayerCardsCard(c, playerId)
	if err != nil {
		return false, err
	}

	result, err := p.CanPlayCard(c, player)
	if !result || err != nil {
		return false, err
	}

	// get target player id from game process
	gameId := player.Game.Id
	gameProgress, err := p.GameProgressRepo.GetGameProgresses(c, playerId, gameId)

	cardId := gameProgress.CardId
	// assume the type is SecretTelegram

	// decide whether accpet or not

	return true, nil
}
