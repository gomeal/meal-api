package app_config

import (
	"fmt"

	"github.com/gomeal/config/pkg/config"
)

const (
	ApplicationName = config.Key("application_name")
	Env             = config.Key("env")

	MealFetcherCronDuration  = config.Key("meal_fetcher_cron_duration")
	MealFetcherCronBatchSize = config.Key("meal_fetcher_cron_batch_size")

	TheMealsDbUrl     = config.Key("the_meals_db_http_url")
	TheMealsDbTimeout = config.Key("the_meals_db_http_timeout")
)

const (
	PostgresUser         = config.Secret("POSTGRES_USER")
	PostgresPassword     = config.Secret("POSTGRES_PASSWORD")
	PostgresHost         = config.Secret("POSTGRES_HOST")
	PostgresPort         = config.Secret("POSTGRES_PORT")
	PostgresDatabaseName = config.Secret("POSTGRES_DB")
)

func PostgresURI(provider config.Provider) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		provider.GetSecretClient().GetSecret(PostgresUser).String(),
		provider.GetSecretClient().GetSecret(PostgresPassword).String(),
		provider.GetSecretClient().GetSecret(PostgresDatabaseName).String(),
		provider.GetSecretClient().GetSecret(PostgresHost).String(),
		provider.GetSecretClient().GetSecret(PostgresPort).String(),
	)
}
