package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
}

// CreateUser inserts a new user into the database
func CreateUser(pool *pgxpool.Pool,
	username string,
	firstName string,
	lastName string,
	email string,
	password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	query := `INSERT INTO users (username, first_name, last_name, email, password)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, created_at, active`

	var user User
	err = pool.QueryRow(context.Background(), query, username, firstName, lastName, email, hashedPassword).Scan(&user.ID, &user.CreatedAt, &user.Active)
	if err != nil {
		return nil, err
	}

	user.Username = username
	user.Email = email
	return &user, nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(pool *pgxpool.Pool, userID int) (*User, error) {
	query := `SELECT id, username, first_name, last_name, email, created_at, active
			  FROM users
			  WHERE id = $1`

	var user User
	err := pool.QueryRow(context.Background(), query, userID).Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Active)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAllUsers retrieves all users
func GetAllUsers(pool *pgxpool.Pool) ([]User, error) {
	query := `SELECT id, username, first_name, last_name, email, created_at, active
			  FROM users`

	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Active); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// UpdateUser updates a user's information
func UpdateUser(pool *pgxpool.Pool, userID int, username string, firstName string, lastName string, email string) (*User, error) {
	query := `UPDATE users
			  SET username = $1, first_name = $2, last_name = $3, email = $4
			  WHERE id = $5
			  RETURNING created_at, active`

	var user User
	err := pool.QueryRow(context.Background(), query, username, firstName, lastName, email, userID).Scan(&user.CreatedAt, &user.Active)
	if err != nil {
		return nil, err
	}

	user.ID = userID
	user.Username = username
	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email

	return &user, nil
}

// DeleteUser removes a user
func DeleteUser(pool *pgxpool.Pool, userID int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := pool.Exec(context.Background(), query, userID)
	return err
}
