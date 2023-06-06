package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
)

type ShopViewRepo struct {
	*mongo.Mongo
}

func NewShopViewRepo(m *mongo.Mongo) *ShopViewRepo {
	return &ShopViewRepo{m}
}

func (r *ShopViewRepo) Store(ctx context.Context, view *entity.ShopView) error {

	view.CreatedAt = time.Now()
	view.ID = uuid.NewString()
	_, err := r.DB.Collection("shopviews").InsertOne(ctx, view)
	if err != nil {
		view.ID = ""
		return fmt.Errorf("error inserting the shopview: %s", err.Error())
	}

	return err
}
