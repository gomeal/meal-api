package transactor

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gomeal/logger/pkg/logger"
	"github.com/jackc/pgx/v5"
)

func (t *transactorImpl) WithinTranasction(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, pgx.TxOptions{
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].WithinTranasction failed to db.BeginTx", t), slog.Any("error", err))
		return err
	}

	if err := f(ctx); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			logger.Error(ctx, fmt.Sprintf("[%T].WithinTranasction failed to tx.Rollback", t), slog.Any("error", err))
			return rbErr
		}

		logger.Error(ctx, fmt.Sprintf("[%T].WithinTranasction failed to f(), transaction rollbacked", t), slog.Any("error", err))
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error(ctx, fmt.Sprintf("[%T].WithinTranasction failed to tx.Commit", t), slog.Any("error", err))
		return err
	}

	return nil
}
