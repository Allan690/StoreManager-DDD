package infrastructure

import (
	"StoreManager-DDD/config"
	"StoreManager-DDD/entity"
	"github.com/gofrs/uuid"
	"github.com/juju/mgosession"
	"gopkg.in/mgo.v2/bson"
)

type salesRepo struct {
	pool *mgosession.Pool
}

func NewMongoSalesRepo(p *mgosession.Pool) *salesRepo{
	return &salesRepo{
		pool: p,
	}
}

// GetSale fetches a sale by id from mongo db
func (repo *salesRepo) GetSale(id uuid.UUID) (*entity.Sale, error)  {
	result := entity.Sale{}
	session := repo.pool.Session(nil)
	coll := session.DB(config.DB_NAME).C("sale")
	err := coll.FindId(entity.ID{UUID: id}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List fetches all sales from mongo db
func (repo *salesRepo) List() ([]*entity.Sale, error)  {
	var results []*entity.Sale
	session := repo.pool.Session(nil)
	coll := session.DB(config.DB_NAME).C("sale")
	err := coll.Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// CreateSale creates a sale in mongo db
func (repo *salesRepo) CreateSale(sale *entity.Sale) (*entity.Sale, error) {
	session := repo.pool.Session(nil)
	coll := session.DB(config.DB_NAME).C("sale")
	err := coll.Insert(sale)
	if err != nil {
		return nil, err
	}
	return sale, nil
}


