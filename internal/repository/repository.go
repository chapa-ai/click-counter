package repository

import (
	"click-counter/internal/model"
	"click-counter/pkg/errors"
	"click-counter/pkg/logger"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
	"time"
)

type ICounterRepository interface {
	IncrementClick(ctx context.Context, bannerID int) error
	GetStats(ctx context.Context, bannerID int, tsFrom time.Time, tsTo time.Time) ([]model.Stats, error)
}

var (
	once        sync.Once
	pool        *pgxpool.Pool
	connPoolErr error
)

type counterRepository struct {
	logger logger.Logger
	db     *pgxpool.Pool
}

func ConnectDB(dsn string) (*pgxpool.Pool, error) {
	once.Do(func() {
		cfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			connPoolErr = err
			return
		}

		pool, err = pgxpool.ConnectConfig(context.Background(), cfg)
		if err != nil {
			connPoolErr = err
			return
		}

		if err = pool.Ping(context.Background()); err != nil {
			connPoolErr = err
			pool.Close()
			return
		}
	})
	return pool, connPoolErr
}

func NewCounterRepository(log logger.Logger, connDB *pgxpool.Pool) ICounterRepository {
	return &counterRepository{
		logger: log,
		db:     connDB,
	}
}

func (r *counterRepository) IncrementClick(ctx context.Context, bannerID int) error {
	_, err := r.db.Exec(ctx, `
INSERT INTO clicks (banner_id, timestamp, count) VALUES ($1, $2, 1) ON CONFLICT (banner_id, timestamp) DO UPDATE SET count = clicks.count + 1`,
		bannerID, time.Now().Truncate(time.Minute))
	if err != nil {
		return errors.Wrap(err, "failed to insert click")
	}

	return nil
}

func (r *counterRepository) GetStats(ctx context.Context, bannerID int,
	tsFrom time.Time, tsTo time.Time) ([]model.Stats, error) {
	rows, err := r.db.Query(ctx, `SELECT banner_id, timestamp, count  FROM clicks WHERE banner_id = $1 
	AND timestamp BETWEEN $2 AND $3`, bannerID, tsFrom, tsTo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query stats")
	}
	defer rows.Close()

	var stats []model.Stats
	for rows.Next() {
		var stat model.Stats
		err = rows.Scan(&stat.BannerID, &stat.Timestamp, &stat.Count)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		stats = append(stats, stat)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error scanning rows")
	}
	return stats, nil
}
