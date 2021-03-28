package sale

import (
	"StoreManager-DDD/entity"
	"github.com/gofrs/uuid"
)

type Reader interface {
	GetSale(s uuid.UUID) (*entity.Sale, error)
	GetSalesByUserId(u uuid.UUID)  ([]*entity.Sale, error)
	List() ([]*entity.Sale, error)
}

type Writer interface {
	Create(s *entity.Sale) (*entity.Sale, error)
	Delete(s uuid.UUID) error
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	CreateSale(s *entity.Sale) (*entity.Sale, error)
	GetSale(s uuid.UUID) (*entity.Sale, error)
	DeleteSale(s uuid.UUID) (*entity.Sale, error)
	GetSalesByProductId(p uuid.UUID) ([]*entity.Sale, error)
	GetSalesByUserId(u uuid.UUID) ([]*entity.Sale, error)
	GetAllSales() ([]*entity.Sale, error)
}
