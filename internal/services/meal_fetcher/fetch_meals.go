package meal_fetcher_service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gomeal/logger/pkg/logger"
	business "github.com/gomeal/meal-api/internal/services/model"
)

func (s *serviceImpl) FetchMeals(ctx context.Context, batchSize int64) error {
	meals := make([]business.Meal, 0, batchSize)
	for range batchSize {
		meal, err := s.mealClient.FetchRandomMeal(ctx)
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("[%T.FetchMeals] failed to mealsClient.FetchRandomMeal", s), slog.Any("error", err))
			return err
		}

		meals = append(meals, meal)
	}

	createdMeals, err := s.mealRepository.CreateMeals(ctx, meals)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T.FetchMeals] failed to mealRepository.CreateMeals", s), slog.Any("error", err))
		return err
	}

	logger.Info(ctx, "successfully saved meals into a database", slog.Int("created_meals_count", len(createdMeals)))
	return nil
}
