package repo

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type BarberAnalyticsRepo struct {
	*mongo.Mongo
}

func NewBarberAnalyticsRepo(m *mongo.Mongo) *BarberAnalyticsRepo {
	return &BarberAnalyticsRepo{m}
}

func (r *BarberAnalyticsRepo) GetAppointmentCountByShop(ctx context.Context, shopID string) (map[string]int, error) {

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

func (r *BarberAnalyticsRepo) GetViewCountByShop(ctx context.Context, shopID string) (map[string]int, error) {

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

func (r *BarberAnalyticsRepo) GetReviewCountByShop(ctx context.Context, shopID string) (map[string]int, error) {

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

func (r *BarberAnalyticsRepo) GetAppointmentViewRatioByShop(ctx context.Context, shopID string) (map[string]float64, error) {

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

func (r *BarberAnalyticsRepo) GetUpDownVoteCountByShop(ctx context.Context, shopID string) (map[string]map[string]int, error) {

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

// This aggregation produces a weighted rating of a Shop based on its Reviews.
// Reviews are weighted depending on Freshness and VoteScore
// Freshness:
//
//		created < 30 days -> 5 points
//	 30 days <= created < 365 days -> 2 points
//	 created > 365 days -> 1 point
//
// VoteScore: #upvotes - #downvotes
// WeightedScore: freshness * voteScore
// WeightedRating: (weightedScore * rating) / sum(weightedScore)
func (r *BarberAnalyticsRepo) GetReviewWeightedRatingByShop(ctx context.Context, shopID string) (float64, error) {

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

	groupStage := bson.D{{
		"$group",
		bson.D{
			{"_id", "$shopId"},
			{"numerator", bson.D{
				{"$sum", bson.D{
					{"$multiply", bson.A{"$weightedScore", "$rating"}},
				}},
			}},
			{"denominator", bson.D{
				{"$sum", "$weightedScore"},
			}},
		},
	}}

	projectStage := bson.D{{
		"$project",
		bson.D{
			{"_id", 0},
			{"weightedRating", bson.D{
				{"$trunc", bson.A{
					bson.D{{"$divide", bson.A{"$numerator", "$denominator"}}},
					2,
				}},
			}},
		},
	}}

	cur, err := r.DB.Collection("reviews").Aggregate(ctx, mongo.Pipeline{matchStage, setStage1, setStage2, setStage3, groupStage, projectStage})
	if err != nil {
		return 0, err
	}

	var mongoResults []bson.M
	err = cur.All(ctx, &mongoResults)
	if err != nil {
		return 0, err
	}

	result := mongoResults[0]["weightedRating"].(float64)

	return result, err

}

// This aggregation is quite complex and can be basically broken into 4 steps:
// - Use a Replace root to get just a single doc with the shopId
// - Find all the users that made an appointment in the Shop in the last 90 days
// - Find all the users that made an appointment in the last 90 days in another Shop and weren't in the new users (active users in other Shops)
// - Find all the users that made an appointment in the past and are active in other Shops
func (r *BarberAnalyticsRepo) GetInactiveUsersByShop(ctx context.Context, shopID string) ([]string, error) {

	matchStage1 := bson.D{{
		"$match", bson.D{
			{"shopId", shopID},
		},
	}}

	groupStage1 := bson.D{{
		"$group",
		bson.D{
			{"_id", "$shopId"},
			{"doc", bson.D{
				{"$first", "$$ROOT"},
			}},
		},
	}}

	replaceRootStage1 := bson.D{{
		"$replaceRoot",
		bson.D{
			{"newRoot", "$doc"},
		},
	}}

	lookupNewClientsShopPipelineMatchStage1 := bson.D{{
		"$match", bson.D{
			{"shopId", shopID},
			{"status", bson.D{
				{"$ne", "canceled"},
			}},
		},
	}}

	lookupNewClientsShopPipelineSetStage1 := bson.D{{
		"$set",
		bson.D{
			{"daysElapsed", bson.D{
				{"$dateDiff", bson.D{
					{"startDate", "$startDate"},
					{"endDate", "$$NOW"},
					{"unit", "day"},
				}},
			}},
		},
	}}

	lookupNewClientsShopPipelineMatchStage2 := bson.D{{
		"$match", bson.D{
			{"$expr", bson.D{
				{"$lt", bson.A{"$daysElapsed", 90}},
			}},
		},
	}}

	lookupNewClientsShopPipelineProjectStage := bson.D{{
		"$project", bson.D{
			{"_id", 0},
			{"userId", 1},
		},
	}}

	lookupNewClientsShopPipeline := bson.A{lookupNewClientsShopPipelineMatchStage1, lookupNewClientsShopPipelineSetStage1, lookupNewClientsShopPipelineMatchStage2, lookupNewClientsShopPipelineProjectStage}

	lookupNewClientsShopStage := bson.D{{
		"$lookup", bson.D{
			{"from", "appointments"},
			{"pipeline", lookupNewClientsShopPipeline},
			{"as", "newClientsShop"},
		},
	}}

	lookupNewClientsNoShopPipelineMatchStage1 := bson.D{{
		"$match", bson.D{
			{"shopId", bson.D{
				{"$ne", shopID},
			}},
			{"status", bson.D{
				{"$ne", "canceled"},
			}},
		},
	}}

	lookupNewClientsNoShopPipelineMatchStage2 := bson.D{{
		"$match", bson.D{
			{"$expr", bson.D{
				{"$not", bson.D{
					{"$in", bson.A{"$userId", "$$newClientsShop.userId"}},
				}},
			}},
		},
	}}

	lookupNewClientsNoShopPipeline := bson.A{lookupNewClientsNoShopPipelineMatchStage1, lookupNewClientsShopPipelineSetStage1, lookupNewClientsShopPipelineMatchStage2, lookupNewClientsNoShopPipelineMatchStage2, lookupNewClientsShopPipelineProjectStage}

	lookupNewClientsNoShopStage := bson.D{{
		"$lookup", bson.D{
			{"from", "appointments"},
			{"let", bson.D{
				{"newClientsShop", "$newClientsShop"},
			}},
			{"pipeline", lookupNewClientsNoShopPipeline},
			{"as", "newClientsNoShop"},
		},
	}}

	lookupOldClientsShopPipelineMatchStage1 := bson.D{{
		"$match", bson.D{
			{"$expr", bson.D{
				{"$gte", bson.A{"$daysElapsed", 90}},
			}},
		},
	}}

	lookupOldClientsShopPipelineMatchStage2 := bson.D{{
		"$match", bson.D{
			{"$expr", bson.D{
				{"$in", bson.A{"$userId", "$$newClientsNoShop.userId"}},
			}},
		},
	}}

	lookupOldClientsShopPipeline := bson.A{lookupNewClientsShopPipelineMatchStage1, lookupNewClientsShopPipelineSetStage1, lookupOldClientsShopPipelineMatchStage1, lookupOldClientsShopPipelineMatchStage2, lookupNewClientsShopPipelineProjectStage}

	lookupOldClientsShopStage := bson.D{{
		"$lookup", bson.D{
			{"from", "appointments"},
			{"let", bson.D{
				{"newClientsNoShop", "$newClientsNoShop"},
			}},
			{"pipeline", lookupOldClientsShopPipeline},
			{"as", "oldClientsShop"},
		},
	}}

	projectStage1 := bson.D{{
		"$project",
		bson.D{
			{"_id", 0},
			{"oldClientsShop", 1},
		},
	}}

	cur, err := r.DB.Collection("appointments").Aggregate(ctx, mongo.Pipeline{matchStage1, groupStage1, replaceRootStage1, lookupNewClientsShopStage, lookupNewClientsNoShopStage, lookupOldClientsShopStage, projectStage1})
	if err != nil {
		return nil, err
	}

	results := []string{}
	var mongoResults []bson.M
	err = cur.All(ctx, &mongoResults)
	if err != nil {
		return nil, err
	}

	userIdMapList := mongoResults[0]["oldClientsShop"].(bson.A)

	for _, user := range userIdMapList {
		results = append(results, user.(bson.M)["userId"].(string))
	}

	return results, err

}
