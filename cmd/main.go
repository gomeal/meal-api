package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/gomeal/config/pkg/config"
	"github.com/gomeal/logger/pkg/logger"
	"github.com/gomeal/meal-api/cmd/app"
	app_config "github.com/gomeal/meal-api/internal/config"
)

func main() {
	var (
		env      = logger.EnvTypeLocal
		logLevel = slog.LevelDebug
	)

	ctx := context.Background()
	logger.InitLogger(logger.EnvType(env), logLevel)
	logger.Info(ctx, "logger initialized successfully", slog.Any("env", env), slog.Any("level", logLevel.String()))

	provider := config.NewProvider(".cfg/values.yaml")
	logger.Info(ctx, "config provider created successfully", slog.String("application-name", provider.GetConfigClient().GetValue(app_config.ApplicationName).String()))

	var (
		services   = app.InitServices(ctx)
		schedulers = app.InitSchedulers(ctx, provider, services)
	)

	_ = schedulers
	time.Sleep(1 * time.Minute)
}
