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

func (r *AnalyticsRepo) GetAppointmentCountByShop(ctx context.Context, shopID string) ([]bson.M, error) {

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
			{"appointmentCount", bson.D{
				{"$sum", 1},
			}},
		},
	}}

	cur, err := r.DB.Collection("appointments").Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		return nil, err
	}

	var mongoResults []bson.M
	err = cur.All(ctx, &mongoResults)
	if err != nil {
		return nil, err
	}

	return mongoResults, err

}
