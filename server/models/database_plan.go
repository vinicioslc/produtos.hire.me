package models

import "gopkg.in/mgo.v2/bson"

// DatabasePlan is a type used to write into database full plan data .
type DatabasePlan struct {
	ID              bson.ObjectId    `bson:"_id" json:"id"`
	PlanSKU         string           `bson:"plan_sku" json:"plan_sku"`
	Spotted         bool             `bson:"plan_spotted" json:"plan_spotted"`
	PlanCarrier     string           `bson:"plan_carrier" json:"plan_carrier"`
	PlanTitle       string           `bson:"plan_title" json:"plan_title"`
	PlanDetails     string           `bson:"plan_details" json:"plan_details"`
	PlanSmallWords  string           `bson:"plan_small_words" json:"plan_small_words"`
	PlanPrice       float64          `bson:"plan_price" json:"plan_price"`
	PlanHighlights  []PlanHighlights `bson:"plan_highlights" json:"plan_highlights"`
	PlanAdvantages  []PlanAdvantage  `bson:"plan_advantages" json:"plan_advantages"`
	PlanLimitMbytes int              `bson:"plan_limit_mbytes" json:"plan_limit_mbytes"`
	PlanLimitDays   int              `bson:"plan_limit_days" json:"plan_limit_days"`
	PlanMoreDetails PlanDetails      `bson:"plan_more_details" json:"plan_more_details"`
}

// AsPlan Convert a database plan to a Plan object
func (dbp *DatabasePlan) AsPlan() Plan {
	return (Plan{
		ID:              dbp.ID,
		PlanSKU:         dbp.PlanSKU,
		Spotted:         dbp.Spotted,
		PlanCarrier:     dbp.PlanCarrier,
		PlanTitle:       dbp.PlanTitle,
		PlanDetails:     dbp.PlanDetails,
		PlanSmallWords:  dbp.PlanSmallWords,
		PlanPrice:       dbp.PlanPrice,
		PlanHighlights:  dbp.PlanHighlights,
		PlanAdvantages:  dbp.PlanAdvantages,
		PlanLimitMbytes: dbp.PlanLimitMbytes,
		PlanLimitDays:   dbp.PlanLimitDays,
	})
}
