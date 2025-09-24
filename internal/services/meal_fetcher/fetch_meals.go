package meal_fetcher_service

import (
	"context"
	"log/slog"

	"github.com/gomeal/logger/pkg/logger"
)

func (s *serviceImpl) FetchMeals(ctx context.Context, batchSize int64) error {
	logger.Info(ctx, "fetching meals from the meals db api...", slog.Int64("batchSize", batchSize))
	return nil
}
