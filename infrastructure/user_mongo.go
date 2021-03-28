package infrastructure

import (
	"StoreManager-DDD/config"
	"StoreManager-DDD/entity"
	"StoreManager-DDD/usecase/user"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/juju/mgosession"
	"gopkg.in/mgo.v2/bson"
)

type repo struct {
	pool *mgosession.Pool
}

func NewMongoRepository(p *mgosession.Pool) user.Repository {
	return &repo{
		pool: p,
	}
}

func (r repo) Get(id uuid.UUID) (*entity.User, error) {
	result := entity.User{}
	session := r.pool.Session(nil)
	coll := session.DB(config.DB_NAME).C("user")
	err := coll.FindId(&entity.ID{UUID: id}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (r repo) GetByEmail(email string) (*entity.User, error) {
	result := entity.User{}
	session := r.pool.Session(nil)
	coll := session.DB(config.DB_NAME).C("user")
	err := coll.Find(bson.M{"email": email}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (r repo) List() ([]*entity.User, error) {
	var userList []*entity.User
	session := r.pool.Session(nil)
	coll := session.DB(config.DB_NAME).C("user")
	err := coll.Find(bson.M{}).All(&userList)
	if err != nil {
		return nil, err
	}
	return userList, err
}

func (r repo) Create(e *entity.User) (uuid.UUID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(config.DB_NAME).C("user")
	user_, _ := r.GetByEmail(e.Email)
	if user_ != nil {
		return uuid.Nil, errors.New("user already exists")
	}
	err := coll.Insert(e)
	if err != nil {
		return uuid.Nil, errors.New("error on user insertion")
	}
	return e.ID.UUID, err
}

func (r repo) Update(e *entity.User) error {
	session := r.pool.Session(nil)
	coll := session.DB(config.DB_NAME).C("user")
	err := coll.UpdateId(e.ID, e)
	if err != nil {
		fmt.Println(err)
		return errors.New("error occurred during update")
	}
	return nil
}

func (r repo) Delete(id uuid.UUID) error {
	panic("implement me")
}


