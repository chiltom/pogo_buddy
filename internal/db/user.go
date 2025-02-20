package db

import (
	"database/sql"
	"errors"
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

// GetUserByEmail retrieves a user from the database by email
func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	query := `
		SELECT id,
			   email,
			   password,
			   first_name,
			   last_name,
			   email_verified,
			   verification_token,
			   verification_expiry,
			   created_at,
			   updated_at
		FROM users
		WHERE email = $1
	`
	row := db.QueryRow(query, email)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.EmailVerified,
		&user.VerificationToken,
		&user.VerificationExpiry,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf("Error retrieving user: %v", err)
		return nil, err
	}

	return &user, nil
}
