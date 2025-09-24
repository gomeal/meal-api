package app

import (
	"context"

	"github.com/gomeal/meal-api/internal/services"
	meal_fetcher_service "github.com/gomeal/meal-api/internal/services/meal_fetcher"
)

type Services struct {
	MealsFetcherService services.MealFetcherService
}

func InitServices(ctx context.Context) Services {
	return Services{
		MealsFetcherService: meal_fetcher_service.New(),
	}
}
