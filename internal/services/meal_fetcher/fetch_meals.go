package meal_fetcher_service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gomeal/logger/pkg/logger"
)

func (s *serviceImpl) FetchMeals(ctx context.Context, batchSize int64) error {
	meal, err := s.mealsClient.FetchRandomMeal(ctx)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T.FetchMeals] failed to mealsClient.FetchRandomMeal", s), slog.Any("error", err))
		return err
	}

	logger.Info(ctx, "got a meal", slog.Any("meal", meal))
	return nil
}
