package mongo

import (
	"context"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/osext"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// check this
var ErrNoDocuments = mongodriver.ErrNoDocuments

type Options struct {
	DBName string
}

type Mongo struct {
	DB  *mongodriver.Database
	sex mongodriver.Session
}

// get url and options as param
// add const

func New(opt *Options) (*Mongo, error) {

	m := &Mongo{}

	dbUri, err := osext.GetStringEnv("MONGO_URI")
	if err != nil {
		return nil, err
	}

	client, err := mongodriver.NewClient(options.Client().ApplyURI(dbUri))
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
