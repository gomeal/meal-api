package meal_repo

import (
	"context"

	business "github.com/gomeal/meal-api/internal/services/model"
)

func (r *repoImpl) CreateMeals(ctx context.Context, meals []business.Meal) ([]business.Meal, error) {
	return meals, nil
}
