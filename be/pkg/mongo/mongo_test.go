package mongo_test

import (
	"context"
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/config"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"github.com/stretchr/testify/assert"
)

func TestMongoSetup(t *testing.T) {

	cfg, err := config.NewConfig()
	assert.Nil(t, err)

	mongo, err := mongo.New(cfg.Mongo.Host, cfg.Mongo.Port, "test")
	assert.Nil(t, err)

	err = mongo.DB.Client().Ping(context.Background(), nil)
	assert.Nil(t, err)
}
