package main

import (
	"click-counter/config"
	"click-counter/internal/app"
	"click-counter/pkg/logger"
	"fmt"
)

func main() {
	log := logger.New()
	cfg, err := config.New()
	if err != nil {
		log.Error(fmt.Sprintf("config error: %s", err))
	}

	a := app.New(cfg, *log)
	if err = a.Run(); err != nil {
		log.Error(fmt.Sprintf("app run: %s", err))
	}
}
