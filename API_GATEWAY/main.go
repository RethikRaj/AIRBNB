package main

import (
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/app"
	"github.com/joho/godotenv"
)

func main() {
	cfg := app.NewConfig(":8080")

	app := app.NewApplication(cfg)

	app.Run()

	godotenv.Load()
}
