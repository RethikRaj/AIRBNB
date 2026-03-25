package main

import (
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/app"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/config"
	envConfig "github.com/RethikRaj/AIRBNB/API_GATEWAY/config/env"
)

func main() {
	// load environment variables
	envConfig.LoadEnv()

	cfg := config.NewConfig()

	app := app.NewApplication(cfg)

	app.Run()

}
