package sale

import (
	"StoreManager-DDD/entity"
	"StoreManager-DDD/usecase/user"
	"errors"
	"github.com/gofrs/uuid"
)

type Service struct {
	salesRepo Repository
	userRepo user.Repository
}

func NewService(sr Repository, ur user.Repository ) *Service {
	return &Service{
		salesRepo: sr,
		userRepo: ur,
	}
}


// CreateSale creates a sale
func (s *Service) CreateSale(sale *entity.Sale) (*entity.Sale, error) {
	err := sale.Validate()
	if err != nil {
		return nil, err
	}
	return s.salesRepo.Create(sale)
}


// GetSale fetches a sale by the sale id
func (s *Service) GetSale(saleId uuid.UUID) (*entity.Sale, error) {
	return s.salesRepo.GetSale(saleId)
}


// DeleteSale deletes a sale by id
func(s *Service) DeleteSale(saleId uuid.UUID) error {
	sale, err := s.GetSale(saleId)
	if err != nil {
		return err
	}
	return s.salesRepo.Delete(sale.ID.UUID)
}

// GetAllSales fetches all available sales
func (s *Service) GetAllSales() ([]*entity.Sale, error) {
	return s.salesRepo.List()
}

// GetSalesByUserId fetches sales of a user by id
func (s *Service) GetSalesByUserId(userId uuid.UUID)  ([]*entity.Sale, error) {
	user_, err := s.userRepo.Get(userId)
	if err != nil {
		return nil, errors.New("user was not found")
	}
	return s.salesRepo.GetSalesByUserId(user_.ID.UUID)
}
