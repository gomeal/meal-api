package meal_fetcher_cron

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/gomeal/config/pkg/config"
	"github.com/gomeal/logger/pkg/logger"
	app_config "github.com/gomeal/meal-api/internal/config"
)

var _ Config = &configImpl{}

type Provider interface {
	config.Provider
}

type ConfigClient interface {
	config.ConfigClient
}

type Value interface {
	config.Value
}

type configImpl struct {
	duration  time.Duration
	batchSize int64

	mu sync.RWMutex
}

func NewConfig(ctx context.Context, provider Provider) (*configImpl, error) {
	c := &configImpl{
		mu: sync.RWMutex{},
	}

	if err := c.updateDuration(ctx, provider.GetConfigClient().GetValue(app_config.MealFetcherCronDuration).Duration()); err != nil {
		logger.Error(ctx, "unable to update duration value", slog.Any("error", err))
		return nil, err
	}

	if err := c.updateBatchSize(ctx, int64(provider.GetConfigClient().GetValue(app_config.MealFetcherCronBatchSize).Int())); err != nil {
		logger.Error(ctx, "unable to update batchSize value", slog.Any("error", err))
		return nil, err
	}

	return c, nil
}

func (c *configImpl) updateDuration(ctx context.Context, duration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.duration = duration
	logger.Info(ctx, "updated duration value", slog.String(string(app_config.MealFetcherCronDuration), duration.String()))
	return nil
}

func (c *configImpl) Duration() time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.duration
}

func (c *configImpl) updateBatchSize(ctx context.Context, batchSize int64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.batchSize = batchSize
	logger.Info(ctx, "updated batchSize value", slog.Int64(string(app_config.MealFetcherCronBatchSize), batchSize))
	return nil
}

func (c *configImpl) BatchSize() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.batchSize
}
