package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chiltom/pogo_buddy/db"
	"github.com/chiltom/pogo_buddy/models"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationHandlers(t *testing.T) {
	db.ConnectDB()
	defer db.CloseDB()
	dbConn := db.GetDB()

	// TODO: Test user registration
	// Test user registration handler
	// Simulate registration
	// body := []byte(`{"username": "testuser", "password": "password123"}`)
	// req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	// req.Header.Set("Content-Type", "application/json")
	// rr := httptest.NewRecorder()
	// RegisterHandler(rr, req)
	// assert.Equal(t, http.StatusOK, rr.Code)

	user, _ := models.CreateUser(dbConn, "testuser", "John", "Doe", "john@example.com", "password123") // CHANGE TO REGISTER USER FUNCTION TEST

	// Simulate login
	body := []byte(`{"username": "testuser", "password": "password123"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	LoginHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Check session
	sessionReq, _ := http.NewRequest("GET", "/session-check", nil)
	sessionReq.Header.Set("Cookie", rr.Header().Get("Set-Cookie"))
	sessionRec := httptest.NewRecorder()
	SessionCheckHandler(sessionRec, sessionReq)

	assert.Equal(t, http.StatusOK, sessionRec.Code)

	// Simulate logout
	logoutReq, _ := http.NewRequest("POST", "/logout", nil)
	logoutRec := httptest.NewRecorder()
	LogoutHandler(logoutRec, logoutReq)

	assert.Equal(t, http.StatusOK, logoutRec.Code)

	// Clean up
	models.DeleteUser(dbConn, user.ID)
}
