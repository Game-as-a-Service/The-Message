package request

type CreateGameRequest struct {
	Players []PlayerInfo `json:"players"`
}

type PlayerInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateGameResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type PlayCardResponse struct {
	Result bool `json:"result"`
}

type PlayCardRequest struct {
	CardID int `json:"card_id"`
}

type AcceptCardRequest struct {
	Accept bool `json:"accept"`
}
