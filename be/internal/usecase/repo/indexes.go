package repo

import (
	"context"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func AddIndexes(mongo *mongo.Mongo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return CreatePointIndex(mongo, ctx, "barbershops")
}

// TODO: make this more general
func CreatePointIndex(m *mongo.Mongo, ctx context.Context, collection string) error {

	indexOpts := options.CreateIndexes().SetMaxTime(time.Second * 10)

	// Index to location 2dsphere type.
	pointIndexModel := mongodriver.IndexModel{
		Keys: bsonx.MDoc{"location": bsonx.String("2dsphere")},
	}

	pointIndexes := m.DB.Collection(collection).Indexes()
	_, err := pointIndexes.CreateOne(ctx, pointIndexModel, indexOpts)
	return err
}
