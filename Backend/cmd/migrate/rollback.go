//go:build migrate

package main

import (
	"github.com/Game-as-a-Service/The-Message/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	m, err := config.NewMigration()
	if err != nil {
		panic(err)
	}

	err = m.Down()
	if err != nil {
		if err.Error() == "no change" {
		} else {
			panic(err)
		}
	}
}
