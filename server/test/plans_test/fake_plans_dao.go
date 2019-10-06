package plans_test

import (
	"bemobi-hire/server/dao"
	"errors"

	"bemobi-hire/server/models"

	"gopkg.in/mgo.v2/bson"
)

var plans = []models.DatabasePlan{
	models.DatabasePlan{
		ID:          bson.NewObjectId(),
		PlanCarrier: "vivo",
		PlanSKU:     "WEB_VIVO100MB",
		PlanMoreDetails: models.PlanDetails{
			Franchise: "Franquia vivo",
		},
		PlanHighlights: []models.PlanHighlights{
			{
				Title: "+10 minutos de ligações para outras operadoras*",
			},
		},
		PlanAdvantages: []models.PlanAdvantage{
			{
				"Minutos ilimitados ",
				"em ligações locais para celulares de outras operadoras",
			},
		},
		PlanTitle:       "100 MB",
		PlanDetails:     "Por dia",
		PlanSmallWords:  "sem renovação automática",
		PlanLimitMbytes: 5000,
		PlanLimitDays:   1,
		PlanPrice:       1.990,
	},
	models.DatabasePlan{
		ID:          bson.NewObjectId(),
		PlanCarrier: "claro",
		PlanSKU:     "WEB_CLARO1GB",
		Spotted:     true,
		PlanHighlights: []models.PlanHighlights{
			{
				Title: "+1GB de bônus com renovação semanal*",
			},
		},
		PlanAdvantages: []models.PlanAdvantage{
			{
				"Ganhe 1GB de bônus, ",
				"válido por 7 dias, ao acumular R$35 em recarga",
			},
		},
		PlanTitle:       "1GB",
		PlanDetails:     "Por semana",
		PlanSmallWords:  "com renovação automática",
		PlanLimitMbytes: 1000,
		PlanLimitDays:   7,
		PlanPrice:       9.99,
		PlanMoreDetails: models.PlanDetails{
			Franchise: "Franquia Loka",
		},
	},
	models.DatabasePlan{
		ID:          bson.NewObjectId(),
		PlanCarrier: "claro",
		PlanSKU:     "WEB_CLARO5GB",
		PlanHighlights: []models.PlanHighlights{
			{
				Title: "Ligações ilimitadas*",
			},
		},
		PlanAdvantages: []models.PlanAdvantage{
			{
				"Ligações ilimitadas, ",
				"válido por 7 dias, ao acumular R$35 em recarga",
			},
			{
				"Minutos ilimitados ",
				"em ligações locais para celulares de outras operadoras",
			},
		},
		PlanTitle:       "5GB",
		PlanDetails:     "Por Semana",
		PlanSmallWords:  "com renovação automática",
		PlanLimitMbytes: 1000,
		PlanLimitDays:   7,
		PlanPrice:       14.99,
	},
}

// FakePlansDAO Contains !FAKE DATABASE! manipulation relative to Plans
type FakePlansDAO struct {
	dao.IPlansDAO
}

// Connect Conecta o DAO ao nosso banco de dados.
func (m *FakePlansDAO) Connect() dao.IPlansDAO {
	return m
}

// ListPlansByCarrier list stored plans by carrier
func (m *FakePlansDAO) ListPlansByCarrier(carrier string) ([]models.Plan, error) {

	var filtered = []models.Plan{}
	for _, element := range plans {
		if element.PlanCarrier == carrier {
			filtered = append(filtered, element.AsPlan())
		}
	}
	return filtered, nil
}

// InsertPlanIntoCarrier insert a new plan in database creating one.
func (m *FakePlansDAO) InsertPlanIntoCarrier(carrier string, plan models.DatabasePlan, result *models.DatabasePlan) error {
	return nil
}

// GetDatabasePlanByCarrierSKU Returns an plan details by sku and carrier
func (m *FakePlansDAO) GetDatabasePlanByCarrierSKU(carrier string, sku string) (models.DatabasePlan, error) {

	for _, element := range plans {
		if element.PlanCarrier == carrier && element.PlanSKU == sku {

			return element, nil
		}
	}
	return models.DatabasePlan{}, errors.New("not found")
}

// DeleteByID delete a plan by id or not
func (m *FakePlansDAO) DeleteByID(id string) error {
	return nil
}

// DeleteBySku Delete a plan by Sku and return count removed
func (m *FakePlansDAO) DeleteBySku(sku string) error {
	return nil
}

// UpdateByID Updates a plan by document id
func (m *FakePlansDAO) UpdateByID(id string, plan models.DatabasePlan) error {
	return nil
}

//UpdateBySku Updates a plan by given sku
func (m *FakePlansDAO) UpdateBySku(sku string, plan models.DatabasePlan) error {

	return nil
}
