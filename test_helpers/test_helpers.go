package test_helpers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/voodoostack/fitstackapi/postgres"
)

func TeardownDB(ctx context.Context, t *testing.T, db *postgres.DB) {
	t.Helper()

	err := db.Truncate(ctx)
	require.NoError(t, err)
}
