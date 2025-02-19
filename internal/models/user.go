package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                 int
	Email              string
	Password           string
	FirstName          string
	LastName           string
	EmailVerified      bool
	VerificationToken  *string
	VerificationExpiry *time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// HashPassword hashes the user's password before storing it
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares a given password with the stored hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
