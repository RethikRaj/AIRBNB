package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Config struct {
	Addr string
}

type Application struct {
	Config *Config
}

func NewConfig(addr string) *Config {
	return &Config{
		Addr: addr,
	}
}

func NewApplication(config *Config) *Application {
	return &Application{
		Config: config,
	}
}

func (app *Application) Run() error {
	// Setup Router
	r := chi.NewRouter()

	// Setup server
	server := http.Server{
		Addr:    app.Config.Addr,
		Handler: r,
	}

	fmt.Println("Starting Server on", app.Config.Addr)

	return server.ListenAndServe()
}
