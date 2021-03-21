package user

import (
	"StoreManager-DDD/entity"
	"fmt"
)

type inmem struct {
	m map[entity.ID]*entity.User
}

func newInMem() *inmem {
	var m = map[entity.ID]*entity.User{}
	return &inmem{
		m: m,
	}
}

//Create a user
func (r *inmem) Create(e *entity.User) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Update handles updating a user
func (r *inmem) Update(e *entity.User) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Get a user by id
func (r *inmem) Get(id entity.ID) (*entity.User, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//GetByEmail Get a user by email
func (r *inmem) GetByEmail(email string) (*entity.User, error) {
	for _, value := range r.m {
		if value.Email == email {
			return value, nil
		}
	}
	return nil, entity.ErrNotFound
}

//List handles listing of users
func (r *inmem) List() ([]*entity.User, error) {
	var userList []*entity.User
	for _, value := range r.m {
		userList = append(userList, value)
	}
	return userList, nil
}

// Delete handles deletion of a user
func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return fmt.Errorf("not found")
	}
	r.m[id] = nil
	return nil
}

