package seeders

import (
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	SeederCards(db)
	SeederGameWithPlayers(db)
}

func OnlyCardsRun(db *gorm.DB) {
	SeederCards(db)
}
