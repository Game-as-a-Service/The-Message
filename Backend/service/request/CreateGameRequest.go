package request

type CreateGameRequest struct {
	Players []PlayerInfo `json:"players"`
}

type PlayerInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
