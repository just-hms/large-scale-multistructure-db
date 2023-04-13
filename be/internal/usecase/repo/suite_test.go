package repo_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/redis"
	"github.com/stretchr/testify/suite"
)

type RepoSuite struct {
	suite.Suite
	db      *mongo.Mongo
	cache   *redis.Redis
	resetDB func()
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(RepoSuite))
}

func (s *RepoSuite) SetupSuite() {

	fmt.Println(">>> From SetupSuite")

	mongo, err := mongo.New(&mongo.Options{DBName: "repotest"})
	s.Require().NoError(err)
	redis, err := redis.New()
	s.Require().NoError(err)

	s.db = mongo
	s.cache = redis
	s.resetDB = func() {
		mongo.DB.Drop(context.Background())
	}
}

func (s *RepoSuite) SetupTest() {
	s.resetDB()
}

func (s *RepoSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}
