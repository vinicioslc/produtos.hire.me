package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// WalkRoutes percorre todas as rotas e verifica erros ao inicializar
func WalkRoutes(router chi.Router) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // percorre e printa todas as rotas
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // Panic se  tiver um erro em uma das rotas
	}
}
