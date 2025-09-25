package clients

import (
	"context"

	business "github.com/gomeal/meal-api/internal/services/model"
)

type TheMealsDbClient interface {
	FetchRandomMeal(ctx context.Context) (business.Meal, error)
}
