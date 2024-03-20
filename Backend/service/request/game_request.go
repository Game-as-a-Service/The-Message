package request

type CreateGameRequest struct {
	RoomID  string       `json:"roomId" binding:"required"`
	Players []PlayerInfo `json:"players" binding:"required"`
}

type GetGameRequest struct {
	GameID uint `json:"gameId" binding:"required"`
}

type PlayerInfo struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"nickname" binding:"required"`
}

type CreateGameResponse struct {
	URL string `json:"url"`
}

type PlayCardResponse struct {
	Result bool `json:"result"`
}

type PlayCardRequest struct {
	CardID uint `json:"card_id"`
}

type AcceptCardRequest struct {
	Accept bool `json:"accept"`
}
