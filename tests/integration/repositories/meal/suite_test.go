package meal_repo_test

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/gomeal/logger/pkg/logger"
	"github.com/gomeal/meal-api/internal/repositories"
	meal_repo "github.com/gomeal/meal-api/internal/repositories/meal"
	"github.com/gomeal/meal-api/internal/repositories/now_timer"
	"github.com/gomeal/meal-api/internal/repositories/transactor"
	db_test_tools "github.com/gomeal/meal-api/tests/integration/tools/db"
	lib_testcontainers "github.com/restinbass/platform-libs/pkg/testcontainers"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	postgresContainer *lib_testcontainers.Container
	mealRepository    repositories.MealRepository
}

func (s *Suite) SetupTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	logger.InitLogger(logger.EnvTypeLocal, slog.LevelDebug)
	env := map[string]string{
		"POSTGRES_USER":     "test_user",
		"POSTGRES_PASSWORD": "test_password",
		"POSTGRES_DB":       "test_db",
	}
	postgresContainer, err := lib_testcontainers.RunContainer(
		ctx,
		lib_testcontainers.WithName("gomeal_meal_api_test_postgres"),
		lib_testcontainers.WithImage("postgres:16.4"),
		lib_testcontainers.WithEnv(env),
		lib_testcontainers.WithExposedPort("5432"),
	)
	if err != nil {
		logger.Error(ctx, "failed to start postgres testcontainer", slog.Any("error", err))
		panic(err)
	}

	logger.Info(ctx, "container is up, running migrations...", slog.String("port", postgresContainer.Port()))
	if err := os.Setenv("PG_DSN", fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		env["POSTGRES_USER"],
		env["POSTGRES_PASSWORD"],
		env["POSTGRES_DB"],
		"localhost",
		postgresContainer.Port(),
	)); err != nil {
		logger.Error(ctx, "failed to set env", slog.Any("error", err))
	}

	if err := db_test_tools.UpMigrations(ctx); err != nil {
		logger.Error(ctx, "failed to UpMigrations()", slog.Any("error", err))
		panic(err)
	}

	logger.Info(ctx, "migrations are up")
	s.postgresContainer = postgresContainer
	s.mealRepository = meal_repo.New(db_test_tools.PgxPool, transactor.New(db_test_tools.PgxPool), now_timer.New(func() time.Time {
		return time.Date(2025, time.September, 26, 21, 01, 37, 0, time.UTC)
	}))
}

func (s *Suite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger.Info(ctx, "running DownMigrations()")
	if err := db_test_tools.DownMigrations(); err != nil {
		logger.Error(ctx, "failed to DownMigrations()", slog.Any("error", err))
	}

	if err := s.postgresContainer.Terminate(ctx); err != nil {
		logger.Error(ctx, "failed to terminate postgres testcontainer", slog.Any("error", err))
		panic(err)
	}

	logger.Info(ctx, "postgres container terminated successfully")
}

func TestMealRepository(t *testing.T) {
	suite.Run(t, new(Suite))
}
