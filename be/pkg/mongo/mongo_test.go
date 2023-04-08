package mongo_test

import (
	"context"
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
)

func TestMongoSetup(t *testing.T) {

	mongo, err := mongo.New(&mongo.Options{
		DBName: "test",
	})

	if err != nil {
		t.Errorf("Failed to connect to Mongo: %v", err)
		return
	}

	if err := mongo.DB.Client().Ping(context.Background(), nil); err != nil {
		t.Errorf("Failed to connect to Mongo: %v", err)
		return
	}
}
