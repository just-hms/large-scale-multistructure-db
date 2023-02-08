package repo

import (
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/pkg/mongo"
)

type BarberShopRepo struct {
	mongo *mongo.Mongo
}

func NewBarberShopRepo(m *mongo.Mongo) *UserRepo {
	return &UserRepo{m}
}

func (r *BarberShopRepo) Find(lat float64, lon float64, name string, radius float64) ([]*entity.BarberShop, error) {
	// TODO : this
	return nil, nil

}
