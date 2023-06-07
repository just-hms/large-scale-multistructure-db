package repo

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type AdminAnalyticsRepo struct {
	*mongo.Mongo
}

func NewAdminAnalyticsRepo(m *mongo.Mongo) *AdminAnalyticsRepo {
	return &AdminAnalyticsRepo{m}
}

func (r *AdminAnalyticsRepo) GetAppointmentCount(ctx context.Context, shopID string) (map[string]int, error) {

	groupStage := bson.D{{
		"$group",
		bson.D{
			{"_id", bson.D{
				{"$dateToString", bson.D{
					{"date", "$startDate"},
					{"format", "%Y-%m"},
				}},
			}},
			{"count", bson.D{
				{"$sum", 1},
			}},
		},
	}}

	cur, err := r.DB.Collection("appointments").Aggregate(ctx, mongo.Pipeline{groupStage})
	if err != nil {
		return nil, err
	}

	results := make(map[string]int)
	var mongoResults []bson.M
	err = cur.All(ctx, &mongoResults)
	if err != nil {
		return nil, err
	}

	for _, result := range mongoResults {
		results[result["_id"].(string)] = int(result["count"].(int32))
	}

	return results, err

}

func (r *AdminAnalyticsRepo) GetViewCount(ctx context.Context, shopID string) (map[string]int, error) {

	groupStage := bson.D{{
		"$group",
		bson.D{
			{"_id", bson.D{
				{"$dateToString", bson.D{
					{"date", "$createdAt"},
					{"format", "%Y-%m"},
				}},
			}},
			{"count", bson.D{
				{"$sum", 1},
			}},
		},
	}}

	cur, err := r.DB.Collection("shopviews").Aggregate(ctx, mongo.Pipeline{groupStage})
	if err != nil {
		return nil, err
	}

	results := make(map[string]int)
	var mongoResults []bson.M
	err = cur.All(ctx, &mongoResults)
	if err != nil {
		return nil, err
	}

	for _, result := range mongoResults {
		results[result["_id"].(string)] = int(result["count"].(int32))
	}

	return results, err

}

func (r *AdminAnalyticsRepo) GetReviewCount(ctx context.Context, shopID string) (map[string]int, error) {

	groupStage := bson.D{{
		"$group",
		bson.D{
			{"_id", bson.D{
				{"$dateToString", bson.D{
					{"date", "$createdAt"},
					{"format", "%Y-%m"},
				}},
			}},
			{"count", bson.D{
				{"$sum", 1},
			}},
		},
	}}

	cur, err := r.DB.Collection("reviews").Aggregate(ctx, mongo.Pipeline{groupStage})
	if err != nil {
		return nil, err
	}

	results := make(map[string]int)
	var mongoResults []bson.M
	err = cur.All(ctx, &mongoResults)
	if err != nil {
		return nil, err
	}

	for _, result := range mongoResults {
		results[result["_id"].(string)] = int(result["count"].(int32))
	}

	return results, err

}
