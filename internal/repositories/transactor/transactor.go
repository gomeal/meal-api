package transactor

import "github.com/jackc/pgx/v5/pgxpool"

type transactorImpl struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *transactorImpl {
	return &transactorImpl{
		db: db,
	}
}
