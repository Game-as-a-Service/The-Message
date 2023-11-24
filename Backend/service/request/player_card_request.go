package request

type PlayerCardsResponse struct {
	ID      string   `json:"id"`
	CardIds []string `json:"cardIds"`
}
