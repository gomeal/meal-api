package meal_repo

import (
	"github.com/gomeal/meal-api/internal/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repoImpl struct {
	db         *pgxpool.Pool
	transactor repositories.Transactor
	nowTimer   repositories.NowTimer
}

func New(
	db *pgxpool.Pool,
	transactor repositories.Transactor,
	nowTimer repositories.NowTimer,
) *repoImpl {
	return &repoImpl{
		db:         db,
		transactor: transactor,
		nowTimer:   nowTimer,
	}
}
