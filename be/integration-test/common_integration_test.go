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
	mongo *mongo.Mongo

	resetDB func()
	params  map[string]string
}

func TestIntegrationSuite(t *testing.T) {
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

		mongo.DB.Drop(context.TODO())

		// create barbershops
		shops := []*entity.BarberShop{
			{Name: "barberShop1", Employees: 2, Latitude: "1", Longitude: "1"},
			{Name: "barberShop2", Employees: 2, Latitude: "1", Longitude: "2"},
		}
		barberShopUsecase := ucs[usecase.BARBER_SHOP].(*usecase.BarberShopUseCase)
		for _, s := range shops {
			barberShopUsecase.Store(context.TODO(), s)
		}

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
			userUsecase.Store(context.TODO(), u)
		}

		s.params["authID"] = users[0].ID
		s.params["authToken"], _ = jwt.CreateToken(users[0].ID)

		s.params["auth2ID"] = users[1].ID
		s.params["auth2Token"], _ = jwt.CreateToken(users[1].ID)

		s.params["adminID"] = users[2].ID
		s.params["adminToken"], _ = jwt.CreateToken(users[2].ID)

		s.params["barberShop1ID"] = shops[0].ID
		s.params["barber1Auth"], _ = jwt.CreateToken(users[3].ID)

		s.params["barberShop2ID"] = shops[1].ID
		s.params["barber2Auth"], _ = jwt.CreateToken(users[4].ID)

		// appointments

		appointments := []*entity.Appointment{
			{Start: time.Now().Add(time.Hour), UserID: users[1].ID, BarbershopID: shops[1].ID},
		}
		appointmentUsecase := ucs[usecase.APPOINTMENT].(*usecase.AppoinmentUseCase)
		for _, a := range appointments {
			appointmentUsecase.Book(context.TODO(), a)
		}

		s.params["appointmentID"] = appointments[0].ID
	}

	// serv the mock server and db
	s.params = make(map[string]string)
	s.srv = controller.Router(ucs, true)
	s.mongo = mongo
}

func (s *IntegrationSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}
