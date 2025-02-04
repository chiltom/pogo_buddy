package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var sessionStore *sessions.CookieStore

func init() {
	env := "development"
	if err := godotenv.Load(".env." + env); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SESSION_SECRET")
	if secretKey == "" {
		log.Fatal("SESSION_SECRET is not set in environment ")
	}

	sessionStore = sessions.NewCookieStore([]byte(secretKey))

	sessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 1 week
		HttpOnly: true,
		Secure:   false, // Set to true in production
		SameSite: http.SameSiteStrictMode,
	}
}

// SetSession creates a session for a logged-in user
func SetSession(w http.ResponseWriter, r *http.Request, userID int) error {
	session, _ := sessionStore.Get(r, "session-name")
	session.Values["user_id"] = userID
	return session.Save(r, w)
}

// GetSession retrieves the logged-in user's ID from session
func GetSession(r *http.Request) (int, bool) {
	session, _ := sessionStore.Get(r, "session-name")
	userID, ok := session.Values["user_id"].(int)
	return userID, ok
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "session-name")
	delete(session.Values, "user_id")
	session.Save(r, w)
}
