package main

import (
	"github.com/tiburciohugo/rest-api-test/config"
	"github.com/tiburciohugo/rest-api-test/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Initialize()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	router.Initialize()
}
