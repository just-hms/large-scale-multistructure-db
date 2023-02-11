package repo

import (
	"context"
	"fmt"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/pkg/mongo"
	"time"
)

type ShopViewRepo struct {
	*mongo.Mongo
}

func NewShopViewRepo(m *mongo.Mongo) *ShopViewRepo {
	return &ShopViewRepo{m}
}

func (r *ShopViewRepo) Store(ctx context.Context, view *entity.ShopView) error {

	view.CreatedAt = time.Now()

	_, err := r.DB.Collection("views").InsertOne(ctx, view)
	if err != nil {
		return fmt.Errorf("Error inserting the user")
	}

	return nil
}
