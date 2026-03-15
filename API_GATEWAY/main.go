package main

import (
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/app"
	config "github.com/RethikRaj/AIRBNB/API_GATEWAY/config/env"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	config.LoadEnv()

	cfg := app.NewConfig()

	app := app.NewApplication(cfg)

	app.Run()

	godotenv.Load()
}
