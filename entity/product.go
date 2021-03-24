package entity

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofrs/uuid"
	"time"
)

// product data
type Product struct {
	ID ID
	User ID
	Name string
	Description string
	Price int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProduct(userId uuid.UUID, name, description string, price int) (*Product, error) {
	product := &Product{
		ID:          ID{NewID()},
		User:        ID{userId},
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return product, nil
}

// Validate validates product struct
func (p Product) Validate() error  {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&p.User, validation.Required, is.UUID),
		validation.Field(&p.Description, validation.Required, validation.Length(3, 1000)),
		validation.Field(&p.Price, validation.Required, validation.Min(1)),
		)
}
