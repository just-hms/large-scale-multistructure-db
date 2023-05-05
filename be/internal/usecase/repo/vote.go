package repo

import (
	"context"
	"fmt"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type VoteRepo struct {
	*mongo.Mongo
}

func NewVoteRepo(m *mongo.Mongo) *VoteRepo {
	return &VoteRepo{m}
}

func (v *VoteRepo) UpVoteByID(ctx context.Context, userID, shopID, reviewID string) error {

	shopFilter := bson.M{"_id": shopID}
	err := v.DB.Collection("barbershops").FindOne(ctx, shopFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified barber shop does not exist")
		}
		return err
	}

	reviewFilter := bson.M{"_id": shopID, "reviews": bson.M{"$elemMatch": bson.M{"reviewId": reviewID}}}
	err = v.DB.Collection("barbershops").FindOne(ctx, reviewFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified review does not exist")
		}
		return err
	}

	reviewFilter = bson.M{"_id": shopID, "reviews.reviewId": reviewID}

	// Check that Upvote wasn't already present
	upvoteFilter := bson.M{"_id": shopID, "reviews.reviewId": reviewID, "reviews.upvotes": bson.M{"$in": bson.A{userID}}}
	err = v.DB.Collection("barbershops").FindOne(ctx, upvoteFilter).Err()
	if err != mongo.ErrNoDocuments {
		return fmt.Errorf("user already upvoted")
	}

	// Check if user has already downvoted the review, and remove if it has
	downvoteUpdate := bson.M{"$pull": bson.M{"reviews.$[].downvotes": userID}}
	_, err = v.DB.Collection("barbershops").UpdateOne(ctx, reviewFilter, downvoteUpdate)
	if err != nil {
		return err
	}

	upvoteUpdate := bson.M{"$push": bson.M{"reviews.$[].upvotes": userID}}
	_, err = v.DB.Collection("barbershops").UpdateOne(ctx, reviewFilter, upvoteUpdate)

	return err

}

func (v *VoteRepo) DownVoteByID(ctx context.Context, userID, shopID, reviewID string) error {

	shopFilter := bson.M{"_id": shopID}
	err := v.DB.Collection("barbershops").FindOne(ctx, shopFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified barber shop does not exist")
		}
		return err
	}

	reviewFilter := bson.M{"_id": shopID, "reviews": bson.M{"$elemMatch": bson.M{"reviewId": reviewID}}}
	err = v.DB.Collection("barbershops").FindOne(ctx, reviewFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified review does not exist")
		}
		return err
	}

	reviewFilter = bson.M{"_id": shopID, "reviews.reviewId": reviewID}

	// Check that the Downvote wasn't already present
	downvoteFilter := bson.M{"_id": shopID, "reviews.reviewId": reviewID, "reviews.downvotes": bson.M{"$in": bson.A{userID}}}
	err = v.DB.Collection("barbershops").FindOne(ctx, downvoteFilter).Err()
	if err != mongo.ErrNoDocuments {
		return fmt.Errorf("user already downvoted")
	}

	// Check if user has already upvoted the review, and remove if it has
	upvoteUpdate := bson.M{"$pull": bson.M{"reviews.$[].upvotes": userID}}
	_, err = v.DB.Collection("barbershops").UpdateOne(ctx, reviewFilter, upvoteUpdate)
	if err != nil {
		return err
	}

	downvoteUpdate := bson.M{"$push": bson.M{"reviews.$[].downvotes": userID}}
	_, err = v.DB.Collection("barbershops").UpdateOne(ctx, reviewFilter, downvoteUpdate)

	return err
}

func (v *VoteRepo) RemoveVoteByID(ctx context.Context, userID, shopID, reviewID string) error {

	shopFilter := bson.M{"_id": shopID}
	err := v.DB.Collection("barbershops").FindOne(ctx, shopFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified barber shop does not exist")
		}
		return err
	}

	reviewFilter := bson.M{"_id": shopID, "reviews": bson.M{"$elemMatch": bson.M{"reviewId": reviewID}}}
	err = v.DB.Collection("barbershops").FindOne(ctx, reviewFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified review does not exist")
		}
		return err
	}

	reviewFilter = bson.M{"_id": shopID, "reviews.reviewId": reviewID}

	upvoteUpdate := bson.M{"$pull": bson.M{"reviews.$[].upvotes": userID}}
	_, err = v.DB.Collection("barbershops").UpdateOne(ctx, reviewFilter, upvoteUpdate)
	if err != nil {
		return err
	}

	downvoteUpdate := bson.M{"$pull": bson.M{"reviews.$[].downvotes": userID}}
	_, err = v.DB.Collection("barbershops").UpdateOne(ctx, reviewFilter, downvoteUpdate)
	return err

}
