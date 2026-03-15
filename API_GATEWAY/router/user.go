package router

import (
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/handlers"
	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userHandler *handlers.UserHandler
}

func NewUserRouter(_userHandler *handlers.UserHandler) Router {
	return &UserRouter{
		userHandler: _userHandler,
	}
}

func (ur *UserRouter) RegisterRoutes(chiRouter chi.Router) {
	// Register user routes here
	chiRouter.Post("/signup", ur.userHandler.CreateUser)
}
