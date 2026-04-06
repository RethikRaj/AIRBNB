package router

import (
	"net/http"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	RegisterRoutes(r chi.Router)
}

func SetupRouter(userRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter()

	chiRouter.Use(middlewares.RateLimiter)
	chiRouter.Use(middleware.Logger) // Chi's built in logger middleware

	// Register Routes
	chiRouter.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	userRouter.RegisterRoutes(chiRouter)

	return chiRouter
}
