package repo_test

import (
	"context"
	"fmt"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

var fixture map[byte]string

const (
	USER1_ID byte = iota
	USER2_ID
	ADMIN_ID
	BARBER1_ID
	BARBER2_ID

	SHOP1_ID
	SHOP2_ID
)

func (s *RepoSuite) SetupAnalyticsTestSuite() {

	fixture = map[byte]string{}

	user1 := &entity.User{Email: "giovanni"}
	user2 := &entity.User{Email: "banana"}
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
	err = userRepo.Store(context.Background(), user2)
	s.Require().NoError(err)
	fixture[USER2_ID] = user2.ID

	err = shopRepo.Store(context.Background(), shop1)
	s.Require().NoError(err)
	fixture[SHOP1_ID] = shop1.ID
	err = shopRepo.Store(context.Background(), shop2)
	s.Require().NoError(err)
	fixture[SHOP2_ID] = shop2.ID

	app1 := &entity.Appointment{
		CreatedAt:    time.Now().AddDate(0, -4, 0),
		StartDate:    time.Now().AddDate(0, -4, 0).Add(1 * time.Hour),
		UserID:       user1.ID,
		BarbershopID: shop1.ID,
	}
	app2 := &entity.Appointment{
		StartDate:    time.Now().Add(1 * time.Hour),
		UserID:       user1.ID,
		BarbershopID: shop1.ID,
	}
	app3 := &entity.Appointment{
		CreatedAt:    time.Now().AddDate(0, -4, 0),
		StartDate:    time.Now().AddDate(0, -4, 0).Add(1 * time.Hour),
		UserID:       user2.ID,
		BarbershopID: shop1.ID,
	}
	app4 := &entity.Appointment{
		StartDate:    time.Now().Add(1 * time.Hour),
		UserID:       user2.ID,
		BarbershopID: shop2.ID,
	}
	app5 := &entity.Appointment{
		StartDate:    time.Now().Add(1 * time.Hour),
		UserID:       user1.ID,
		BarbershopID: shop2.ID,
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

func (s *RepoSuite) TestGetAppointmentCountByShop() {

	s.SetupAnalyticsTestSuite()

	year, month, _ := time.Now().Date()
	monthKey := fmt.Sprintf("%02d-%02d", year, month)

	analyticsRepo := repo.NewBarberAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetAppointmentCountByShop(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 1)
	analytics, err = analyticsRepo.GetAppointmentCountByShop(context.Background(), fixture[SHOP2_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 2)

}

func (s *RepoSuite) TestGetViewCountByShop() {

	s.SetupAnalyticsTestSuite()

	year, month, _ := time.Now().Date()
	monthKey := fmt.Sprintf("%02d-%02d", year, month)

	analyticsRepo := repo.NewBarberAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetViewCountByShop(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 4)

	analytics, err = analyticsRepo.GetViewCountByShop(context.Background(), fixture[SHOP2_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 4)

}

func (s *RepoSuite) TestGetReviewCountByShop() {

	s.SetupAnalyticsTestSuite()

	year, month, _ := time.Now().Date()
	monthKey := fmt.Sprintf("%02d-%02d", year, month)

	analyticsRepo := repo.NewBarberAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetReviewCountByShop(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 1)

	analytics, err = analyticsRepo.GetReviewCountByShop(context.Background(), fixture[SHOP2_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 0)
}

func (s *RepoSuite) TestGetAppointmentViewRatioByShop() {

	s.SetupAnalyticsTestSuite()

	year, month, _ := time.Now().Date()
	monthKey := fmt.Sprintf("%02d-%02d", year, month)

	analyticsRepo := repo.NewBarberAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetAppointmentViewRatioByShop(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 0.25)

	analytics, err = analyticsRepo.GetAppointmentViewRatioByShop(context.Background(), fixture[SHOP2_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 0.5)
}

func (s *RepoSuite) TestGetUpDownVoteCountByShop() {

	s.SetupAnalyticsTestSuite()

	year, month, _ := time.Now().Date()
	monthKey := fmt.Sprintf("%02d-%02d", year, month)

	analyticsRepo := repo.NewBarberAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetUpDownVoteCountByShop(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey]["upCount"], 2)
	s.Require().Equal(analytics[monthKey]["downCount"], 0)

	analytics, err = analyticsRepo.GetUpDownVoteCountByShop(context.Background(), fixture[SHOP2_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey]["upCount"], 0)
	s.Require().Equal(analytics[monthKey]["downCount"], 0)

}

func (s *RepoSuite) TestGetReviewWeightedRatingByShop() {

	s.SetupAnalyticsTestSuite()

	analyticsRepo := repo.NewBarberAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetReviewWeightedRatingByShop(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics, 3.53)

	analytics, err = analyticsRepo.GetReviewWeightedRatingByShop(context.Background(), fixture[SHOP2_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics, 2.0)

}

func (s *RepoSuite) TestGetInactiveUsersByShop() {

	s.SetupAnalyticsTestSuite()

	analyticsRepo := repo.NewBarberAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetInactiveUsersByShop(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(len(analytics), 1)
	s.Require().Equal(analytics[0], fixture[USER2_ID])

	analytics, err = analyticsRepo.GetInactiveUsersByShop(context.Background(), fixture[SHOP2_ID])
	s.Require().NoError(err)
	s.Require().Equal(len(analytics), 0)

}
