package meal_fetcher_service

import "github.com/gomeal/meal-api/internal/clients"

type serviceImpl struct {
	mealsClient clients.TheMealsDbClient
}

func New(mealsClient clients.TheMealsDbClient) *serviceImpl {
	return &serviceImpl{
		mealsClient: mealsClient,
	}
}
