package user

import (
	"StoreManager-DDD/entity"
	"github.com/gofrs/uuid"
)

// Reader interface
type Reader interface {
	Get(id uuid.UUID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	List()([]*entity.User, error)
}


// Writer interface
type Writer interface {
	Create(e *entity.User) (uuid.UUID, error)
	Update(e *entity.User) error
	Delete(id uuid.UUID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	GetUser(id uuid.UUID) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(email, password, firstName, lastName string) (uuid.UUID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id uuid.UUID) error
}
