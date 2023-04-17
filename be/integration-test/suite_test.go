package integration_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/app"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
	srv   *gin.Engine
	db    *mongo.Mongo
	cache *redis.Redis

	resetDB func()
	fixture map[byte]string
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) SetupTest() {
	s.resetDB()
}

func (s *IntegrationSuite) SetupSuite() {

	fmt.Println(">>> From SetupSuite")

	mongo, err := mongo.New(&mongo.Options{DBName: "test"})
	s.Require().NoError(err)
	redis, err := redis.New()
	s.Require().NoError(err)

	ucs := app.BuildRequirements(mongo, redis)

	s.resetDB = func() {

		err := mongo.DB.Drop(context.Background())
		s.Require().NoError(err)

		// TODO: move this somewhere else
		err = mongo.CreateIndex(context.Background())
		s.Require().NoError(err)

		s.fixture, err = InitFixture(ucs)
		s.Require().NoError(err)
	}

	// serv the mock server and db
	s.srv = controller.Router(ucs, true)
	s.db = mongo
	s.cache = redis
}

func (s *IntegrationSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}
