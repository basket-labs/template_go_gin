package app

import (
	"awesomeProject/config"
)

func GetVersion() string {
	return config.AppConfig.Version
}
