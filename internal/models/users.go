package models

import (
	"time"
)

type User struct {
	ID                 int
	Email              string
	Password           string // Hashed password
	FirstName          string
	LastName           string
	EmailVerified      bool
	VerificationToken  *string
	VerificationExpiry *time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
