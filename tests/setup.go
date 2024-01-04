package tests

import (
	"testing"

	"gorm.io/gorm"
)

type SetupOptions struct {
	DB *gorm.DB
}

func SetupTest(t *testing.T, opts *SetupOptions) func(t *testing.T) {
	if &opts.DB != nil {
		TruncateTables(opts.DB)
	}

	return func(t *testing.T) {
		if &opts.DB != nil {
			TruncateTables(opts.DB)
		}
	}
}
