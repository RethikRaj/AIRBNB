package app

import (
	"fmt"
	"net/http"

	config "github.com/RethikRaj/AIRBNB/API_GATEWAY/config/env"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/router"
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
	// Setup Router
	router := router.SetupRouter()

	// Setup server
	server := http.Server{
		Addr:    app.Config.Addr,
		Handler: router,
	}

	fmt.Println("Starting Server on", app.Config.Addr)

	return server.ListenAndServe()
}
