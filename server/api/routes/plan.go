package routes

import (
	"bemobi-hire/server/config"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Plans Rotas para api de planos
func Plans(config config.AppConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{carrier}", GetAllCarrierPlans)
	return router
}

// GetAllCarrierPlans Retorna todos os planos da rota plans
func GetAllCarrierPlans(w http.ResponseWriter, r *http.Request) {
	carrier := chi.URLParam(r, "carrier")

	render.JSON(w, r, carrier)
}
