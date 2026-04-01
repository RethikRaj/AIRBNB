package main

import (
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/app"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/config"
	envConfig "github.com/RethikRaj/AIRBNB/API_GATEWAY/config/env"
)

func main() {
	// Load environment variables from .env file into the process environment
	envConfig.LoadEnv()

	// Create the config
	cfg := config.NewConfig()

	// Create the application
	app := app.NewApplication(cfg)

	// Wire up dependencies and start the server
	app.Run()

}
