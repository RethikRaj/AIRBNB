package router

import (
	"net/http"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(userHandler *handlers.UserHandler) *chi.Mux {

	chiRouter := chi.NewRouter()

	// Register Routes
	chiRouter.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	chiRouter.Post("/signup", userHandler.CreateUser)

	return chiRouter
}
