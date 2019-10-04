package test

import (
	"testing"
)

func TestReturningPlans(t *testing.T) {
	plansDao := FakePlansDAO{}
	ctrl := controllers.PlanController{}

	ctrl.InjectDAO(
		plansDao,
	).GetAllCarrierPlans()

	// results :=
	// if len(results) < 1 {
	// 	t.Error("Expected at least 1 intem")
	// }

}
