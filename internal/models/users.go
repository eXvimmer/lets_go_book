package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// UserModel type wraps a database connection pool.
type UserModel struct {
	DB *sql.DB
}

// TODO: implement these methods

// We'll use the Insert method to add a new record to the "users" table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// The Authenticate method verifies whether a user exists with the provided
// email address and password or not. This will return the relevant user ID if
// they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// check if a user exists with a specific ID.
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
