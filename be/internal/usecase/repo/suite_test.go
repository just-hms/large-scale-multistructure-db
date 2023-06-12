package repo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/config"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
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

var fixture map[byte]string

const (
	USER1_ID byte = iota
	USER2_ID
	ADMIN_ID
	BARBER1_ID
	BARBER2_ID

	USER1_USERNAME
	USER2_USERNAME

	SHOP1_ID
	SHOP2_ID

	SHOP1_NAME
	SHOP2_NAME
)

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

func (s *RepoSuite) SetupAnalyticsTestSuite() {

	fixture = map[byte]string{}

	user1 := &entity.User{
		Email:      "giovanni",
		Username:   "bigG",
		SignupDate: time.Now().AddDate(-1, -4, 0),
	}
	user2 := &entity.User{
		Email:      "banana",
		Username:   "MangoLoco",
		SignupDate: time.Now().AddDate(-1, -6, 0),
	}
	shop1 := &entity.BarberShop{Name: "brownies"}
	shop2 := &entity.BarberShop{Name: "choco"}

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)
	appointmentRepo := repo.NewAppointmentRepo(s.db)
	viewRepo := repo.NewShopViewRepo(s.db)
	reviewRepo := repo.NewReviewRepo(s.db)
	voteRepo := repo.NewVoteRepo(s.db)

	err := userRepo.Store(context.Background(), user1)
	s.Require().NoError(err)
	fixture[USER1_ID] = user1.ID
	fixture[USER1_USERNAME] = user1.Username
	err = userRepo.Store(context.Background(), user2)
	s.Require().NoError(err)
	fixture[USER2_ID] = user2.ID
	fixture[USER2_USERNAME] = user2.Username

	err = shopRepo.Store(context.Background(), shop1)
	s.Require().NoError(err)
	fixture[SHOP1_ID] = shop1.ID
	fixture[SHOP1_NAME] = shop1.Name
	err = shopRepo.Store(context.Background(), shop2)
	s.Require().NoError(err)
	fixture[SHOP2_ID] = shop2.ID
	fixture[SHOP2_NAME] = shop2.Name

	app1 := &entity.Appointment{
		CreatedAt:      time.Now().AddDate(0, -4, 0),
		StartDate:      time.Now().AddDate(0, -4, 0).Add(1 * time.Hour),
		UserID:         user1.ID,
		Username:       user1.Username,
		BarbershopID:   shop1.ID,
		BarbershopName: shop1.Name,
	}
	app2 := &entity.Appointment{
		StartDate:      time.Now().Add(1 * time.Hour),
		UserID:         user1.ID,
		Username:       user1.Username,
		BarbershopID:   shop1.ID,
		BarbershopName: shop1.Name,
	}
	app3 := &entity.Appointment{
		CreatedAt:      time.Now().AddDate(0, -4, 0),
		StartDate:      time.Now().AddDate(0, -4, 0).Add(1 * time.Hour),
		UserID:         user2.ID,
		Username:       user2.Username,
		BarbershopID:   shop1.ID,
		BarbershopName: shop1.Name,
	}
	app4 := &entity.Appointment{
		StartDate:      time.Now().Add(1 * time.Hour),
		UserID:         user2.ID,
		Username:       user2.Username,
		BarbershopID:   shop2.ID,
		BarbershopName: shop2.Name,
	}
	app5 := &entity.Appointment{
		StartDate:      time.Now().Add(1 * time.Hour),
		UserID:         user1.ID,
		Username:       user1.Username,
		BarbershopID:   shop2.ID,
		BarbershopName: shop2.Name,
	}
	app6 := &entity.Appointment{
		CreatedAt:      time.Now().AddDate(0, -4, 0),
		StartDate:      time.Now().AddDate(0, -4, 0).Add(1 * time.Hour),
		UserID:         user1.ID,
		Username:       user1.Username,
		BarbershopID:   shop1.ID,
		BarbershopName: shop1.Name,
		Status:         "canceled",
	}
	view1 := &entity.ShopView{
		UserID:       user1.ID,
		BarbershopID: shop1.ID,
	}
	view2 := &entity.ShopView{
		UserID:       user2.ID,
		BarbershopID: shop1.ID,
	}
	view3 := &entity.ShopView{
		UserID:       user1.ID,
		BarbershopID: shop2.ID,
	}
	review1 := &entity.Review{
		Rating:  4,
		Content: "test1",
		UserID:  user1.ID,
	}
	review2 := &entity.Review{
		Rating:  2,
		Content: "test2",
		UserID:  user2.ID,
	}

	err = appointmentRepo.Book(context.Background(), app1)
	s.Require().NoError(err)
	err = appointmentRepo.Book(context.Background(), app2)
	s.Require().NoError(err)
	err = appointmentRepo.Book(context.Background(), app3)
	s.Require().NoError(err)
	err = appointmentRepo.Book(context.Background(), app4)
	s.Require().NoError(err)
	err = appointmentRepo.Book(context.Background(), app5)
	s.Require().NoError(err)
	err = appointmentRepo.Book(context.Background(), app6)
	s.Require().NoError(err)

	err = viewRepo.Store(context.Background(), view1)
	s.Require().NoError(err)
	err = viewRepo.Store(context.Background(), view1)
	s.Require().NoError(err)
	err = viewRepo.Store(context.Background(), view1)
	s.Require().NoError(err)
	err = viewRepo.Store(context.Background(), view2)
	s.Require().NoError(err)
	err = viewRepo.Store(context.Background(), view3)
	s.Require().NoError(err)
	err = viewRepo.Store(context.Background(), view3)
	s.Require().NoError(err)
	err = viewRepo.Store(context.Background(), view3)
	s.Require().NoError(err)
	err = viewRepo.Store(context.Background(), view3)
	s.Require().NoError(err)

	err = reviewRepo.Store(context.Background(), review1, shop1.ID)
	s.Require().NoError(err)
	review2.CreatedAt = time.Now().AddDate(0, -2, 0)
	err = reviewRepo.Store(context.Background(), review2, shop1.ID)
	s.Require().NoError(err)
	review2.CreatedAt = time.Now().AddDate(-1, -2, 0)
	err = reviewRepo.Store(context.Background(), review2, shop1.ID)
	s.Require().NoError(err)
	review2.CreatedAt = time.Now().AddDate(0, -2, 0)
	err = reviewRepo.Store(context.Background(), review2, shop2.ID)
	s.Require().NoError(err)
	review2.CreatedAt = time.Now().AddDate(-1, -2, 0)
	err = reviewRepo.Store(context.Background(), review2, shop2.ID)
	s.Require().NoError(err)

	err = voteRepo.UpVoteByID(context.Background(), user1.ID, shop1.ID, review1.ID)
	s.Require().NoError(err)
	err = voteRepo.UpVoteByID(context.Background(), user2.ID, shop1.ID, review1.ID)
	s.Require().NoError(err)
	err = voteRepo.DownVoteByID(context.Background(), user1.ID, shop2.ID, review2.ID)
	s.Require().NoError(err)
	err = voteRepo.UpVoteByID(context.Background(), user2.ID, shop2.ID, review2.ID)
	s.Require().NoError(err)

}
