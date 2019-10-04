package models

// DatabasePlan is a type used to write into database full plan data .
type DatabasePlan struct {
	Plan
	PlanMoreDetails PlanDetails `bson:"plan_more_details" json:"plan_more_details"`
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
