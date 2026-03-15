package app

import (
	"fmt"
	"net/http"

	config "github.com/RethikRaj/AIRBNB/API_GATEWAY/config/env"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/handlers"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/repositories"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/router"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/services"
)

type Config struct {
	Addr string
}

type Application struct {
	Config *Config
}

func NewConfig() *Config {
	port := config.GetStringValue("PORT", ":8081")

	return &Config{
		Addr: port,
	}
}

func NewApplication(config *Config) *Application {
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
		Addr:    app.Config.Addr,
		Handler: router,
	}

	fmt.Println("Starting Server on", app.Config.Addr)

	return server.ListenAndServe()
}
