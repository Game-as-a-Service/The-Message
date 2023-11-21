package request

import (
	"github.com/Game-as-a-Service/The-Message/service/repository"
)

type PlayerCardsResponse struct {
	ID    string `json:"id"`
	Cards []repository.PlayerCard
}
