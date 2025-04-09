package app

import (
	"click-counter/config"
	"click-counter/internal/handler"
	"click-counter/internal/repository"
	"click-counter/internal/service"
	"click-counter/pkg/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const migrationsPath = "migrations"

type App struct {
	cfg config.Config
	log logger.Logger
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	connDB, err := repository.ConnectDB(a.cfg.GetDbConfig().GetDsn())
	if err != nil {
		a.log.Errorf("connect to db failed: %v", err)
		return err
	}

	if err = repository.MigrateUp(a.cfg.DB, migrationsPath); err != nil {
		a.log.Errorf("migrations failed: %v", err)
		return err
	}

	c := service.NewCounterService(ctx, a.cfg, connDB, a.log)
	h := handler.NewCounterHandler(c)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		<-exit
		h.Shutdown(ctx)
		cancel()
	}()

	return h.Start(fmt.Sprintf(":%s", a.cfg.App.Port))
}

func New(cfg config.Config, log logger.Logger) *App {
	return &App{
		cfg: cfg,
		log: log,
	}

}
