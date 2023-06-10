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
	review.ShopID = shopID
	review.Username = user.Username
	review.Reported = false
	review.UpVotes = []string{}
	review.DownVotes = []string{}
	if review.CreatedAt.IsZero() {
		review.CreatedAt = time.Now()
	}

	_, err = r.DB.Collection("reviews").InsertOne(ctx, review)
	if err != nil {
		review.ID = ""
		return fmt.Errorf("error inserting the review: %s", err.Error())
	}

	return err
}

func (r *ReviewRepo) GetByBarberShopID(ctx context.Context, shopID string) ([]*entity.Review, error) {

	err := r.DB.Collection("barbershops").FindOne(ctx, bson.M{"_id": shopID}).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("the specified barber shop does not exist")
		}
		return nil, err
	}

	cur, err := r.DB.Collection("reviews").Find(ctx, bson.M{"shopId": shopID})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	reviews := []*entity.Review{}

	for cur.Next(ctx) {
		var review entity.Review

		if err := cur.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}
	return reviews, nil
}

func (r *ReviewRepo) DeleteByID(ctx context.Context, reviewID string) error {

	res, err := r.DB.Collection("reviews").DeleteOne(ctx, bson.M{"_id": reviewID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("review not found")
	}
	return nil
}
