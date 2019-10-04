package routes

import (
	"bemobi-hire/server/config"
	"bemobi-hire/server/dao"
	"bemobi-hire/server/handlers"

	"github.com/go-chi/chi"
)

// planControoller instantiate plan_controller used in this package

// Plans routes that manipulate plans
func Plans(config config.AppConfig) *chi.Mux {
	planController := handlers.PlanController{}
	plansDAO := dao.PlansDAO{
		Server:   config.DBHost,
		Database: config.DBDatabase,
	}

	plansDAO.Connect()
	planController.InjectDAO(plansDAO)

	router := chi.NewRouter()
	router.Get("/{carrier}/plans", planController.GetAllCarrierPlans)
	router.Get("/{carrier}/plans/{sku}/details", planController.GetPlanDetails)
	router.Post("/{carrier}/plans", planController.InsertPlanFromCarrier)
	router.Put("/{carrier}/plans/{sku}", planController.UpdatePlanBySKU)
	return router
}
