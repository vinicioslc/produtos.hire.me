package models

import "gopkg.in/mgo.v2/bson"

// Plan descreve um plano de assinatura para os usuários
type Plan struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	PlanType       string        `bson:"plan_type" json:"plan_type"`
	PlanDetails    string        `bson:"plan_details" json:"plan_details"`
	PlanAdvantages []string      `bson:"plan_advantages" json:"plan_advantages"`
	PlanPrice      float64       `bson:"plan_price" json:"plan_price"`
}
