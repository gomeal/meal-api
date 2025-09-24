package meal_fetcher_cron

import (
	"context"
	"log/slog"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/gomeal/logger/pkg/logger"
)

type Config interface {
	Duration() time.Duration
	BatchSize() int64
}

type Service interface {
	FetchMeals(ctx context.Context, batchSize int64) error
}

type Cron struct {
	config  Config
	service Service
	cron    *cron.Cron
}

func New(config Config, service Service, cron *cron.Cron) *Cron {
	return &Cron{
		config:  config,
		service: service,
		cron:    cron,
	}
}

func (c *Cron) Start(ctx context.Context) {
	c.cron.Schedule(cron.Every(c.config.Duration()), cron.FuncJob(func() {
		c.Do(ctx)
	}))

	c.cron.Start()
	logger.Info(ctx, "meal fetcher cron successfully started", slog.String("duration", c.config.Duration().String()))
}

func (c *Cron) Do(ctx context.Context) {
	start := time.Now()
	defer func() {
		logger.Info(ctx, "successfully done meal fetching job", slog.String("timeEstimated", time.Since(start).String()))
	}()

	if err := c.service.FetchMeals(ctx, c.config.BatchSize()); err != nil {
		logger.Error(ctx, "error in meal fetcher cron tick", slog.Any("error", err))
		return
	}
}

func (c *Cron) Stop(ctx context.Context) {
	stopCtx := c.cron.Stop()

	select {
	case <-stopCtx.Done():
		logger.Info(ctx, "meal fetcher cron successfully stopped")
	case <-ctx.Done():
		logger.Warn(ctx, "meal fetcher cron stop interrupted by context cancellation")
	}
}
