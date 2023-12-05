//go:build migrate

package main

import (
	"github.com/Game-as-a-Service/The-Message/config"
	"github.com/Game-as-a-Service/The-Message/database/seeders"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := config.NewDatabase()
	seeders.Run(db)
}
