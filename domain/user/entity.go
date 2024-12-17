package user

import (
	"errors"
)

type User struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"test@email.com"`
}

func NewUser(name, email string) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email cannot be empty")
	}
	return &User{Name: name, Email: email}, nil
}
