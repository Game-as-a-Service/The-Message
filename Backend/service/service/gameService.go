package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/Game-as-a-Service/The-Message/enums"
	domain "github.com/Game-as-a-Service/The-Message/service/repository"
	repository "github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/gin-gonic/gin"
	"math/rand"
)

type GameService struct {
	GameRepo   repository.GameRepository
	PlayerRepo repository.PlayerRepository
}

func NewGameService(gameRepo repository.GameRepository, playerRepo repository.PlayerRepository) *GameService {
	return &GameService{
		GameRepo:   gameRepo,
		PlayerRepo: playerRepo,
	}
}

func (g *GameService) StartGame(cxt *gin.Context, req request.CreateGameRequest) (*domain.Game, error) {
	game := new(domain.Game)
	jwtToken := "the-message" // 先亂寫Token
	jwtBytes := []byte(jwtToken)
	hash := sha256.Sum256(jwtBytes)
	hashString := hex.EncodeToString(hash[:])
	game.Token = hashString

	game, err := g.GameRepo.CreateGame(cxt, game)
	if err != nil {
		return nil, err
	}

	// 建立身份牌牌堆
	identityCards := initIdentityCards(len(req.Players))

	for i, reqPlayer := range req.Players {
		player := new(domain.Player)
		player.Name = reqPlayer.Name
		player.GameId = game.Id
		player.IdentityCard = identityCards[i]
		player, err = g.PlayerRepo.CreatePlayer(cxt, player)
		if err != nil {
			return nil, err
		}
	}

	return game, err
}

func initIdentityCards(count int) []string {
	identityCards := make([]string, count)
	// if count == 3 had 1 UndercoverFront, 1 MilitaryIntel, 1 Bystander
	if count == 3 {
		identityCards[0] = enums.UndercoverFront
		identityCards[1] = enums.MilitaryIntel
		identityCards[2] = enums.Bystander
	}
	identityCards = shuffle(identityCards)
	return identityCards
}

func shuffle(cards []string) []string {
	shuffledCards := make([]string, len(cards))
	for i, j := range rand.Perm(len(cards)) {
		shuffledCards[i] = cards[j]
	}
	return shuffledCards
}
