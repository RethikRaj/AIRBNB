package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from the .env file.
// It should be called at the beginning of the application before accessing any environment variables.
func LoadEnv() {
	err := godotenv.Load() // you can mention the path to your .env file if it's not in the root directory of your project.

	if err != nil {
		log.Fatal("Error loading .env file : ", err)
	}
}

func GetStringValue(key string, fallback string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}

func GetIntValue(key string, fallback int) int {
	value, ok := os.LookupEnv(key) // LookupEnv always returns a string so we need to convert it to an integer.

	if !ok {
		return fallback
	}
	intValue, err := strconv.Atoi(value)

	if err != nil {
		fmt.Printf("Error converting %s to integer: %v", key, err)
		return fallback
	}

	return intValue
}

func GetBoolValue(key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}
	boolValue, err := strconv.ParseBool(value)

	if err != nil {
		fmt.Printf("Error converting %s to boolean: %v", key, err)
		return fallback
	}

	return boolValue
}
