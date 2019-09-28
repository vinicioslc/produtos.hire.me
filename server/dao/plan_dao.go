package dao

import (
	"log"

	"bemobi-hire/server/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PlansDAO Contém conexão com banco de dados
type PlansDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// PLANCOLLECTION Nome da coleção no mongodb
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

// ListPlans Lista todos os planos dado a operadora
func (m *PlansDAO) ListPlans(carrier string) ([]models.Plan, error) {
	var plans []models.Plan
	err := db.C(PLANCOLLECTION).Find(bson.M{
		"plan_carrier": carrier,
	}).All(&plans)
	return plans, err
}

// CreatePlan cria uma nova assinatura no banco
func (m *PlansDAO) CreatePlan(plan models.Plan) error {
	err := db.C(PLANCOLLECTION).Insert(&plan)
	return err
}
