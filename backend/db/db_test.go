package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	// Test database connection
	ConnectDB()
	defer CloseDB()

	db := GetDB()
	err := db.Ping(context.Background())

	assert.NoError(t, err, "Database should connect without errors")
}
