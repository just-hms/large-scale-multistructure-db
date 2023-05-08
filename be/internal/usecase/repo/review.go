package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type ReviewRepo struct {
	*mongo.Mongo
}

func NewReviewRepo(m *mongo.Mongo) *ReviewRepo {
	return &ReviewRepo{m}
}

func (r *ReviewRepo) Store(ctx context.Context, review *entity.Review, shopID string) error {

	err := r.DB.Collection("barbershops").FindOne(ctx, bson.M{"_id": shopID}).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified barber shop does not exist")
		}
		return err
	}

	user := &entity.User{}
	err = r.DB.Collection("users").FindOne(ctx, bson.M{"_id": review.UserID}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified user does not exist")
		}
		return err
	}

	review.ID = uuid.NewString()
	review.Username = user.Username
	review.Reported = false
	review.UpVotes = []string{}
	review.DownVotes = []string{}
	review.CreatedAt = time.Now()

	filter := bson.M{"_id": shopID}
	update := bson.M{"$push": bson.M{"reviews": review}}

	_, err = r.DB.Collection("barbershops").UpdateOne(ctx, filter, update)
	return err
}

func (r *ReviewRepo) GetByBarberShopID(ctx context.Context, shopID string) ([]*entity.Review, error) {

	shop := &entity.BarberShop{}
	err := r.DB.Collection("barbershops").FindOne(ctx, bson.M{"_id": shopID}).Decode(&shop)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("the specified barber shop does not exist")
		}
		return nil, err
	}

	return shop.Reviews, nil
}

func (r *ReviewRepo) DeleteByID(ctx context.Context, shopID, reviewID string) error {

	shopFilter := bson.M{"_id": shopID}
	err := r.DB.Collection("barbershops").FindOne(ctx, shopFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified barber shop does not exist")
		}
		return err
	}

	reviewFilter := bson.M{"_id": shopID, "reviews": bson.M{"$elemMatch": bson.M{"_id": reviewID}}}
	err = r.DB.Collection("barbershops").FindOne(ctx, reviewFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified review does not exist")
		}
		return err
	}

	update := bson.M{"$pull": bson.M{"reviews": bson.M{"_id": reviewID}}}

	_, err = r.DB.Collection("barbershops").UpdateOne(ctx, shopFilter, update)

	return err
}
