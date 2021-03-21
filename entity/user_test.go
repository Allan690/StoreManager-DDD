package entity_test

import (
	"StoreManager-DDD/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	u, err := entity.NewUser("sjobs@apple.com", "new_password", "Steve", "Jobs")
	assert.Nil(t, err)
	assert.Equal(t, u.FirstName, "Steve")
	assert.NotNil(t, u.ID)
	assert.NotEqual(t, u.Password, "new_password")
}

func TestValidatePassword(t *testing.T) {
	u, _ := entity.NewUser("sjobs@apple.com", "new_password123", "Steve", "Jobs")
	err := u.ValidatePassword("new_password123")
	assert.Nil(t, err)
	err = u.ValidatePassword("wrong_password123")
	assert.NotNil(t, err)
}


