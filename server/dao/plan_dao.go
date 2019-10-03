package dao

import (
	"log"

	"bemobi-hire/server/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PlansDAO Contains all database manipulation relative to Plans
type PlansDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// PLANCOLLECTION Collection that have plans.
	PLANCOLLECTION = "carrier_plans"
)

// Connect Conecta o DAO ao nosso banco de dados.
func (m *PlansDAO) Connect() *PlansDAO {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	return m
}

// ListPlansByCarrier list stored plans by carrier
func (m *PlansDAO) ListPlansByCarrier(carrier string) ([]models.Plan, error) {
	var plans []models.Plan = ([]models.Plan{})
	err := db.C(PLANCOLLECTION).Find(bson.M{
		"plan_carrier": carrier,
	}).All(&plans)
	return plans, err
}

// InserPlan insert a new plan in database creating one.
func (m *PlansDAO) InserPlan(plan models.Plan) error {
	err := db.C(PLANCOLLECTION).Insert(&plan)
	return err
}

// GetPlanByCarrierAndSku Returns an plan details by sku and carrier
func (m *PlansDAO) GetPlanByCarrierAndSku(Carrier string, Sku string) (models.DatabasePlan, error) {
	var plan models.DatabasePlan

	err := db.C(PLANCOLLECTION).Find(bson.M{
		"plan_carrier": Carrier,
		"plan_sku":     Sku,
	}).One(&plan)

	return plan, err
}

// DeleteByID delete a plan by id or not
func (m *PlansDAO) DeleteByID(id string) error {
	err := db.C(PLANCOLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

// DeleteBySku Delete a plan by Sku and return count removed
func (m *PlansDAO) DeleteBySku(sku string) (int, error) {
	info, err := db.C(PLANCOLLECTION).RemoveAll(bson.M{"plan_sku": sku})
	return info.Removed, err
}

// UpdateByID Updates a plan by document id
func (m *PlansDAO) UpdateByID(id string, plan models.DatabasePlan) error {
	err := db.C(PLANCOLLECTION).UpdateId(bson.ObjectIdHex(id), &plan)
	return err
}
