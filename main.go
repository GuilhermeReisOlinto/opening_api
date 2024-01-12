package main

import (
	"opening_api/config"
	"opening_api/internal/presentation/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
	}

	//Initialize router
	router.InitializeRouter()
}
