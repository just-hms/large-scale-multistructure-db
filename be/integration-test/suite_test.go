package integration_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/app"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/jwt"
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
	params  []string
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

const (
	USER1_ID int = iota
	USER2_ID
	ADMIN_ID
	BARBER1_ID
	BARBER2_ID

	USER1_TOKEN
	USER2_TOKEN
	ADMIN_TOKEN

	BARBER1_TOKEN
	BARBER2_TOKEN

	SHOP1_ID
	SHOP2_ID

	USER1_SHOP1_APPOINTMENT
	PARAM_LEN
)

func (s *IntegrationSuite) SetupSuite() {

	fmt.Println(">>> From SetupSuite")

	mongo, err := mongo.New(&mongo.Options{DBName: "test"})
	s.Require().NoError(err)
	redis, err := redis.New()
	s.Require().NoError(err)

	ucs := app.BuildRequirements(mongo, redis)

	s.resetDB = func() {

		mongo.DB.Drop(context.Background())

		// create barbershops
		shops := []*entity.BarberShop{
			{Name: "barberShop1", Employees: 2, Location: entity.NewLocation(1, 1)},
			{Name: "barberShop2", Employees: 2, Location: entity.NewLocation(1, 2)},
		}
		barberShopUsecase := ucs[usecase.BARBER_SHOP].(*usecase.BarberShopUseCase)
		for _, s := range shops {
			barberShopUsecase.Store(context.Background(), s)
		}
		s.params[SHOP1_ID] = shops[0].ID
		s.params[SHOP2_ID] = shops[1].ID

		// create users
		users := []*entity.User{
			{Email: "correct@example.com", Password: "password", Type: entity.USER},
			{Email: "another@example.com", Password: "password", Type: entity.USER},
			{Email: "admin@example.com", Password: "password", Type: entity.ADMIN},
			{Email: "to.filter@example.com", Password: "password", Type: entity.USER},

			{
				Email: "barber1@example.com", Password: "password", Type: entity.BARBER,
				OwnedShops: []*entity.BarberShop{{Name: shops[0].Name, ID: shops[0].ID}},
			},
			{
				Email: "barber2@example.com", Password: "password", Type: entity.BARBER,
				OwnedShops: []*entity.BarberShop{{Name: shops[1].Name, ID: shops[1].ID}},
			},
		}
		userUsecase := ucs[usecase.USER].(*usecase.UserUseCase)
		for _, u := range users {
			userUsecase.Store(context.Background(), u)
		}

		s.params[USER1_ID] = users[0].ID
		s.params[USER1_TOKEN], _ = jwt.CreateToken(users[0].ID)

		s.params[USER2_ID] = users[1].ID
		s.params[USER2_TOKEN], _ = jwt.CreateToken(users[1].ID)

		s.params[ADMIN_ID] = users[2].ID
		s.params[ADMIN_TOKEN], _ = jwt.CreateToken(users[2].ID)

		s.params[BARBER1_ID] = users[4].ID
		s.params[BARBER1_TOKEN], _ = jwt.CreateToken(users[4].ID)
		s.params[BARBER2_ID] = users[5].ID
		s.params[BARBER2_TOKEN], _ = jwt.CreateToken(users[5].ID)

		// appointments

		appointments := []*entity.Appointment{
			{
				Start: time.Now().Add(time.Hour), UserID: s.params[USER1_ID],
				BarbershopID: s.params[SHOP1_ID],
			},
		}
		appointmentUsecase := ucs[usecase.APPOINTMENT].(*usecase.AppoinmentUseCase)
		for _, a := range appointments {
			appointmentUsecase.Book(context.Background(), a)
		}

		s.params[USER1_SHOP1_APPOINTMENT] = appointments[0].ID
	}

	// serv the mock server and db
	s.params = make([]string, PARAM_LEN)
	s.srv = controller.Router(ucs, true)
	s.db = mongo
	s.cache = redis
}

func (s *IntegrationSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}
