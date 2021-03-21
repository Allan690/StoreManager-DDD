package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"
)

type Sale struct {
	Product ID
	User ID
	Total int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewSale create a new sale
func NewSale(productId, userId ID, total int) (*Sale, error) {
	sale := &Sale{
		Product:   productId,
		User:      userId,
		Total:     total,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := sale.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return sale, err
}

func (sale Sale) Validate() error  {
	return validation.ValidateStruct(&sale,
		validation.Field(&sale.User, validation.Required, is.UUID),
		validation.Field(&sale.Product, validation.Required, is.UUID),
		validation.Field(&sale.Total, validation.Required, validation.Min(1)),
		)
}
