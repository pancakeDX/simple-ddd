package user

import (
	"errors"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func NewUser(name, email string) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email cannot be empty")
	}
	return &User{Name: name, Email: email}, nil
}
