package repo

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
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

func (r *AnalyticsRepo) GetUpDownVoteCountByShop(ctx context.Context, shopID string) (map[string]map[string]int, error) {

	matchStage := bson.D{{"$match", bson.D{{"shopId", shopID}}}}

	projectStage := bson.D{{
		"$project",
		bson.D{
			{"createdAt", 1},
			{"upCount", bson.D{
				{"$size", "$upvotes"},
			}},
			{"downCount", bson.D{
				{"$size", "$downvotes"},
			}},
		},
	}}

	groupStage := bson.D{{
		"$group",
		bson.D{
			{"_id", bson.D{
				{"$dateToString", bson.D{
					{"date", "$createdAt"},
					{"format", "%Y-%m"},
				}},
			}},
			{"upCount", bson.D{
				{"$sum", "$upCount"},
			}},
			{"downCount", bson.D{
				{"$sum", "$downCount"},
			}},
		},
	}}

	cur, err := r.DB.Collection("reviews").Aggregate(ctx, mongo.Pipeline{matchStage, projectStage, groupStage})
	if err != nil {
		return nil, err
	}

	results := make(map[string]map[string]int)
	var mongoResults []bson.M
	err = cur.All(ctx, &mongoResults)
	if err != nil {
		return nil, err
	}

	for _, result := range mongoResults {
		results[result["_id"].(string)] = make(map[string]int)
		results[result["_id"].(string)]["upCount"] = int(result["upCount"].(int32))
		results[result["_id"].(string)]["downCount"] = int(result["downCount"].(int32))
	}

	return results, err

}

func (r *AnalyticsRepo) GetWeightRankedReviewByShop(ctx context.Context, shopID string) ([]*entity.Review, error) {

	matchStage := bson.D{{"$match", bson.D{{"shopId", shopID}}}}

	setStage1 := bson.D{{
		"$set",
		bson.D{
			{"upCount", bson.D{
				{"$size", "$upvotes"},
			}},
			{"downCount", bson.D{
				{"$size", "$downvotes"},
			}},
			{"daysElapsed", bson.D{
				{"$dateDiff", bson.D{
					{"startDate", "$createdAt"},
					{"endDate", "$$NOW"},
					{"unit", "day"},
				}},
			}},
		},
	}}

	setStage2 := bson.D{{
		"$set",
		bson.D{
			{"freshnessScore", bson.D{
				{"$switch", bson.D{
					{"branches", bson.A{
						bson.D{
							{"case", bson.D{
								{"$and", bson.A{
									bson.D{{"$gte", bson.A{"$daysElapsed", 0}}},
									bson.D{{"$lt", bson.A{"$daysElapsed", 30}}},
								}},
							}},
							{"then", 5},
						},
						bson.D{
							{"case", bson.D{
								{"$and", bson.A{
									bson.D{{"$gte", bson.A{"$daysElapsed", 30}}},
									bson.D{{"$lt", bson.A{"$daysElapsed", 365}}},
								}},
							}},
							{"then", 2},
						},
					}},
					{"default", 1},
				}},
			}},
			{"voteScore", bson.D{
				{"$cond", bson.A{
					bson.D{{"$eq", bson.A{bson.D{{"$subtract", bson.A{"$upCount", "$downCount"}}}, 0}}},
					1,
					bson.D{{"$subtract", bson.A{"$upCount", "$downCount"}}},
				}},
			}},
		},
	}}

	setStage3 := bson.D{{
		"$set",
		bson.D{
			{"weightedScore", bson.D{
				{"$multiply", bson.A{"$freshnessScore", "$voteScore"}},
			}},
		},
	}}

	unsetStage := bson.D{{
		"$unset",
		bson.A{"upCount", "downCount", "freshnessScore", "voteScore"},
	}}

	sortStage := bson.D{{
		"$sort",
		bson.D{
			{"weightedScore", -1},
		},
	}}

	cur, err := r.DB.Collection("reviews").Aggregate(ctx, mongo.Pipeline{matchStage, setStage1, setStage2, setStage3, unsetStage, sortStage})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	results := []*entity.Review{}

	for cur.Next(ctx) {
		var review entity.Review

		if err := cur.Decode(&review); err != nil {
			return nil, err
		}
		results = append(results, &review)
	}

	return results, err

}
