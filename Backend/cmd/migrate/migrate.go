//go:build migrate

package main

import (
	"fmt"
	"github.com/Game-as-a-Service/The-Message/config"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

func main() {
	m, err := config.NewMigration()
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("no change")
			return
		}
		panic(err)
	}
}

func MigrationMysql(db *gorm.DB) {
	err := db.AutoMigrate(&repository.Game{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&repository.Player{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&repository.Card{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&repository.Deck{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&repository.PlayerCard{})
	if err != nil {
		return
	}
}

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(&repository.Game{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&repository.Player{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&repository.Card{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&repository.Deck{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&repository.PlayerCard{})
	if err != nil {
		return
	}
}
