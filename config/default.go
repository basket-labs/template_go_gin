package config

import "os"

type Server struct {
	Port string
}

type appConfig struct {
	Env     string
	Version string
	Server  Server
}

var AppConfig *appConfig

func LoadAppConfig() {
	if AppConfig != nil {
		return
	}
	if os.Getenv("VERSION") == "" {
		os.Setenv("VERSION", "unknown")
	}
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "default")
	}
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}
	AppConfig = &appConfig{
		Env:     os.Getenv("ENV"),
		Version: os.Getenv("VERSION"),
		Server: Server{
			Port: os.Getenv("PORT"),
		},
	}
}
