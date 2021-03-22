package auth

import (
	"StoreManager-DDD/entity"
	"StoreManager-DDD/usecase/user"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func newFixtureUser() *entity.User {
	return &entity.User{
		ID:        entity.NewID(),
		Email:     "test@test.com",
		Password:  "123434",
		FirstName: "Allan",
		LastName:  "Mogusu",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestInmem_Create(t *testing.T) {
	repo := user.NewInMem()
	m := user.NewService(repo)
	u := newFixtureUser()
	_, _ = m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	authUseCase := NewAuthUseCase(repo, []byte("secret"), time.Hour)
	token, err := authUseCase.SignIn(u.Email, u.Password)
	assert.Nil(t, err)
	assert.IsType(t, token, "")
}

func TestInmem_CreateUser(t *testing.T) {
	repo := user.NewInMem()
	m := user.NewService(repo)
	u := newFixtureUser()
	_, _ = m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	authUseCase := NewAuthUseCase(repo, []byte("secret"), time.Hour)
	_, err := authUseCase.SignIn(u.Email, "wrong_password")
	assert.Equal(t, err, entity.ErrInvalidUserEmailOrPassword)
}

func TestInmem_ValidateToken(t *testing.T) {
	repo := user.NewInMem()
	m := user.NewService(repo)
	u := newFixtureUser()
	_, _ = m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	authUseCase := NewAuthUseCase(repo, []byte("secret"), time.Hour)
	token, err := authUseCase.SignIn(u.Email, u.Password)
	claims, err := authUseCase.ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, claims.Email, u.Email)
}

func TestInmem_ValidateInvalidToken(t *testing.T) {
	repo := user.NewInMem()
	m := user.NewService(repo)
	u := newFixtureUser()
	_, _ = m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	authUseCase := NewAuthUseCase(repo, []byte("secret"), time.Hour)
	token, err := authUseCase.SignIn(u.Email, "")
	_, err = authUseCase.ValidateToken(token)
	assert.NotNil(t, err)
}