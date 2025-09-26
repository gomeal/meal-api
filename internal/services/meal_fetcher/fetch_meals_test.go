package meal_fetcher_service_test

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/gomeal/logger/pkg/logger"
	"github.com/gomeal/meal-api/internal/clients"
	cl_mocks "github.com/gomeal/meal-api/internal/clients/mocks"
	"github.com/gomeal/meal-api/internal/repositories"
	repo_mocks "github.com/gomeal/meal-api/internal/repositories/mocks"
	meal_fetcher_service "github.com/gomeal/meal-api/internal/services/meal_fetcher"
	business "github.com/gomeal/meal-api/internal/services/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchMeals(t *testing.T) {
	type testCase struct {
		name           string
		mealClientMock func() clients.TheMealsDbClient
		mealRepoMock   func() repositories.MealRepository
		batchSize      int64
		expectedErr    error
	}

	tests := []testCase{
		{
			name: "valid fetch and create",
			mealClientMock: func() clients.TheMealsDbClient {
				clMock := cl_mocks.NewTheMealsDbClient(t)
				clMock.EXPECT().FetchRandomMeal(mock.Anything).
					Return(business.Meal{
						ExternalID: 5218,
					}, nil).
					Times(1)

				return clMock
			},
			mealRepoMock: func() repositories.MealRepository {
				repoMock := repo_mocks.NewMealRepository(t)
				repoMock.EXPECT().CreateMeals(mock.Anything, []business.Meal{
					{
						ExternalID: 5218,
					},
				}).Return([]business.Meal{
					{
						ID:         1,
						ExternalID: 5218,
					},
				}, nil)

				return repoMock
			},
			batchSize:   1,
			expectedErr: nil,
		},
		{
			name: "client error",
			mealClientMock: func() clients.TheMealsDbClient {
				clMock := cl_mocks.NewTheMealsDbClient(t)
				clMock.EXPECT().FetchRandomMeal(mock.Anything).
					Return(business.Meal{}, errors.New("client error")).
					Times(1)

				return clMock
			},
			mealRepoMock: func() repositories.MealRepository {
				return repo_mocks.NewMealRepository(t)
			},
			batchSize:   1,
			expectedErr: errors.New("client error"),
		},
		{
			name: "repo error",
			mealClientMock: func() clients.TheMealsDbClient {
				clMock := cl_mocks.NewTheMealsDbClient(t)
				clMock.EXPECT().FetchRandomMeal(mock.Anything).
					Return(business.Meal{
						ExternalID: 5218,
					}, nil).
					Times(1)

				return clMock
			},
			mealRepoMock: func() repositories.MealRepository {
				repoMock := repo_mocks.NewMealRepository(t)
				repoMock.EXPECT().CreateMeals(mock.Anything, []business.Meal{
					{
						ExternalID: 5218,
					},
				}).Return(nil, errors.New("repo error"))

				return repoMock
			},
			batchSize:   1,
			expectedErr: errors.New("repo error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logger.InitLogger(logger.EnvTypeTesting, slog.LevelDebug)
			ctx := context.Background()

			mealService := meal_fetcher_service.New(tc.mealClientMock(), tc.mealRepoMock())
			err := mealService.FetchMeals(ctx, tc.batchSize)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
