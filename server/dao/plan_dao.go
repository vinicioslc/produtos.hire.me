package dao

import (
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
	// COLLECTION Nome da coleção no mongodb
	COLLECTION = "plans"
)

// ListPlans Lista todos os planos dado a operadora
func ListPlans(carrier string) ([]models.Plan, error) {
	var movies []models.Plan
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}
