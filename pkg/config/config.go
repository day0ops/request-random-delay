package config

import (
	"os"
	"strconv"
)

var (
	BaseDelay = os.Getenv("BASE_DELAY")
	ServerId  = os.Getenv("SERVER_ID")
	LogLevel  = os.Getenv("LOG_LEVEL")
)

func GetBaseDelay() int {
	baseDelay, err := strconv.Atoi(BaseDelay)
	if err != nil {
		return 0
	}
	return baseDelay
}

func GetServerId() string {
	if ServerId == "" {
		panic("no SERVER_ID defined")
	}
	return ServerId
}
