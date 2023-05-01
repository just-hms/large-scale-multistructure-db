package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/env"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
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

func New(opt *Options) (*Mongo, error) {

	m := &Mongo{}

	dbHost := env.GetStringWithDefault("MONGO_HOST", "localhost")
	dbPort, err := env.GetInt("MONGO_PORT")
	if err != nil {
		return nil, err
	}

	mongoAddr := fmt.Sprintf("mongodb://%s:%d", dbHost, dbPort)

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

	m.DB = client.Database(opt.DBName)

	return m, nil
}

// TODO: make this more general
func (m *Mongo) CreateIndex(ctx context.Context) error {
	indexOpts := options.CreateIndexes().SetMaxTime(time.Second * 10)

	// Index to location 2dsphere type.
	pointIndexModel := mongodriver.IndexModel{
		Keys: bsonx.MDoc{"location": bsonx.String("2dsphere")},
	}

	pointIndexes := m.DB.Collection("barbershops").Indexes()
	_, err := pointIndexes.CreateOne(ctx, pointIndexModel, indexOpts)
	return err
}
