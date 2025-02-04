package models

import (
	"testing"

	"github.com/chiltom/pogo_buddy/db"
	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {
	db.ConnectDB()
	defer db.CloseDB()
	dbConn := db.GetDB()

	// Create a new user
	user, err := CreateUser(dbConn, "testuser", "John", "Doe", "john@example.com", "password123")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)

	// Get user by ID
	retrievedUser, err := GetUserByID(dbConn, user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, retrievedUser.Username)

	// Delete user
	err = DeleteUser(dbConn, user.ID)
	assert.NoError(t, err)

	// Verify user deletion
	deletedUser, err := GetUserByID(dbConn, user.ID)
	assert.Error(t, err, "User should not exist after deletion")
	assert.Nil(t, deletedUser)
}
