package seeders

import (
	"context"

	"github.com/Game-as-a-Service/The-Message/enums"
	"github.com/Game-as-a-Service/The-Message/service/repository"
	"github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"gorm.io/gorm"
)

var actionColors = map[string]map[string]int{
	enums.LockOn:      {"紅": 5, "藍": 5, "黑": 4},
	enums.LureAway:    {"紅": 2, "藍": 2, "黑": 4},
	enums.Intercept:   {"紅": 1, "藍": 1, "黑": 6},
	enums.Diversion:   {"紅": 2, "藍": 2, "黑": 2},
	enums.Decipher:    {"紅": 3, "藍": 3, "黑": 1},
	enums.Burn:        {"紅": 1, "藍": 1, "黑": 4},
	enums.SeeThrough:  {"紅": 3, "藍": 3, "黑": 6},
	enums.Probe:       {"紅": 6, "藍": 6, "黑": 6},
	enums.BlurOfTruth: {"紅": 1, "藍": 1},
}

func SeederCards(db *gorm.DB) {
	for actionType, colors := range actionColors {
		for color, count := range colors {
			for i := 0; i < count; i++ {
				mysql.NewCardRepository(db).CreateCard(context.TODO(), &repository.Card{
					Name:             actionType,
					Color:            color,
					IntelligenceType: enums.ToIntelligenceType(actionType),
				})
			}
		}
	}
}
