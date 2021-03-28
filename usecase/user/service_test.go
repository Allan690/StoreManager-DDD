package user

import (
	"StoreManager-DDD/entity"
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

func TestInmem_Create(t *testing.T) {
	repo := NewInMem()
	m := NewService(repo)
	u := newFixtureUser()
	_, err := m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
}

func TestInmem_Update(t *testing.T) {
	repo := NewInMem()
	m := NewService(repo)
	u := newFixtureUser()
	id, err := m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	assert.Nil(t, err)
	saved, _ := m.GetUser(id)
	saved.FirstName = "Nyabayo"
	assert.Nil(t, m.UpdateUser(saved))
	updated, err := m.GetUser(id)
	assert.Nil(t, err)
	assert.Equal(t, "Nyabayo", updated.FirstName)
	assert.False(t, updated.UpdatedAt.IsZero())
}


func TestInmem_Delete(t *testing.T) {
	repo := NewInMem()
	m := NewService(repo)
	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2ID, _ := m.CreateUser(u2.Email, u2.Password, u2.FirstName, u2.LastName)
	err := m.DeleteUser(u1.ID.UUID)
	assert.Equal(t, entity.ErrNotFound, err)
	err = m.DeleteUser(u2ID)
	assert.Nil(t, err)
	_, err = m.GetUser(u2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}

func TestFindByEmail(t *testing.T) {
	repo := NewInMem()
	m := NewService(repo)
	u := newFixtureUser()
	_, _ = m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	user, _ := m.GetUserByEmail(u.Email)
	assert.Equal(t, user.Email, u.Email)
	assert.Equal(t, user.FirstName, u.FirstName)
}

func TestFindByEmail_NotFound(t *testing.T) {
	repo := NewInMem()
	m := NewService(repo)
	u := newFixtureUser()
	_, _ = m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	user, err := m.GetUserByEmail("tes@test.com")
	assert.Equal(t, entity.ErrNotFound, err)
	assert.Nil(t, user)
}

func TestInmem_List(t *testing.T) {
	repo := NewInMem()
	m := NewService(repo)
	u := newFixtureUser()
	u2 := newFixtureUser()
	_, _ = m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
	_, _ = m.CreateUser(u2.Email, u2.Password, u2.FirstName, u2.LastName)
	users, err := m.ListUsers()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}
