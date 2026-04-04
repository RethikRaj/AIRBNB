package router

import (
	"net/http"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/handlers"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/middlewares"
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
	myHandler := http.HandlerFunc(ur.userHandler.CreateUser)
	chiRouter.Post("/signup", middlewares.ReadAndValidateCreateUserRequest(myHandler).(http.HandlerFunc))
	chiRouter.Get("/{userID}", ur.userHandler.GetUserByID)
	chiRouter.Post("/signin", ur.userHandler.LoginUser)
}
