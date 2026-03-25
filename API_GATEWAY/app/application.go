package app

import (
	"fmt"
	"net/http"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/config"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/handlers"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/repositories"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/router"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/services"
)

type Application struct {
	Config *config.Config
}

func NewApplication(config *config.Config) *Application {
	return &Application{
		Config: config,
	}
}

func (app *Application) Run() error {

	// Repositories
	userRepository := repositories.NewUserRepository()

	// Service
	userService := services.NewUserService(userRepository)

	// Handlers
	userHandler := handlers.NewUserHandler(userService)

	// Setup Router
	userRouter := router.NewUserRouter(userHandler)
	router := router.SetupRouter(userRouter)

	// Setup server
	server := http.Server{
		Addr:    app.Config.HTTP.Addr,
		Handler: router,
	}

	fmt.Println("Starting Server on", app.Config.HTTP.Addr)

	return server.ListenAndServe()
}
