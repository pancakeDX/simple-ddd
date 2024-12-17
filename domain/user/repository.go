package user

type Repository interface {
	Save(user *User) error
	GetByID(id int) (*User, error)
	DeleteByID(id int) error
	GetAll() ([]*User, error)
}
