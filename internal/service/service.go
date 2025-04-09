package service

import (
	"click-counter/config"
	"click-counter/internal/model"
	"click-counter/internal/repository"
	"click-counter/pkg/logger"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Service struct {
	Ctx    context.Context
	Cfg    config.Config
	DB     repository.ICounterRepository
	Logger *logger.Logger
}

func NewCounterService(ctx context.Context, cfg config.Config, connDB *pgxpool.Pool, log logger.Logger) *Service {
	return &Service{
		Ctx:    ctx,
		Cfg:    cfg,
		DB:     repository.NewCounterRepository(log, connDB),
		Logger: &log,
	}
}

func (s *Service) IncrementClick(bannerID int) error {
	return s.DB.IncrementClick(s.Ctx, bannerID)
}

func (s *Service) GetStats(bannerID int, tsFrom, tsTo time.Time) ([]model.Stats, error) {
	return s.DB.GetStats(s.Ctx, bannerID, tsFrom, tsTo)
}
