package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chiltom/pogo_buddy/auth"
	"github.com/chiltom/pogo_buddy/db"
	"github.com/chiltom/pogo_buddy/models"
)

// LoginHandler authenticates a user and creates a session
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := models.AuthenticateUser(db.GetDB(), credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Set session
	err = auth.SetSession(w, r, user.ID)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

// LogoutHandler clears the user's session
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.ClearSession(w, r)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}

// SessionCheckHandler verifies if a user is logged in
func SessionCheckHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetSession(r)
	if !ok {
		http.Error(w, "No active session", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User is logged in",
		"user_id": userID,
	})
}

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newUser, err := models.CreateUser(db.GetDB(), user.Username, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "User registered successfully",
		"username": newUser.Username,
		"user_id":  newUser.ID,
	})
}
