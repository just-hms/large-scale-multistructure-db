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

func (r *AdminAnalyticsRepo) GetAppointmentCount(ctx context.Context) (map[string]int, error) {

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

func (r *AdminAnalyticsRepo) GetViewCount(ctx context.Context) (map[string]int, error) {

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

func (r *AdminAnalyticsRepo) GetReviewCount(ctx context.Context) (map[string]int, error) {

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

func (r *AdminAnalyticsRepo) GetNewUsersCount(ctx context.Context) (map[string]int, error) {

	groupStage := bson.D{{
		"$group",
		bson.D{
			{"_id", bson.D{
				{"$dateToString", bson.D{
					{"date", "$signupDate"},
					{"format", "%Y-%m"},
				}},
			}},
			{"count", bson.D{
				{"$sum", 1},
			}},
		},
	}}

	cur, err := r.DB.Collection("users").Aggregate(ctx, mongo.Pipeline{groupStage})
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

func (r *AdminAnalyticsRepo) GetAppointmentCancellationUserRanking(ctx context.Context) ([]bson.M, error) {

	setStage := bson.D{{
		"$set",
		bson.D{
			{"isCanceled", bson.D{
				{"$cond", bson.A{
					bson.D{{"$eq", bson.A{"$status", "canceled"}}},
					1,
					0,
				}},
			}},
		},
	}}

	groupStage := bson.D{{
		"$group",
		bson.D{
			{"_id", "$username"},
			{"cancelCount", bson.D{
				{"$sum", "$isCanceled"},
			}},
		},
	}}

	projectStage := bson.D{{
		"$project",
		bson.D{
			{"_id", 0},
			{"username", "$_id"},
			{"cancelCount", 1},
		},
	}}

	sortStage := bson.D{{
		"$sort",
		bson.D{
			{"cancelCount", -1},
		},
	}}

	cur, err := r.DB.Collection("appointments").Aggregate(ctx, mongo.Pipeline{setStage, groupStage, projectStage, sortStage})
	if err != nil {
		return nil, err
	}

	var mongoResults []bson.M
	err = cur.All(ctx, &mongoResults)
	if err != nil {
		return nil, err
	}

	for _, result := range mongoResults {
		result["cancelCount"] = int(result["cancelCount"].(int32))
	}

	return mongoResults, err

}

func (r *AdminAnalyticsRepo) GetAppointmentCancellationShopRanking(ctx context.Context) ([]bson.M, error) {

	setStage := bson.D{{
		"$set",
		bson.D{
			{"isCanceled", bson.D{
				{"$cond", bson.A{
					bson.D{{"$eq", bson.A{"$status", "canceled"}}},
					1,
					0,
				}},
			}},
		},
	}}

	groupStage := bson.D{{
		"$group",
		bson.D{
			{"_id", "$shopId"},
			{"shopName", bson.D{
				{"$first", "$shopName"},
			}},
			{"cancelCount", bson.D{
				{"$sum", "$isCanceled"},
			}},
		},
	}}

	projectStage := bson.D{{
		"$project",
		bson.D{
			{"_id", 0},
			{"shopName", "$shopName"},
			{"cancelCount", 1},
		},
	}}

	sortStage := bson.D{{
		"$sort",
		bson.D{
			{"cancelCount", -1},
		},
	}}

	cur, err := r.DB.Collection("appointments").Aggregate(ctx, mongo.Pipeline{setStage, groupStage, projectStage, sortStage})
	if err != nil {
		return nil, err
	}

	var mongoResults []bson.M
	err = cur.All(ctx, &mongoResults)
	if err != nil {
		return nil, err
	}

	for _, result := range mongoResults {
		result["cancelCount"] = int(result["cancelCount"].(int32))
	}

	return mongoResults, err

}

func (r *AdminAnalyticsRepo) GetEngagementShopRanking(ctx context.Context) ([]bson.M, error) {

	setStage := bson.D{{
		"$set",
		bson.D{
			{"upCount", bson.D{
				{"$size", "$upvotes"},
			}},
			{"downCount", bson.D{
				{"$size", "$downvotes"},
			}},
		},
	}}

	setStage2 := bson.D{{
		"$set",
		bson.D{
			{"voteEngagement", bson.D{
				{"$sum", bson.A{"$upCount", "$downCount"}},
			}},
		},
	}}

	groupStage := bson.D{{
		"$group",
		bson.D{
			{"_id", "$shopId"},
			{"shopName", bson.D{
				{"$first", "$shopName"},
			}},
			{"voteEngagement", bson.D{
				{"$sum", "$voteEngagement"},
			}},
		},
	}}

	lookupGroupAndScoreAppointmentsStage := bson.D{{
		"$group",
		bson.D{
			{"_id", "$shopId"},
			{"appointmentEngagement", bson.D{
				{"$sum", 5},
			}},
		},
	}}

	lookupScoreAppointmentsPipeline := bson.A{lookupGroupAndScoreAppointmentsStage}

	lookupScoreAppointmentsStage := bson.D{{
		"$lookup", bson.D{
			{"from", "appointments"},
			{"localField", "_id"},
			{"foreignField", "shopId"},
			{"pipeline", lookupScoreAppointmentsPipeline},
			{"as", "appointmentEngagementList"},
		},
	}}

	setStage3 := bson.D{{
		"$set",
		bson.D{
			{"engagementScoreElem", bson.D{
				{"$arrayElemAt", bson.A{"$appointmentEngagementList", 0}},
			}},
		},
	}}

	setStage4 := bson.D{{
		"$set",
		bson.D{
			{"engagementScore", bson.D{
				{"$add", bson.A{"$voteEngagement", "$engagementScoreElem.appointmentEngagement"}},
			}},
		},
	}}

	projectStage := bson.D{{
		"$project",
		bson.D{
			{"_id", 0},
			{"shopName", "$shopName"},
			{"engagementScore", 1},
		},
	}}

	sortStage := bson.D{{
		"$sort",
		bson.D{
			{"engagementScore", -1},
		},
	}}

	cur, err := r.DB.Collection("reviews").Aggregate(ctx, mongo.Pipeline{setStage, setStage2, groupStage, lookupScoreAppointmentsStage, setStage3, setStage4, projectStage, sortStage})
	if err != nil {
		return nil, err
	}

	var mongoResults []bson.M
	err = cur.All(ctx, &mongoResults)
	if err != nil {
		return nil, err
	}

	for _, result := range mongoResults {
		result["engagementScore"] = int(result["engagementScore"].(int32))
	}

	return mongoResults, err

}
