package router

import (
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

	// Way One of registering middlewares
	// myHandler := http.HandlerFunc(ur.userHandler.CreateUser)
	// chiRouter.Post("/signup", middlewares.ReadAndValidateCreateUserRequest(myHandler).(http.HandlerFunc))

	// Way Two
	chiRouter.With(middlewares.ReadAndValidateCreateUserRequest).Post("/signup", ur.userHandler.CreateUser)

	chiRouter.With(middlewares.ReadAndValidateSignInUserRequest).Post("/signin", ur.userHandler.LoginUser)

	// ProtectedRoutes
	chiRouter.With(middlewares.JWTAuthMiddleware).Get("/profile", ur.userHandler.GetUserByID)

}
