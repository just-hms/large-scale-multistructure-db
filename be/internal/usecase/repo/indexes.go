package repo

import (
	"context"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func AddTestIndexes(mongo *mongo.Mongo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return CreateTestIndexes(mongo, ctx)
}

func CreateTestIndexes(m *mongo.Mongo, ctx context.Context) error {

	indexOpts := options.CreateIndexes().SetMaxTime(time.Second * 10)

	//Make email and username unique
	usernameIndexModel := mongodriver.IndexModel{
		Keys:    bson.D{{"username", 1}},
		Options: options.Index().SetUnique(true),
	}
	emailIndexModel := mongodriver.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	}
	// Index to location 2dsphere type.
	pointIndexModel := mongodriver.IndexModel{
		Keys: bsonx.MDoc{"location": bsonx.String("2dsphere")},
	}

	userIndexes := m.DB.Collection("users").Indexes()
	_, err := userIndexes.CreateOne(ctx, usernameIndexModel, indexOpts)
	if err != nil {
		return err
	}
	_, err = userIndexes.CreateOne(ctx, emailIndexModel, indexOpts)
	if err != nil {
		return err
	}

	shopIndexes := m.DB.Collection("barbershops").Indexes()
	_, err = shopIndexes.CreateOne(ctx, pointIndexModel, indexOpts)
	return err
}
