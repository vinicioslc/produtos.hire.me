package api

import (
	"bemobi-hire/server/config"

	"github.com/go-chi/cors"
)

// GenerateCors inicializa o middlware cors
func GenerateCors(appConfig config.AppConfig) *cors.Cors {

	// Cors básico padrão liberado
	return cors.New(cors.Options{
		AllowedOrigins:   []string{appConfig.AllowedOrigins},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

}
