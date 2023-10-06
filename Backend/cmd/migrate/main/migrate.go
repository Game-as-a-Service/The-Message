package main

import (
	"github.com/Game-as-a-Service/The-Message/database"
	domain "github.com/Game-as-a-Service/The-Message/service/repository"
	repository "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"gorm.io/gorm"
)

func main() {
	db := database.InitDB()
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
