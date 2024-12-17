package user

import (
	"errors"
	"simple-ddd/domain/user"
	"sync"
)

var ErrUserNotFound = errors.New("user not found")

type MemoryRepository struct {
	data      map[int]*user.User
	mu        sync.RWMutex
	autoIDSeq int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make(map[int]*user.User),
	}
}

func (r *MemoryRepository) Save(u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.autoIDSeq++
	u.ID = r.autoIDSeq
	r.data[u.ID] = u
	return nil
}

func (r *MemoryRepository) GetByID(id int) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, exists := r.data[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return u, nil
}

func (r *MemoryRepository) DeleteByID(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.data[id]; !exists {
		return ErrUserNotFound
	}
	delete(r.data, id)
	return nil
}

func (r *MemoryRepository) GetAll() ([]*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	users := make([]*user.User, 0, len(r.data))
	for _, u := range r.data {
		users = append(users, u)
	}
	return users, nil
}
