package repositories

import (
	"context"
	"time"

	business "github.com/gomeal/meal-api/internal/services/model"
)

type NowTimer interface {
	Now() time.Time
}

type Transactor interface {
	WithinTranasction(ctx context.Context, f func(ctx context.Context) error) error
}

type MealRepository interface {
	CreateMeals(ctx context.Context, meals []business.Meal) ([]business.Meal, error)
}
