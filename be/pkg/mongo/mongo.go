package mongo

import (
	"context"
	"time"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// check this
var ErrNoDocuments = mongodriver.ErrNoDocuments

const (
	DB_URI  = "mongodb://db:27017"
	DB_NAME = "db"
)

type Mongo struct {
	DB *mongodriver.Database
}

// get url and options as param
// add const

func New() (*Mongo, error) {

	m := &Mongo{}

	client, err := mongodriver.NewClient(options.Client().ApplyURI(DB_URI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	m.DB = client.Database(DB_NAME)

	return m, nil
}
