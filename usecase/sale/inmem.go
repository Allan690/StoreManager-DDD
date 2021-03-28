package sale

import (
	"StoreManager-DDD/entity"
	"fmt"
	"github.com/gofrs/uuid"
)

type inmem struct {
	m map[entity.ID]*entity.Sale
	u map[entity.ID]*entity.User
}


func NewInMem() *inmem{
	var m = map[entity.ID]*entity.Sale{}
	var u = map[entity.ID]*entity.User{}
	return &inmem{
		m: m,
		u: u,
	}
}

func (r *inmem) Create(e *entity.Sale) (*entity.Sale, error) {
	r.m[e.ID] = e
	return e, nil
}

func (r *inmem) Delete(id uuid.UUID) error {
	if r.m[entity.ID{UUID: id}] == nil {
		return fmt.Errorf("not found")
	}
	r.m[entity.ID{UUID: id}] = nil
	return nil
}

func  (r *inmem) GetSale(id uuid.UUID) (*entity.Sale, error) {
	if r.m[entity.ID{UUID: id}] == nil {
		return nil, fmt.Errorf("not found")
	}
	return r.m[entity.ID{UUID: id}], nil
}

func  (r *inmem) List() ([]*entity.Sale, error) {
	var salesList []*entity.Sale
	for _, v := range r.m {
		salesList = append(salesList, v)
	}
	return salesList, nil
}


func ( r *inmem) GetSalesByUserId(user uuid.UUID)([]*entity.Sale, error) {
	var salesList []*entity.Sale
	for _, v := range r.m {
		if r.m[v.ID].User.UUID == user {
			salesList = append(salesList, r.m[v.ID])
		}
	}
	return salesList, nil
}
