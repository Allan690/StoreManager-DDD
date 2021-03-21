package entity_test

import (
	"StoreManager-DDD/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	prod, err := entity.NewProduct(entity.NewID(), "Mac", "Macbook PRO 2020", 10000)
	assert.Nil(t, err)
	assert.Equal(t, prod.Name, "Mac")
	assert.NotNil(t, prod.User)
	assert.NotNil(t, prod.ID)
}

func TestProduct_Validate(t *testing.T) {
	type test struct {
		user    entity.ID
		name   string
		description string
		price int
		want     error
	}
	tests := []test{
		{
			user:  entity.NewID(),
			name:  "Mac",
			description: "Macbook",
			price: 2000,
			want:      nil,
		},
		{
			user:  entity.NewID(),
			name:  "",
			description: "Macbook",
			price: 2000,
			want:      entity.ErrInvalidEntity,
		},
		{
			user:  entity.NewID(),
			name:  "Mac",
			description: "Macbook",
			price: 0,
			want:      entity.ErrInvalidEntity,
		},
		{
			user:  entity.NewID(),
			name:  "Mac",
			description: "",
			price: 0,
			want:  entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := entity.NewProduct(tc.user, tc.name, tc.description, tc.price)
		assert.Equal(t, err, tc.want)
	}
}
