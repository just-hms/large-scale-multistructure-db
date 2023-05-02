package repo_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/config"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
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

	cfg, err := config.NewConfig()
	s.Require().NoError(err)

	fmt.Println(">>> From SetupSuite")

	mongo, err := mongo.New(cfg.Mongo.Host, cfg.Mongo.Port, "repotest")
	s.Require().NoError(err)

	redis, err := redis.New(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password)
	s.Require().NoError(err)

	s.db = mongo
	s.cache = redis
	s.resetDB = func() {
		err := redis.Clear()
		s.Require().NoError(err)

		err = mongo.DB.Drop(context.Background())
		s.Require().NoError(err)

		err = repo.AddIndexes(mongo)
		s.Require().NoError(err)

	}
}

func (s *RepoSuite) SetupTest() {
	s.resetDB()
}

func (s *RepoSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}
