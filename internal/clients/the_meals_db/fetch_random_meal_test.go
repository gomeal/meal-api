package themealsdb_client_test

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	"github.com/gomeal/logger/pkg/logger"
	"github.com/gomeal/meal-api/internal/clients"
	"github.com/gomeal/meal-api/internal/clients/mocks"
	themealsdb_client "github.com/gomeal/meal-api/internal/clients/the_meals_db"
	business "github.com/gomeal/meal-api/internal/services/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	_ "embed"
)

//go:embed testdata/mock_random_meal_response.json
var mockRandomMealResponse string

func TestFetchRandomMeal(t *testing.T) {
	type testCase struct {
		name           string
		configMock     func() clients.TheMealsDbConfig
		httpClientMock func() clients.HTTPClient
		expectedResp   business.Meal
		expectedErr    error
	}

	tests := []testCase{
		{
			name: "valid response",
			configMock: func() clients.TheMealsDbConfig {
				mock := mocks.NewTheMealsDbConfig(t)
				mock.EXPECT().Url().
					Return("http://some.url")

				return mock
			},
			httpClientMock: func() clients.HTTPClient {
				clientMock := mocks.NewHTTPClient(t)
				clientMock.EXPECT().Do(mock.Anything).
					Return(&http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(mockRandomMealResponse)),
					}, nil)

				return clientMock
			},
			expectedResp: business.Meal{
				ExternalID: 52821,
				Name:       "Laksa King Prawn Noodles",
				Category: business.MealCategory{
					Name: "Seafood",
				},
				Cuisine: business.MealCuisine{
					Name: "Malaysian",
				},
				Instructions: "Heat the oil in a medium saucepan and ...",
				ImageURL:     "https://www.themealdb.com/images/media/meals/rvypwy1503069308.jpg",
				Tags:         []string{"Shellfish", "Seafood"},
				YouTubeURL:   "https://www.youtube.com/watch?v=OcarztU8cYo",
				Ingridients: []business.MealIngridient{
					{
						Name:    "Olive Oil",
						Measure: "1 tbsp",
					},
					{
						Name:    "Red Chilli",
						Measure: "1 finely sliced",
					},
					{
						Name:    "Thai red curry paste",
						Measure: "2 tbsp",
					},
				},
				RecipeURL: "https://www.bbcgoodfood.com/recipes/prawn-laksa-curry-bowl",
			},
		},
		{
			name: "error response",
			configMock: func() clients.TheMealsDbConfig {
				mock := mocks.NewTheMealsDbConfig(t)
				mock.EXPECT().Url().
					Return("http://some.url")

				return mock
			},
			httpClientMock: func() clients.HTTPClient {
				clientMock := mocks.NewHTTPClient(t)
				clientMock.EXPECT().Do(mock.Anything).
					Return(&http.Response{
						StatusCode: http.StatusInternalServerError,
					}, errors.New("some error"))

				return clientMock
			},
			expectedResp: business.Meal{},
			expectedErr:  errors.New("some error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logger.InitLogger(logger.EnvTypeTesting, slog.LevelDebug)
			ctx := context.Background()

			theMealsDbClient := themealsdb_client.New(tc.configMock(), tc.httpClientMock())
			meal, err := theMealsDbClient.FetchRandomMeal(ctx)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedResp, meal)
		})
	}
}
