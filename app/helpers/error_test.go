package helpers

import (
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
)

func TestErrorHelper(t *testing.T) {
	t.Run("PGConflictError", func(t *testing.T) {
		t.Run("should return an error", func(t *testing.T) {
			t.Run("when error is a PgError with code 23505", func(t *testing.T) {
				pgErr := pgconn.PgError{Code: "23505"}
				err := PGConflictError(&pgErr)

				assert.NotNil(t, err)
			})
		})

		t.Run("should not return an error", func(t *testing.T) {
			t.Run("when error is not a PgError with code 23505", func(t *testing.T) {
				anyError := errors.New("any error")
				err := PGConflictError(anyError)

				assert.Nil(t, err)
			})
		})
	})
}
