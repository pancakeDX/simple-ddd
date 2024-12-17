package user

import (
	"errors"
	"testing"

	"simple-ddd/domain/user"
)

func TestMemoryRepository(t *testing.T) {
	// Arrange: create MemoryRepository
	repo := NewMemoryRepository()

	// prepare user data
	user1, err := user.NewUser("Alice", "alice@example.com")
	if err != nil {
		t.Fatalf("unexpected error when creating user1: %v", err)
	}
	user2, err := user.NewUser("Bob", "bob@example.com")
	if err != nil {
		t.Fatalf("unexpected error when creating user2: %v", err)
	}

	// Act: 儲存用戶
	err = repo.Save(user1)
	if err != nil {
		t.Fatalf("unexpected error when saving user1: %v", err)
	}
	err = repo.Save(user2)
	if err != nil {
		t.Fatalf("unexpected error when saving user2: %v", err)
	}

	// Assert: get user by id
	t.Run("find existing user", func(t *testing.T) {
		foundUser, err := repo.GetByID(1)
		if err != nil {
			t.Fatalf("unexpected error when finding user1: %v", err)
		}
		if foundUser.ID != 1 || foundUser.Name != "Alice" || foundUser.Email != "alice@example.com" {
			t.Errorf("user data does not match, got %+v", foundUser)
		}
	})

	// Assert: get error user
	t.Run("find non-existent user", func(t *testing.T) {
		_, err := repo.GetByID(3)
		if !errors.Is(err, ErrUserNotFound) {
			t.Errorf("expected error %v, but got %v", ErrUserNotFound, err)
		}
	})

	// Assert: delete user
	t.Run("delete user", func(t *testing.T) {
		err := repo.DeleteByID(2)
		if err != nil {
			t.Fatalf("unexpected error during delete: %v", err)
		}
		_, err = repo.GetByID(2)
		if !errors.Is(err, ErrUserNotFound) {
			t.Errorf("expected error %v when finding deleted user, but got %v", ErrUserNotFound, err)
		}
	})
}
