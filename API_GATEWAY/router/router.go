package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	RegisterRoutes(r chi.Router)
}

func SetupRouter(userRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter()

	// Register Routes
	chiRouter.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	userRouter.RegisterRoutes(chiRouter)

	return chiRouter
}
