package repo

import (
	"context"
	"errors"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

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
	// add index to shopId to improve performance
	shopIDIndexModel := mongodriver.IndexModel{
		Keys:    bson.D{{"shopId", 1}},
		Options: options.Index(),
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
	if err != nil {
		return err
	}

	appointmentsIndexes := m.DB.Collection("appointments").Indexes()
	_, err = appointmentsIndexes.CreateOne(ctx, shopIDIndexModel, indexOpts)
	if err != nil {
		return err
	}

	return err
}

func DropTestIndexes(m *mongo.Mongo, ctx context.Context) error {
	_, err1 := m.DB.Collection("users").Indexes().DropOne(ctx, "username_1")
	_, err2 := m.DB.Collection("users").Indexes().DropOne(ctx, "email_1")
	_, err3 := m.DB.Collection("barbershops").Indexes().DropOne(ctx, "location_2dsphere")
	return errors.Join(err1, err2, err3)
}
