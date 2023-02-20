package integration_test

import (
	"context"
	"fmt"
	"large-scale-multistructure-db/be/internal/controller"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/internal/usecase/auth"
	"large-scale-multistructure-db/be/internal/usecase/repo"
	"large-scale-multistructure-db/be/pkg/jwt"
	"large-scale-multistructure-db/be/pkg/mongo"
	"testing"

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

	mongo, err := mongo.New(&mongo.Options{
		DB_NAME: "test",
	})

	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}

	// create repos and usecases
	userRepo := repo.NewUserRepo(mongo)
	barberShopRepo := repo.NewBarberShopRepo(mongo)
	viewShopRepo := repo.NewShopViewRepo(mongo)

	password := auth.NewPasswordAuth()

	usecases := []usecase.Usecase{
		usecase.NewUserUseCase(
			userRepo,
			password,
		),
		usecase.NewBarberShopUseCase(
			barberShopRepo,
			viewShopRepo,
		),
	}

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
			Name: "barberShop1",
			Coordinates: entity.Coordinates{
				Latitude:  "1",
				Longitude: "1",
			},
			Employees: 2,
		}
		barberShopRepo.Store(context.TODO(), barberShop1)

		barber1 := &entity.User{
			Email:    "barber1@example.com",
			Password: p,
			Type:     entity.BARBER,
			OwnedShops: []*entity.BarberShop{
				barberShop1,
			},
		}

		userRepo.Store(context.TODO(), barber1)

		s.params["barberShop1ID"] = barberShop1.ID
		s.params["barber1Auth"], _ = jwt.CreateToken(barber1.ID)

		barberShop2 := &entity.BarberShop{
			Name: "barberShop2",
			Coordinates: entity.Coordinates{
				Latitude:  "1",
				Longitude: "2",
			},
			Employees: 2,
		}

		barberShopRepo.Store(context.TODO(), barberShop2)

		barber2 := &entity.User{
			Email:    "barber2@example.com",
			Password: p,
			Type:     entity.BARBER,
			OwnedShops: []*entity.BarberShop{
				barberShop2,
			},
		}

		userRepo.Store(context.TODO(), barber2)

		s.params["barberShop2ID"] = barberShop2.ID
		s.params["barber2Auth"], _ = jwt.CreateToken(barber2.ID)

	}

	// serv the mock server and db
	s.params = make(map[string]string)
	s.srv = controller.Router(usecases)
	s.mongo = mongo
}

func (s *IntegrationSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}
