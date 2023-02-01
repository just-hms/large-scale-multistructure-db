package mongo_test

import (
	"context"
	"large-scale-multistructure-db/be/pkg/mongo"
	"testing"
)

func TestMongoSetup(t *testing.T) {

	mongo, err := mongo.New()

	if err != nil {
		t.Errorf("Failed to connect to Mongo: %v", err)
		return
	}

	if err := mongo.DB.Client().Ping(context.TODO(), nil); err != nil {
		t.Errorf("Failed to connect to Mongo: %v", err)
		return
	}
}
