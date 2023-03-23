package integration_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/auth"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
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

	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}

	redis, err := redis.New()
	if err != nil {
		fmt.Printf("redis-error: %s", err.Error())
		return
	}
	// create repos and usecases

	userRepo := repo.NewUserRepo(mongo)
	barberShopRepo := repo.NewBarberShopRepo(mongo)
	viewShopRepo := repo.NewShopViewRepo(mongo)
	appintmentRepo := repo.NewAppointmentRepo(mongo)
	slotRepo := repo.NewSlotRepo(redis)

	password := auth.NewPasswordAuth()

	userUseCase := usecase.NewUserUseCase(
		userRepo,
		password,
	)

	barberUseCase := usecase.NewBarberShopUseCase(
		barberShopRepo,
		viewShopRepo,
	)

	appointmentUseCase := usecase.NewAppoinmentUseCase(
		appintmentRepo,
		slotRepo,
	)
	calendarUseCase := usecase.NewCalendarUseCase(
		slotRepo,
	)

	// Fill the test DB

	s.resetDB = func() {

		mongo.DB.Drop(context.TODO())

		p, _ := password.HashAndSalt("password")

		// create users
		us := &entity.User{
			Email:    "correct@example.com",
			Password: p,
			Type:     entity.USER,
		}

		userRepo.Store(context.TODO(), us)

		s.params["authID"] = us.ID
		s.params["authToken"], _ = jwt.CreateToken(us.ID)

		us2 := &entity.User{
			Email:    "another@example.com",
			Password: p,
			Type:     entity.USER,
		}

		userRepo.Store(context.TODO(), us2)

		s.params["auth2ID"] = us2.ID
		s.params["auth2Token"], _ = jwt.CreateToken(us2.ID)

		admin := &entity.User{
			Email:    "admin@example.com",
			Password: p,
			Type:     entity.ADMIN,
		}

		userRepo.Store(context.TODO(), admin)

		s.params["adminID"] = admin.ID
		s.params["adminToken"], _ = jwt.CreateToken(admin.ID)

		userRepo.Store(context.TODO(), &entity.User{
			Email:    "to.filter@example.com",
			Password: p,
			Type:     entity.USER,
		})

		// create barberShops

		barberShop1 := &entity.BarberShop{
			Name:      "barberShop1",
			Latitude:  "1",
			Longitude: "1",
			Employees: 2,
		}
		barberShopRepo.Store(context.TODO(), barberShop1)

		barber1 := &entity.User{
			Email:    "barber1@example.com",
			Password: p,
			Type:     entity.BARBER,
			OwnedShops: []*entity.BarberShop{
				{
					Name: barberShop1.Name,
					ID:   barberShop1.ID,
				},
			},
		}

		userRepo.Store(context.TODO(), barber1)

		s.params["barberShop1ID"] = barberShop1.ID
		s.params["barber1Auth"], _ = jwt.CreateToken(barber1.ID)

		barberShop2 := &entity.BarberShop{
			Name:      "barberShop2",
			Latitude:  "1",
			Longitude: "2",
			Employees: 2,
		}

		barberShopRepo.Store(context.TODO(), barberShop2)

		barber2 := &entity.User{
			Email:    "barber2@example.com",
			Password: p,
			Type:     entity.BARBER,

			OwnedShops: []*entity.BarberShop{
				{
					Name: barberShop2.Name,
					ID:   barberShop2.ID,
				},
			},
		}

		userRepo.Store(context.TODO(), barber2)

		s.params["barberShop2ID"] = barberShop2.ID
		s.params["barber2Auth"], _ = jwt.CreateToken(barber2.ID)

		appointment := &entity.Appointment{
			Start:        time.Now().Add(time.Hour),
			UserID:       us2.ID,
			BarbershopID: barberShop2.ID,
		}

		appointmentUseCase.Book(context.TODO(), appointment)

		s.params["appointmentID"] = appointment.ID
	}

	// serv the mock server and db
	s.params = make(map[string]string)

	s.srv = controller.Router(
		userUseCase,
		barberUseCase,
		appointmentUseCase,
		calendarUseCase,
	)

	s.mongo = mongo
}

func (s *IntegrationSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}
