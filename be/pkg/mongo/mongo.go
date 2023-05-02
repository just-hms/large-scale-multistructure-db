package mongo

import (
	"context"
	"fmt"
	"time"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// check this
var ErrNoDocuments = mongodriver.ErrNoDocuments

type Options struct {
	DBName string
}

type Mongo struct {
	DB *mongodriver.Database
}

// get url and options as param
// add const

func New(host string, port int, dbName string) (*Mongo, error) {

	m := &Mongo{}

	mongoAddr := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err := mongodriver.NewClient(
		options.Client().ApplyURI(mongoAddr),
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	m.DB = client.Database(dbName)

	return m, nil
}
