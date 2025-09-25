package clients

import (
	"context"
	"net/http"

	business "github.com/gomeal/meal-api/internal/services/model"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type TheMealsDbClient interface {
	FetchRandomMeal(ctx context.Context) (business.Meal, error)
}
