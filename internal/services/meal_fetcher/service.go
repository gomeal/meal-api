package meal_fetcher_service

import (
	"github.com/gomeal/meal-api/internal/clients"
	"github.com/gomeal/meal-api/internal/repositories"
)

type serviceImpl struct {
	mealClient     clients.TheMealsDbClient
	mealRepository repositories.MealRepository
}

func New(mealClient clients.TheMealsDbClient, mealRepository repositories.MealRepository) *serviceImpl {
	return &serviceImpl{
		mealClient:     mealClient,
		mealRepository: mealRepository,
	}
}
