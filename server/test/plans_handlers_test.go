package test

import (
	"bemobi-hire/server/handlers"
	"bemobi-hire/server/test/mock_http"
	"bemobi-hire/server/test/plans_test"
	"context"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestGetPlansByCarrier(t *testing.T) {
	t.Logf("Started TestGetPlansByCarrier")
	var plansDao = plans_test.FakePlansDAO{}
	var ctrl = handlers.PlanController{}
	ctrl.InjectDAO(
		&plansDao,
	)
	w := mock_http.ResponseWriter{}
	r := httptest.NewRequest("GET", "/v1/carrier/{carrier}/plans", nil)

	// FIRST TEST WITH CLARO CARRIER
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("carrier", "claro")

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	ctrl.GetAllCarrierPlans(&w, r)
	finalCount := len(w.GetBodyJSONArray())
	if finalCount != 2 {
		t.Error("Wont returned plans for requested carrier: claro ")
	}
	t.Logf("Finished CLARO CARRIER (with [%v] count)", finalCount)

	// SECOND TEST WITH VIVO CARRIER
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("carrier", "vivo")

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	ctrl.GetAllCarrierPlans(&w, r)
	finalCount = len(w.GetBodyJSONArray())

	if finalCount != 1 {
		t.Error("Wont returned plans for requested carrier: vivo ")
	}
	t.Logf("Finished VIVO CARRIER (with [%v] count)", finalCount)

	// THIRD TEST WITH ANOYING CARRIER
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("carrier", "anoying_carrier")

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	ctrl.GetAllCarrierPlans(&w, r)
	finalCount = len(w.GetBodyJSONArray())

	if finalCount != 0 {
		t.Error("Wont returned plans for requested carrier: vivo ")
	}
	t.Logf("Finished ANOYING CARRIER (with %v count)", finalCount)
}

func TestGettingPlanDetails(t *testing.T) {
	t.Logf("Started TestGettingPlanDetails")
	var plansDao = plans_test.FakePlansDAO{}
	var ctrl = handlers.PlanController{}
	ctrl.InjectDAO(
		&plansDao,
	)
	w := mock_http.ResponseWriter{}
	r := httptest.NewRequest("GET", "/v1/carrier/{carrier}/plans/{sku}/details", nil)

	// TEST WITH CLARO
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("carrier", "claro")
	rctx.URLParams.Add("sku", "WEB_CLARO1GB")

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	ctrl.GetPlanDetails(&w, r)
	jsonBody := w.GetBodyJSON()

	if jsonBody["franchise"] != "Franquia Loka" {
		t.Error("Wont returned [ franchise : Franquia Loka ] in plan details for sku WEB_CLARO1GB")
	}
	t.Logf("Finished TestGettingPlanDetails (with %v == franchise)", jsonBody["franchise"])

	// TEST WITH VIVO
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("carrier", "vivo")
	rctx.URLParams.Add("sku", "WEB_VIVO100MB")

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	ctrl.GetPlanDetails(&w, r)
	jsonBody = w.GetBodyJSON()

	if jsonBody["franchise"] != "Franquia vivo" {
		t.Errorf("Wont returned [ franchise : %v ] in plan details for sku WEB_CLARO1GB", jsonBody["franchise"])
	}
	t.Logf("Finished TestGettingPlanDetails getting plan details correctly")

}
