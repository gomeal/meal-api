package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/gomeal/config/pkg/config"
	"github.com/gomeal/logger/pkg/logger"
	"github.com/gomeal/meal-api/internal/closer"
	app_config "github.com/gomeal/meal-api/internal/config"
	"github.com/gomeal/meal-api/internal/repositories"
	meal_repo "github.com/gomeal/meal-api/internal/repositories/meal"
	"github.com/gomeal/meal-api/internal/repositories/now_timer"
	"github.com/gomeal/meal-api/internal/repositories/transactor"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	Meal repositories.MealRepository
}

func InitRepositories(ctx context.Context, provider config.Provider) Repositories {
	pool, err := pgxpool.New(ctx, app_config.PostgresURI(provider))
	if err != nil {
		logger.Error(ctx, "can not pgxpool.New", slog.Any("error", err))
		panic(err)
	}
	closer.Add(func(ctx context.Context) error {
		pool.Close()
		return nil
	})

	if err := pool.Ping(ctx); err != nil {
		logger.Error(ctx, "failed to ping pgxpool", slog.Any("error", err))
		panic(err)
	}

	return Repositories{
		Meal: meal_repo.New(pool, transactor.New(pool), now_timer.New(func() time.Time {
			return time.Now().UTC()
		})),
	}
}
