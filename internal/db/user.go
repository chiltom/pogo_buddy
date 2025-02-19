package db

import (
	"database/sql"
	"log"

	"github.com/chiltom/pogo_buddy/internal/models"
)

// CreateUser creates a new user in the database
func CreateUser(db *sql.DB, user models.User) (int, error) {
	query := `
		INSERT INTO users (
			email, 
			password, 
			first_name, 
			last_name, 
			email_verified,
			verification_token,
			verification_expiry,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
		) RETURNING id
	`

	err := db.QueryRow(
		query,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.EmailVerified,
		user.VerificationToken,
		user.VerificationExpiry,
	).Scan(&user.ID)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return 0, err
	}

	return user.ID, nil
}
