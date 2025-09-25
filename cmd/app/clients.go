package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gomeal/config/pkg/config"
	"github.com/gomeal/logger/pkg/logger"
	"github.com/gomeal/meal-api/internal/clients"
	themealsdb_client "github.com/gomeal/meal-api/internal/clients/the_meals_db"
)

type Clients struct {
	TheMealsDbClient clients.TheMealsDbClient
}

func InitClients(ctx context.Context, provider config.Provider) Clients {
	mealsClientCfg, err := themealsdb_client.NewConfig(ctx, provider)
	if err != nil {
		logger.Error(ctx, "failed to init TheMealsDB client config", slog.Any("error", err))
		panic(err)
	}

	return Clients{
		TheMealsDbClient: themealsdb_client.New(mealsClientCfg, &http.Client{
			Timeout: mealsClientCfg.Timeout(),
		}),
	}
}
