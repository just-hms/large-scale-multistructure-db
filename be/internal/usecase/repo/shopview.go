package repo

import (
	"context"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/pkg/mongo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type ShopViewRepo struct {
	*mongo.Mongo
}

func NewShopViewRepo(m *mongo.Mongo) *ShopViewRepo {
	return &ShopViewRepo{m}
}

func (r *ShopViewRepo) Store(ctx context.Context, view *entity.ShopView) error {
	view.CreatedAt = time.Now()

	filter := bson.M{"_id": view.BarberShopID}
	update := bson.M{"$push": bson.M{"views": view.ViewerID}}

	_, err := r.DB.Collection("barbershops").UpdateOne(ctx, filter, update)
	return err
}
