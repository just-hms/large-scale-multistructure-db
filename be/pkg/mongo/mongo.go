package mongo

import (
	"context"
	"time"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// check this
var ErrNoDocuments = mongodriver.ErrNoDocuments

type Options struct {
	DB_URI  string
	DB_NAME string
}

type Mongo struct {
	DB  *mongodriver.Database
	sex mongodriver.Session
}

// get url and options as param
// add const

func New(opt *Options) (*Mongo, error) {

	m := &Mongo{}

	if opt.DB_URI == "" {
		opt.DB_URI = "mongodb://db:27017"
	}

	if opt.DB_NAME == "" {
		opt.DB_NAME = "test"
	}

	client, err := mongodriver.NewClient(options.Client().ApplyURI(opt.DB_URI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	m.DB = client.Database(opt.DB_NAME)

	return m, nil
}

func (m *Mongo) Record() {
	m.sex, _ = m.DB.Client().StartSession(&options.SessionOptions{
		CausalConsistency: &options.DefaultCausalConsistency,
	})
	m.sex.StartTransaction()
}

func (m *Mongo) RollBack(ctx context.Context) {
	m.sex.AbortTransaction(ctx)
	m.sex.EndSession(ctx)
}
