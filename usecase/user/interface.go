package user

import "StoreManager-DDD/entity"

// Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	List()([]*entity.User, error)
}


// Writer interface
type Writer interface {
	Create(e *entity.User) (entity.ID, error)
	Update(e *entity.User) error
	Delete(id entity.ID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	GetUser(id entity.ID) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(email, password, firstName, lastName string) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id entity.ID) error
}
