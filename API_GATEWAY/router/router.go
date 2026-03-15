package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {

	chiRouter := chi.NewRouter()

	// Register Routes
	chiRouter.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	return chiRouter
}
