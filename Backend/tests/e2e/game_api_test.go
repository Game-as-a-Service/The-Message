package e2e

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/Game-as-a-Service/The-Message/database/seeders"
	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/service/request"
	"github.com/Game-as-a-Service/The-Message/utils"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func (suite *IntegrationTestSuite) TestStartGameE2E() {
	// Set the cards
	seeders.OnlyCardsRun(suite.db)

	api := "/api/v1/games"

	suite.T().Run("it can validate room id and players", func(t *testing.T) {
		invalidRoomReq := utils.RequestToJsonBody(struct {
			Players []request.PlayerInfo `json:"players"`
		}{
			Players: []request.PlayerInfo{},
		})

		invalidPlayerReq := utils.RequestToJsonBody(struct {
			RoomID string `json:"room_id"`
		}{
			RoomID: faker.UUIDDigit(),
		})

		invalidRoomRes := suite.requestJson(api, invalidRoomReq, http.MethodPost)
		invalidPlayerRes := suite.requestJson(api, invalidPlayerReq, http.MethodPost)

		// Convert response body from json to map
		resBodyAsByteArray, _ := io.ReadAll(invalidRoomRes.Body)
		resBody := make(map[string]interface{})
		_ = json.Unmarshal(resBodyAsByteArray, &resBody)

		assert.Equal(t, http.StatusBadRequest, invalidRoomRes.StatusCode)
		assert.Equal(t, http.StatusBadRequest, invalidPlayerRes.StatusCode)
	})

	suite.T().Run("it can fail when player count less than 3", func(t *testing.T) {
		num := rand.Intn(2) + 1

		var players []request.PlayerInfo
		for i := 0; i < num; i++ {
			player := request.PlayerInfo{
				ID:   faker.UUIDDigit(),
				Name: faker.FirstName(),
			}
			players = append(players, player)
		}

		req := utils.RequestToJsonBody(request.CreateGameRequest{
			RoomID:  faker.UUIDDigit(),
			Players: players,
		})

		res := suite.requestJson(api, req, http.MethodPost)

		// Convert response body from json to map
		resBody := utils.JsonBodyToMap(res)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, "Players must be at least 3 and at most 9", resBody["message"])
	})

	suite.T().Run("it can fail when player count more than 9", func(t *testing.T) {
		num := 10

		var players []request.PlayerInfo
		for i := 0; i < num; i++ {
			player := request.PlayerInfo{
				ID:   faker.UUIDDigit(),
				Name: faker.FirstName(),
			}
			players = append(players, player)
		}

		req := utils.RequestToJsonBody(request.CreateGameRequest{
			RoomID:  faker.UUIDDigit(),
			Players: players,
		})

		res := suite.requestJson(api, req, http.MethodPost)

		// Convert response body from json to map
		resBody := utils.JsonBodyToMap(res)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, "Players must be at least 3 and at most 9", resBody["message"])
	})

	suite.T().Run("it can success when player count between 3 and 9", func(t *testing.T) {
		// Set up environment variables
		frontendURL := os.Getenv("APP_FRONTEND_URL")
		version := os.Getenv("APP_VERSION")
		suite.T().Setenv("APP_FRONTEND_URL", frontendURL)
		suite.T().Setenv("APP_VERSION", version)

		num := rand.Intn(7) + 3
		roomID := faker.UUIDDigit()
		url := frontendURL + version + "/games/" + strconv.Itoa(1)

		var players []request.PlayerInfo
		for i := 0; i < num; i++ {
			player := request.PlayerInfo{
				ID:   faker.UUIDDigit(),
				Name: faker.FirstName(),
			}
			players = append(players, player)
		}

		req := utils.RequestToJsonBody(request.CreateGameRequest{
			RoomID:  roomID,
			Players: players,
		})

		res := suite.requestJson(api, req, http.MethodPost)

		// Convert response body from json to map
		resBody := utils.JsonBodyToMap(res)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, url, resBody["url"])

		// Assert Game and Players
		game := &repository.Game{}
		_ = suite.db.First(&game, "id = ?", 1)

		assert.Equal(t, roomID, game.RoomID)
		assert.Equal(t, enums.ActionCardStage, game.Status)

		gamePlayers := &[]repository.Player{}
		_ = suite.db.Find(&gamePlayers, "game_id = ?", game.ID)

		assert.Equal(t, num, len(*gamePlayers))
	})
}
