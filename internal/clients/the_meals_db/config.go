package themealsdb_client

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/gomeal/config/pkg/config"
	"github.com/gomeal/logger/pkg/logger"
	"github.com/gomeal/meal-api/internal/clients"
	app_config "github.com/gomeal/meal-api/internal/config"
)

var _ clients.TheMealsDbConfig = &configImpl{}

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
	url     string
	timeout time.Duration

	mu sync.RWMutex
}

func NewConfig(ctx context.Context, provider Provider) (*configImpl, error) {
	c := &configImpl{
		mu: sync.RWMutex{},
	}

	if err := c.updateUrl(ctx, provider.GetConfigClient().GetValue(app_config.TheMealsDbUrl).String()); err != nil {
		logger.Error(ctx, "unable to update url value", slog.Any("error", err))
		return nil, err
	}

	return c, nil
}

func (c *configImpl) updateUrl(ctx context.Context, url string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.url = url
	logger.Info(ctx, "updated url value", slog.String(string(app_config.TheMealsDbUrl), url))
	return nil
}

func (c *configImpl) Url() string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.url
}

func (c *configImpl) updateTimeout(ctx context.Context, timeout time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.timeout = timeout
	logger.Info(ctx, "updated timeout value", slog.String(string(app_config.TheMealsDbTimeout), timeout.String()))
	return nil
}

func (c *configImpl) Timeout() time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.timeout
}
