package app

import (
	"github.com/gomeal/meal-api/internal/repositories"
	meal_repo "github.com/gomeal/meal-api/internal/repositories/meal"
)

type Repositories struct {
	Meal repositories.MealRepository
}

func InitRepositories() Repositories {
	return Repositories{
		Meal: meal_repo.New(),
	}
}
