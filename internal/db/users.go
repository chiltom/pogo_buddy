package db

import (
	"database/sql"
	"log"

	"github.com/chiltom/pogo_buddy/internal/models"
)

type UserStore struct {
  db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
  return &UserStore{db: db}
}

func (s *UserStore) Create(user models.User) (int, error) {
	query := `
		INSERT INTO users (
			email, password, first_name, last_name, email_verified,
			verification_token, verification_expiry, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) 
    RETURNING id
	`
  var id int
  err := s.db.QueryRow(
    query,
    user.Email,
    user.Password,
    user.FirstName,
    user.LastName,
    user.EmailVerified,
    user.VerificationToken,
    user.VerificationExpiry).Scan(&id)
  if err != nil {
    log.Printf("Error inserting user: %v", err)
    return 0, err
  }
  return id, nil
}

func (s *UserStore) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, password, first_name, last_name, email_verified,
			     verification_token, verification_expiry, created_at, updated_at
		FROM users
		WHERE email = $1
	`
  var user models.User
  err := s.db.QueryRow(query, email).Scan(
    &user.ID, &user.Email, &user.FirstName, &user.LastName,
    &user.EmailVerified, &user.VerificationToken, &user.VerificationExpiry,
    &user.CreatedAt, &user.UpdatedAt)
  if err == sql.ErrNoRows{
    log.Printf("No user exists with email: %v", email)
    return nil, nil
  }
  if err != nil {
    log.Printf("Error retrieving user: %v", err)
    return nil, err
  }
  return &user, nil
}

func (s *UserStore) Update(user models.User) error {
	query := `
		UPDATE users
		SET email = $1, password = $2, first_name = $3, last_name = $4,
			  email_verified = $5, verification_token = $6, verification_expiry = $7,
			  updated_at = NOW()
		WHERE id = $8
	`
	_, err := s.db.Exec(
		query, 
    user.Email, user.Password, user.FirstName, user.LastName,
    user.EmailVerified, user.VerificationToken, user.VerificationExpiry,
		user.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

func (s *UserStore) Delete(id int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`
	_, err := s.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
