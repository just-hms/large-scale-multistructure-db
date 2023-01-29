package db_test

import (
	"context"
	"testing"
	"time"

	"large-scale-multistructure-db/be/infrastructures/db"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoSetup(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.URI))
	if err != nil {
		t.Errorf("Failed to connect to MongoDB: %v", err)
		return
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		t.Errorf("Failed to connect to MongoDB: %v", err)
		return
	}
}
