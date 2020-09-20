package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDatabaste(t *testing.T) {
	var testPath = "test.db"
	db, err := New(testPath)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close())
		os.RemoveAll(testPath)
	})
}
