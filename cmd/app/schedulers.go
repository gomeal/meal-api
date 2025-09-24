package app

import (
	"context"
	"log/slog"

	"github.com/robfig/cron/v3"

	"github.com/gomeal/config/pkg/config"
	"github.com/gomeal/logger/pkg/logger"
	"github.com/gomeal/meal-api/internal/closer"
	meal_fetcher_cron "github.com/gomeal/meal-api/internal/schedulers/meal_fetcher"
)

type Schedulers struct {
	MealFetcherCron *meal_fetcher_cron.Cron
}

func InitSchedulers(ctx context.Context, provider config.Provider, services Services) Schedulers {
	c := cron.New()
	mealFetcherCronCfg, err := meal_fetcher_cron.NewConfig(ctx, provider)
	if err != nil {
		logger.Error(ctx, "failed to init meal fetcher cron config", slog.Any("error", err))
		panic(err)
	}

	mealFetcherCron := meal_fetcher_cron.New(mealFetcherCronCfg, services.MealsFetcherService, c)
	mealFetcherCron.Start(ctx)
	closer.Add(func(ctx context.Context) error {
		logger.Info(ctx, "stopping meal fetcher cron")
		mealFetcherCron.Stop(ctx)
		return nil
	})

	return Schedulers{
		MealFetcherCron: mealFetcherCron,
	}
}
