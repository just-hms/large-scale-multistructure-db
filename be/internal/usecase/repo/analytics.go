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

func (r *AnalyticsRepo) GetAppointmentViewReviewCount(ctx context.Context, shopID string) (map[string]int, error) {

	matchStage := bson.D{{"$match", bson.D{{"_id", shopID}}}}

	/*
		groupStage := bson.D{{
			"$group",
			bson.D{
				{"_id", bson.D{
					{"$dateToString", bson.D{
						{"date", "$startDate"},
						{"format", "%Y-%m"},
					}},
				}},
			},
		}}
	*/

	projectStage := bson.D{{
		"$project",
		bson.D{
			{"appointmentCount", bson.D{
				{"$size", bson.D{
					{"$ifNull", bson.A{
						"$appointments", bson.A{},
					}},
				}},
			}},
			{"viewCount", bson.D{
				{"$size", bson.D{
					{"$ifNull", bson.A{
						"$views", bson.A{},
					}},
				}},
			}},
			{"reviewCount", bson.D{
				{"$size", bson.D{
					{"$ifNull", bson.A{
						"$reviews", bson.A{},
					}},
				}},
			}},
			{"_id", 0},
		},
	}}

	cur, err := r.DB.Collection("barbershops").Aggregate(ctx, mongo.Pipeline{matchStage, projectStage})
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
		for k, v := range result {
			results[k] = int(v.(int32))
		}
	}

	return results, err

}
