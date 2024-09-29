package config

import (
	"os"
)

var (
	BaseDelay = os.Getenv("BASE_DELAY")
	LogLevel  = os.Getenv("LOG_LEVEL")
)
