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
	router.Get("/{carrier}", GetAllCarrierPlans)
	return router
}

// GetAllCarrierPlans Retorna todos os planos da rota plans
func GetAllCarrierPlans(w http.ResponseWriter, r *http.Request) {
	carrier := chi.URLParam(r, "carrier")

	plans, err := pdao.ListPlans(carrier)
	if err != nil {
		log.Fatal(err.Error())
	}
	render.JSON(w, r, plans)
}
