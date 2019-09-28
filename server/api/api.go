package api

import (
	"bemobi-hire/server/api/routes"
	"bemobi-hire/server/config"
	"net/http"

	"github.com/go-chi/chi"
)

// RunAPI inicializa api
func RunAPI(appConfig config.AppConfig) *chi.Mux {
	router := chi.NewRouter()

	InitMiddlewares(router, appConfig)
	// Toda modificação na api que modifique os contratos ou crie uma breaking change deverá incrementar a versão da mesma
	// router.Route("/v2", func(r chi.Router) {
	// 	r.Mount("/plans", routes.PlansV2(appConfig))
	// })
	router.Route("/v1", func(r chi.Router) { // sinto um bad smells não sei por que, depois investigar
		r.Mount("/carrier", routes.Plans(appConfig)) // rotas de planos
	})

	// checa rotas por erros
	WalkRoutes(router)
	// Inicializa servidor
	http.ListenAndServe(":"+appConfig.AppPort, router)
	return router
}
