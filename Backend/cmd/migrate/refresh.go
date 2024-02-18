//go:build migrate

package main

import (
	"github.com/Game-as-a-Service/The-Message/config"
	"github.com/Game-as-a-Service/The-Message/database/seeders"
)

func main() {
	config.RunRefresh()
	db := config.NewDatabase()
	seeders.Run(db)
}
