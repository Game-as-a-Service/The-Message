//go:build migrate

package main

import (
	"github.com/Game-as-a-Service/The-Message/config"
	"github.com/Game-as-a-Service/The-Message/domain"
	repository "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func main() {
	db := config.InitDB()
	MigrationMysql(db)
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
}

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Game{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&domain.Player{})
	if err != nil {
		return
	}
}
