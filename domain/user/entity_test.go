package user

import (
	"testing"
)

func TestUserCreation(t *testing.T) {
	// success case
	t.Run("successful creation", func(t *testing.T) {
		name := "Alice"
		email := "alice@example.com"

		user, err := NewUser(name, email)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if user.Name != name {
			t.Errorf("expected Name to be %s, got %s", name, user.Name)
		}
		if user.Email != email {
			t.Errorf("expected Email to be %s, got %s", email, user.Email)
		}
	})

	// error case
	t.Run("creation fails with empty fields", func(t *testing.T) {
		tests := []struct {
			name  string
			email string
		}{
			{"", "alice@example.com"},
			{"Alice", ""},
		}

		for _, tt := range tests {
			user, err := NewUser(tt.name, tt.email)
			if err == nil {
				t.Errorf("expected error, but got none for input: name=%s, email=%s", tt.name, tt.email)
			}
			if user != nil {
				t.Errorf("expected user to be nil, but got %+v", user)
			}
		}
	})
}
