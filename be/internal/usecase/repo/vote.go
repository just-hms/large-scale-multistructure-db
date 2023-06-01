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

	err := v.DB.Collection("barbershops").FindOne(ctx, bson.M{"_id": shopID}).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified barber shop does not exist")
		}
		return err
	}

	reviewFilter := bson.M{"_id": reviewID}
	err = v.DB.Collection("reviews").FindOne(ctx, reviewFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified review does not exist")
		}
		return err
	}

	// Check that Upvote wasn't already present
	upvoteFilter := bson.M{"_id": reviewID, "upvotes": bson.M{"$in": bson.A{userID}}}
	err = v.DB.Collection("reviews").FindOne(ctx, upvoteFilter).Err()
	if err != mongo.ErrNoDocuments {
		return fmt.Errorf("user already upvoted")
	}

	// Check if user has already downvoted the review, and remove if it has
	downvoteUpdate := bson.M{"$pull": bson.M{"downvotes": userID}}
	_, err = v.DB.Collection("reviews").UpdateOne(ctx, reviewFilter, downvoteUpdate)
	if err != nil {
		return err
	}

	upvoteUpdate := bson.M{"$push": bson.M{"upvotes": userID}}
	_, err = v.DB.Collection("reviews").UpdateOne(ctx, reviewFilter, upvoteUpdate)

	return err

}

func (v *VoteRepo) DownVoteByID(ctx context.Context, userID, shopID, reviewID string) error {

	err := v.DB.Collection("barbershops").FindOne(ctx, bson.M{"_id": shopID}).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified barber shop does not exist")
		}
		return err
	}

	reviewFilter := bson.M{"_id": reviewID}
	err = v.DB.Collection("reviews").FindOne(ctx, reviewFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified review does not exist")
		}
		return err
	}

	// Check that the Downvote wasn't already present
	downvoteFilter := bson.M{"_id": reviewID, "downvotes": bson.M{"$in": bson.A{userID}}}
	err = v.DB.Collection("reviews").FindOne(ctx, downvoteFilter).Err()
	if err != mongo.ErrNoDocuments {
		return fmt.Errorf("user already downvoted")
	}

	// Check if user has already upvoted the review, and remove if it has
	upvoteUpdate := bson.M{"$pull": bson.M{"upvotes": userID}}
	_, err = v.DB.Collection("reviews").UpdateOne(ctx, reviewFilter, upvoteUpdate)
	if err != nil {
		return err
	}

	downvoteUpdate := bson.M{"$push": bson.M{"downvotes": userID}}
	_, err = v.DB.Collection("reviews").UpdateOne(ctx, reviewFilter, downvoteUpdate)

	return err
}

func (v *VoteRepo) RemoveVoteByID(ctx context.Context, userID, shopID, reviewID string) error {

	reviewFilter := bson.M{"_id": reviewID}
	err := v.DB.Collection("reviews").FindOne(ctx, reviewFilter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("the specified review does not exist")
		}
		return err
	}

	upvoteUpdate := bson.M{"$pull": bson.M{"upvotes": userID}}
	_, err = v.DB.Collection("reviews").UpdateOne(ctx, reviewFilter, upvoteUpdate)
	if err != nil {
		return err
	}

	downvoteUpdate := bson.M{"$pull": bson.M{"downvotes": userID}}
	_, err = v.DB.Collection("reviews").UpdateOne(ctx, reviewFilter, downvoteUpdate)
	return err

}
