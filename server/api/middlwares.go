package api

import (
	"bemobi-hire/server/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// InitMiddlewares Inicializa middlewares
func InitMiddlewares(router *chi.Mux, appConfig config.AppConfig) {

	// Utiliza middlewares para tratamentos comuns como logger Recuperar de panics sem reiniciar a aplicação e requestId para parsear ids no url
	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		GenerateCors(appConfig).Handler,
	)
}
