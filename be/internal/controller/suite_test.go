package controller_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/config"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/app"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ControllerSuite struct {
	suite.Suite
	srv   *gin.Engine
	db    *mongo.Mongo
	cache *redis.Redis

	resetDB func()
	fixture map[byte]string
}

func TestControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	suite.Run(t, new(ControllerSuite))
}

func (s *ControllerSuite) SetupTest() {
	s.resetDB()
}

func (s *ControllerSuite) SetupSuite() {

	cfg, err := config.NewConfig()
	s.Require().NoError(err)

	fmt.Println(">>> From SetupSuite")

	mongo, err := mongo.New(cfg.Mongo.Host, cfg.Mongo.Port, "controller-test")
	s.Require().NoError(err)

	redis, err := redis.New(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password)
	s.Require().NoError(err)

	ucs, err := app.BuildRequirements(mongo, redis, cfg)
	s.Require().NoError(err)

	s.resetDB = func() {

		err := mongo.DB.Drop(context.Background())
		s.Require().NoError(err)

		err = repo.AddIndexes(mongo)
		s.Require().NoError(err)

		s.fixture, err = InitFixture(ucs)
		s.Require().NoError(err)
	}

	// serv the mock server and db
	s.srv = controller.Router(ucs, true)
	s.db = mongo
	s.cache = redis
}

func (s *ControllerSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

const (
	USER1_ID byte = iota
	USER1_EMAIL
	USER2_ID
	USER3_ID
	ADMIN_ID
	BARBER1_ID
	BARBER2_ID

	USER1_TOKEN
	USER2_TOKEN
	USER3_TOKEN
	ADMIN_TOKEN

	BARBER1_TOKEN
	BARBER2_TOKEN

	SHOP1_ID
	SHOP2_ID
	EMPTY_SHOP

	USER1_SHOP1_APPOINTMENT_ID
)

func InitFixture(ucs map[byte]usecase.Usecase) (map[byte]string, error) {

	fixture := map[byte]string{}

	// create barbershops
	shops := []*entity.BarberShop{
		{Name: "barberShop1", Employees: 2, Location: entity.NewLocation(1, 1)},
		{Name: "barberShop2", Employees: 2, Location: entity.NewLocation(1, 2)},
		{Name: "empty", Employees: 0, Location: entity.NewLocation(1, 2)},
	}
	barberShopUsecase := ucs[usecase.BARBER_SHOP].(usecase.BarberShop)
	for _, s := range shops {
		if err := barberShopUsecase.Store(context.Background(), s); err != nil {
			return nil, err
		}
	}
	fixture[SHOP1_ID] = shops[0].ID
	fixture[SHOP2_ID] = shops[1].ID
	fixture[EMPTY_SHOP] = shops[2].ID

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
			OwnedShops: []*entity.BarberShop{
				{Name: shops[1].Name, ID: shops[1].ID},
				{Name: shops[2].Name, ID: shops[2].ID},
			},
		},
	}
	userUsecase := ucs[usecase.USER].(usecase.User)
	for _, u := range users {
		if err := userUsecase.Store(context.Background(), u); err != nil {
			return nil, err
		}
	}
	tokenUsecase := ucs[usecase.TOKEN].(usecase.Token)

	fixture[USER1_ID] = users[0].ID
	fixture[USER1_TOKEN], _ = tokenUsecase.CreateToken(users[0].ID)
	fixture[USER1_EMAIL] = users[0].Email

	fixture[USER2_ID] = users[1].ID
	fixture[USER2_TOKEN], _ = tokenUsecase.CreateToken(users[1].ID)

	fixture[ADMIN_ID] = users[2].ID
	fixture[ADMIN_TOKEN], _ = tokenUsecase.CreateToken(users[2].ID)

	fixture[USER3_ID] = users[3].ID
	fixture[USER3_TOKEN], _ = tokenUsecase.CreateToken(users[3].ID)

	fixture[BARBER1_ID] = users[4].ID
	fixture[BARBER1_TOKEN], _ = tokenUsecase.CreateToken(users[4].ID)
	fixture[BARBER2_ID] = users[5].ID
	fixture[BARBER2_TOKEN], _ = tokenUsecase.CreateToken(users[5].ID)

	// appointments

	appointments := []*entity.Appointment{
		{
			Start:        time.Now().Add(time.Hour * 2),
			UserID:       fixture[USER1_ID],
			BarbershopID: fixture[SHOP1_ID],
		},
	}
	appointmentUsecase := ucs[usecase.APPOINTMENT].(usecase.Appointment)
	for _, a := range appointments {
		err := appointmentUsecase.Book(context.Background(), a)
		if err != nil {
			return nil, err
		}
	}

	fixture[USER1_SHOP1_APPOINTMENT_ID] = appointments[0].ID

	return fixture, nil
}
