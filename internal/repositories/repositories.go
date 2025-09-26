package repositories

import (
	"context"

	business "github.com/gomeal/meal-api/internal/services/model"
)

type MealRepository interface {
	CreateMeals(ctx context.Context, meals []business.Meal) ([]business.Meal, error)
}
