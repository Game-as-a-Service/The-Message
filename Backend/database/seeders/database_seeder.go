package seeders

import (
	"github.com/Game-as-a-Service/The-Message/config"
)

func Run() {
	db := config.NewDatabase()
	SeederCards(db)
}
