package sale_test

import (
	"StoreManager-DDD/entity"
	"StoreManager-DDD/usecase/sale"
	"StoreManager-DDD/usecase/user"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func newFixtureUser() *entity.User {
	return &entity.User{
		ID:        entity.ID{UUID: entity.NewID()},
		Email:     "test@test.com",
		Password:  "123434",
		FirstName: "Allan",
		LastName:  "Mogusu",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func newFixtureSale() *entity.Sale {
	user_ := newFixtureUser()
	return &entity.Sale{
		ID:        entity.ID{UUID: entity.NewID()},
		Product:   entity.ID{UUID: entity.NewID()},
		User:      user_.ID,
		Total:     100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
func TestService_CreateSale(t *testing.T) {
	repo := sale.NewInMem()
	userRepo := user.NewInMem()
	m := sale.NewService(repo, userRepo)
	sale_, err := m.CreateSale(newFixtureSale())
	assert.Nil(t, err)
	assert.NotEmpty(t, sale_.CreatedAt)
}

func TestService_CreateSale_Zero_Total(t *testing.T) {
	user_ := newFixtureUser()
	fixture := &entity.Sale{
		ID:        entity.ID{UUID: entity.NewID()},
		Product:   entity.ID{UUID: entity.NewID()},
		User:      user_.ID,
		Total:     0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo := sale.NewInMem()
	userRepo := user.NewInMem()
	m := sale.NewService(repo, userRepo)
	_, err := m.CreateSale(fixture)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Total: cannot be blank.")
}

func TestService_GetAllSales(t *testing.T) {
	repo := sale.NewInMem()
	userRepo := user.NewInMem()
	m := sale.NewService(repo, userRepo)
	_, _ = m.CreateSale(newFixtureSale())
	sales, err := m.GetAllSales()
	assert.Len(t, sales, 1)
	assert.Nil(t, err)
}

func TestService_GetSale(t *testing.T) {
	repo := sale.NewInMem()
	userRepo := user.NewInMem()
	m := sale.NewService(repo, userRepo)
	s := newFixtureSale()
	_, _ = m.CreateSale(s)
	sale_, err := m.GetSale(s.ID.UUID)
	assert.NotNil(t, sale_)
	assert.Nil(t, err)
}

func TestService_DeleteSale(t *testing.T) {
	repo := sale.NewInMem()
	userRepo := user.NewInMem()
	m := sale.NewService(repo, userRepo)
	s := newFixtureSale()
	_, _ = m.CreateSale(s)
	err := m.DeleteSale(s.ID.UUID)
	assert.Nil(t, err)
}

func TestService_GetSalesByUserId(t *testing.T) {
	repo := sale.NewInMem()
	userRepo := user.NewInMem()
	m := sale.NewService(repo, userRepo)
	user_ := newFixtureUser()
	sale_ := &entity.Sale{
		ID:        entity.ID{UUID: entity.NewID()},
		Product:   entity.ID{UUID: entity.NewID()},
		User:      user_.ID,
		Total:     100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	sale2_ := &entity.Sale{
		ID:        entity.ID{UUID: entity.NewID()},
		Product:   entity.ID{UUID: entity.NewID()},
		User:      user_.ID,
		Total:     200,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, _ = userRepo.Create(user_)
	_, _ = m.CreateSale(sale_)
	_, _ = m.CreateSale(sale2_)
	sales, err := m.GetSalesByUserId(user_.ID.UUID)
	assert.NotNil(t, sales)
	assert.Len(t, sales, 2)
	assert.Nil(t, err)
	assert.Equal(t, sales[0].User.UUID, user_.ID.UUID)
}

