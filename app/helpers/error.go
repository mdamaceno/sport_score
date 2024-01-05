package helpers

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func PGConflictError(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return err
	}

	return nil
}
