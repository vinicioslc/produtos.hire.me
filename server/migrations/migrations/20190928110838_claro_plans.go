package migrations

import (
	"bemobi-hire/server/models"

	migrate "github.com/eminetto/mongo-migrate"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

/// Baseado nos exemplos de https://github.com/xakep666/mongo-migrate
func init() {
	carrier := "claro"
	migrate.Register(func(db *mgo.Database) error { // uP
		return db.C(plansColl).Insert(
			models.Plan{
				ID:          bson.NewObjectId(),
				PlanCarrier: carrier,
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
				PlanTitle:       "3GB",
				PlanDetails:     "de internet 4.5G",
				PlanSmallWords:  "com renovação automática",
				PlanLimitMbytes: 2000,
				PlanLimitDays:   15,
				PlanPrice:       14.990,
				PlanMoreDetails: models.PlanDetails{
					Franchise:      ``,
					AdditionalInfo: "",
					Regulation:     "",
				},
			},
			models.Plan{
				ID:          bson.NewObjectId(),
				PlanCarrier: carrier,
				PlanSKU:     "WEB_CLARO5GB",
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
				PlanTitle:       "5GB",
				PlanDetails:     "de internet 5G",
				PlanSmallWords:  "sem renovação automática",
				PlanLimitMbytes: 5000,
				PlanLimitDays:   7,
				PlanPrice:       8.990,
				PlanMoreDetails: models.PlanDetails{
					Franchise:      "",
					AdditionalInfo: "",
					Regulation:     "",
				},
			},
		)
	}, func(db *mgo.Database) error { // dOWN
		_, err := db.C(plansColl).RemoveAll(bson.M{
			"plan_carrier": carrier,
		})
		return err
	})
}
