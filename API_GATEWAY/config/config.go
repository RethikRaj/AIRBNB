package config

import config "github.com/RethikRaj/AIRBNB/API_GATEWAY/config/env"

type Config struct {
	HTTP HTTPServerConfig
	DB   DBConfig
}

type HTTPServerConfig struct {
	Addr         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

func NewConfig() *Config {
	httpConfig := HTTPServerConfig{
		Addr:         config.GetStringValue("PORT", ":8081"),
		ReadTimeout:  config.GetIntValue("READ_TIMEOUT", 30),
		WriteTimeout: config.GetIntValue("WRITE_TIMEOUT", 30),
		IdleTimeout:  config.GetIntValue("IDLE_TIMEOUT", 60),
	}

	dbConfig := DBConfig{
		Host:     config.GetStringValue("DB_HOST", "localhost"),
		Port:     config.GetStringValue("DB_PORT", "5432"),
		User:     config.GetStringValue("DB_USER", "postgres"),
		Name:     config.GetStringValue("DB_NAME", ""),
		Password: config.GetStringValue("DB_PASSWORD", ""),
	}

	cfg := &Config{
		HTTP: httpConfig,
		DB:   dbConfig,
	}

	return cfg
}
