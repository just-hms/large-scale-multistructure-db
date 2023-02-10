package repo

import (
	"context"
	"fmt"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/pkg/mongo"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShopViewRepo struct {
	*mongo.Mongo
}

func NewShopViewRepo(m *mongo.Mongo) *ShopViewRepo {
	return &ShopViewRepo{m}
}

func (r *ShopViewRepo) Store(ctx context.Context, view *entity.ShopView) (string, error) {
	view.CreatedAt = time.Now()

	res, err := r.DB.Collection("views").InsertOne(ctx, view)
	if err != nil {
		return "", fmt.Errorf("Error inserting the user")
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("Error retriving the userID")
	}

	return oid.Hex(), err
}
