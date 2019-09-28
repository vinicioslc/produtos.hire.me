package migrations

import (
	"bemobi-hire/server/models"
	"fmt"

	migrate "github.com/eminetto/mongo-migrate"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

var (
	plansColl = "carrier_plans"
)

/// Baseado nos exemplos de https://github.com/xakep666/mongo-migrate
func init() {
	fmt.Sprintln("ERROR ----")
	migrate.Register(func(db *mgo.Database) error { // uP
		return db.C(plansColl).Insert(
			models.Plan{
				ID:          bson.NewObjectId(),
				PlanCarrier: "vivo",
				PlanSKU:     "WEB_VIVOTURBO15DIAS1499",
				PlanAdvantages: []models.PlanAdvantage{
					{
						"Ganhe 1GB de bônus, ",
						"válido por 7 dias, ao acumular R$35 em recarga",
					},
					{
						"Minutos ilimitados ",
						"em ligações locais para celulares de outras operadoras",
					},
				},
				PlanHighlights: []models.PlanHighlights{
					{
						"whattsicon.png",
						"WhatsApp Ilimitado",
						"para mensagens, vídeos e fotos",
					},
				},
				PlanTitle:       "3GB",
				PlanDetails:     "de internet 4.5G",
				PlanSmallWords:  "com renovação automática",
				PlanLimitMbytes: 2000,
				PlanLimitDays:   7,
				PlanPrice:       15.990,
				PlanMoreDetails: models.PlanDetails{
					Franchise:      "",
					AdditionalInfo: "",
					Regulation:     "",
				},
			},
		)
	}, func(db *mgo.Database) error { // dOWN
		return db.C(plansColl).DropCollection()
	})
}
