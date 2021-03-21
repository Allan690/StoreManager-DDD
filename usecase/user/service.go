package user

import (
	"StoreManager-DDD/entity"
	"time"
)

// Service interface
type Service struct {
	repo Repository
}


// NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateUser handles creation of a user
func (s *Service) CreateUser(email, password, firstName, lastName string)(entity.ID, error) {
	u, err := entity.NewUser(email, password, firstName, lastName)
	if err != nil {
		return entity.NewID(), err
	}
	return s.repo.Create(u)
}

// GetUser handles getting a user by id
func (s *Service) GetUser(id entity.ID) (*entity.User, error) {
	return s.repo.Get(id)
}

// GetUserByEmail handles getting user by email
func (s *Service) GetUserByEmail(email string) (*entity.User, error) {
	return s.repo.GetByEmail(email)
}

// ListUsers handles listing users
func (s *Service) ListUsers() ([]*entity.User, error) {
	return s.repo.List()
}


// DeleteUser handles deleting a user
func (s *Service) DeleteUser(id entity.ID) error {
	user, err := s.GetUser(id)
	if user == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// UpdateUser handles updating a user
func (s *Service) UpdateUser(e *entity.User) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
