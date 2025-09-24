package services

import "context"

type MealFetcherService interface {
	FetchMeals(ctx context.Context, batchSize int64) error
}
