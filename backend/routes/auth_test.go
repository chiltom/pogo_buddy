package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAuthRoutes(t *testing.T) {
	router := mux.NewRouter()
	AuthRoutes(router)

	// Test if login route is registered
	req, _ := http.NewRequest("POST", "/login", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusNotFound, rr.Code, "Login route should be registered")

	// Test if logout route is registered
	req, _ = http.NewRequest("POST", "/logout", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusNotFound, rr.Code, "Logout route should be registered")

	// Test if session-check route is registered
	req, _ = http.NewRequest("GET", "/session-check", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusNotFound, rr.Code, "Session-check route should be registered")

	// Test if register route is registered
	req, _ = http.NewRequest("POST", "/register", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusNotFound, rr.Code, "Register route should be registered")
}
