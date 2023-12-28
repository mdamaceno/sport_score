package helpers

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

type Error struct {
	OriginalError error       `json:"-"`
	Data          interface{} `json:"data"`
	Name          string      `json:"name"`
	Message       string      `json:"message"`
}

func PGConflictError(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return err
	}

	return nil
}
