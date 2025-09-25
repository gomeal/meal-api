package themealsdb_client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gomeal/logger/pkg/logger"
	transport_converter "github.com/gomeal/meal-api/internal/clients/converter"
	transport "github.com/gomeal/meal-api/internal/clients/model"
	business "github.com/gomeal/meal-api/internal/services/model"
)

var (
	ErrStatusCodeIsNotOk = errors.New("http response status code is not 200 OK")
)

func (c *Client) FetchRandomMeal(ctx context.Context) (business.Meal, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.config.Url(), nil)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T.FetchRandomMeal] failed to http.NewRequestWithContext", c), slog.Any("error", err))
		return business.Meal{}, err
	}

	resp, err := c.cl.Do(req)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T.FetchRandomMeal] failed to cl.Do", c), slog.Any("error", err))
		return business.Meal{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error(ctx, fmt.Sprintf("[%T.FetchRandomMeal] response status code is not 200", c), slog.Int("staus_code", resp.StatusCode))
		return business.Meal{}, ErrStatusCodeIsNotOk
	}

	var mealResponse transport.RandomMealResponse
	if err := json.NewDecoder(resp.Body).Decode(&mealResponse); err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T.FetchRandomMeal] failed to Decode", c), slog.Any("error", err))
		return business.Meal{}, err
	}

	return transport_converter.TransportMealToBusinessMeal(mealResponse.Meals[0]), nil
}
