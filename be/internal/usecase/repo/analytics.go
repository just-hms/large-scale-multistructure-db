package repo

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type AnalyticsRepo struct {
	*mongo.Mongo
}

func NewAnalyticsRepo(m *mongo.Mongo) *AnalyticsRepo {
	return &AnalyticsRepo{m}
}

func (r *AnalyticsRepo) GetAppointmentCountByShop(ctx context.Context, shopID string) (map[string]int, error) {

	matchStage := bson.D{{"$match", bson.D{{"shopId", shopID}}}}

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

	cur, err := r.DB.Collection("appointments").Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
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

func (r *AnalyticsRepo) GetViewCountByShop(ctx context.Context, shopID string) (map[string]int, error) {

	matchStage := bson.D{{"$match", bson.D{{"shopId", shopID}}}}

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

	cur, err := r.DB.Collection("shopviews").Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
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

func (r *AnalyticsRepo) GetReviewCountByShop(ctx context.Context, shopID string) (map[string]int, error) {

	matchStage := bson.D{{"$match", bson.D{{"shopId", shopID}}}}

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

	cur, err := r.DB.Collection("reviews").Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
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

func (r *AnalyticsRepo) GetAppointmentViewRatioByShop(ctx context.Context, shopID string) (map[string]float64, error) {

	viewCount, err := r.GetViewCountByShop(ctx, shopID)
	if err != nil {
		return nil, err
	}

	appointmentCount, err := r.GetAppointmentCountByShop(ctx, shopID)
	if err != nil {
		return nil, err
	}

	results := make(map[string]float64)
	for month, vCount := range viewCount {
		aCount, ok := appointmentCount[month]
		if ok {
			results[month] = float64(aCount) / float64(vCount)
		} else {
			results[month] = 0.0
		}
	}

	return results, err

}