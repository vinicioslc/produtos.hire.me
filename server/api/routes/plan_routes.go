package routes

import (
	"bemobi-hire/server/config"
	"bemobi-hire/server/dao"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var pdao = dao.PlansDAO{}

// Plans Rotas para api de planos
func Plans(config config.AppConfig) *chi.Mux {
	pdao.Server = config.DBHost
	pdao.Database = config.DBDatabase
	pdao.Connect()
	router := chi.NewRouter()
	router.Get("/{carrier}/plans", GetAllCarrierPlans)
	router.Get("/{carrier}/plans/{sku}/details", GetPlanDetails)
	return router
}

// GetAllCarrierPlans Return all plans from given carrier
func GetAllCarrierPlans(w http.ResponseWriter, r *http.Request) {
	carrier := chi.URLParam(r, "carrier")

	plans, err := pdao.ListPlans(carrier)
	if err != nil {
		log.Fatal(err.Error())
	}
	render.JSON(w, r, plans)
}

// GetPlanDetails Return Plan Details from given carrier and sku
func GetPlanDetails(w http.ResponseWriter, r *http.Request) {
	carrier := chi.URLParam(r, "carrier")

	plans, err := pdao.ListPlans(carrier)
	if err != nil {
		log.Fatal(err.Error())
	}
	render.JSON(w, r, plans)
}
