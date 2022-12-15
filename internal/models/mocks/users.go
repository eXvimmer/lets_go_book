package mocks

import (
	"time"

	"github.com/exvimmer/lets_go/snippetbox/internal/models"
)

type UserModel struct{}

func (u *UserModel) Insert(name, email, password string) error {
	switch email {
	case "duplicated@gmail.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "mustafa@gmail.com" && password == "password" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	if id == 1 {
		u := &models.User{
			ID:      1,
			Name:    "Mustafa",
			Email:   "mustafa@gmail.com",
			Created: time.Now(),
		}
		return u, nil
	}
	return nil, models.ErrNoRecord
}
