package user

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(name, email string) (*User, error) {
	user, err := NewUser(name, email)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetAllUsers() ([]*User, error) {
	return s.repo.GetAll()
}
