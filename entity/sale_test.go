package entity_test

import (
	"StoreManager-DDD/entity"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestNewSale(t *testing.T) {
	sale, err := entity.NewSale(entity.NewID(), entity.NewID(), 300)
	assert.Nil(t, err)
	assert.Equal(t, sale.Total, 300)
	assert.IsType(t, sale.Product, entity.ID{})
}

func TestSale_Validate(t *testing.T) {
	type test struct {
		product uuid.UUID
		user  uuid.UUID
		total int
		want  error
	}
	tests := []test{
		{
			user:  entity.NewID(),
			product: entity.NewID(),
			total: 300,
			want:      nil,
		},
		{
			user:  entity.NewID(),
			product: entity.NewID(),
			want:      entity.ErrInvalidEntity,
		},
		{
			user:  entity.NewID(),
			total:  0,
			product: entity.NewID(),
			want:      entity.ErrInvalidEntity,
		},
		{
			user:  entity.NewID(),
			product: entity.NewID(),
			total: -1000,
			want:  entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := entity.NewSale(tc.product, tc.user, tc.total)
		assert.Equal(t, err, tc.want)
	}
}
