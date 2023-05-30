package repo_test

import (
	"context"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

func (s *RepoSuite) TestGetAppointmentViewReviewCount() {

	user1 := &entity.User{Email: "giovanni"}
	user2 := &entity.User{Email: "banana"}
	shop1 := &entity.BarberShop{Name: "brownies"}
	shop2 := &entity.BarberShop{Name: "choco"}

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)
	appointmentRepo := repo.NewAppointmentRepo(s.db)
	viewRepo := repo.NewShopViewRepo(s.db)
	reviewRepo := repo.NewReviewRepo(s.db)
	analyticsRepo := repo.NewAnalyticsRepo(s.db)

	err := userRepo.Store(context.Background(), user1)
	s.Require().NoError(err)
	err = userRepo.Store(context.Background(), user2)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop1)
	s.Require().NoError(err)
	err = shopRepo.Store(context.Background(), shop2)
	s.Require().NoError(err)

	app1 := &entity.Appointment{
		StartDate:    time.Now().Add(1 * time.Hour),
		UserID:       user1.ID,
		BarbershopID: shop1.ID,
	}
	app2 := &entity.Appointment{
		StartDate:    time.Now().Add(1 * time.Hour),
		UserID:       user2.ID,
		BarbershopID: shop1.ID,
	}
	view1 := &entity.ShopView{
		ViewerID:     user1.ID,
		BarberShopID: shop1.ID,
	}
	view2 := &entity.ShopView{
		ViewerID:     user2.ID,
		BarberShopID: shop1.ID,
	}
	view3 := &entity.ShopView{
		ViewerID:     user1.ID,
		BarberShopID: shop2.ID,
	}
	review1 := &entity.Review{
		Rating:  4,
		Content: "test1",
		UserID:  user1.ID,
	}
	review2 := &entity.Review{
		Rating:  5,
		Content: "test2",
		UserID:  user2.ID,
	}

	err = appointmentRepo.Book(context.Background(), app1)
	s.Require().NoError(err)
	err = appointmentRepo.Book(context.Background(), app2)
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

	err = reviewRepo.Store(context.Background(), review1, shop1.ID)
	s.Require().NoError(err)
	err = reviewRepo.Store(context.Background(), review2, shop2.ID)
	s.Require().NoError(err)
	err = reviewRepo.Store(context.Background(), review2, shop2.ID)
	s.Require().NoError(err)

	analytics, err := analyticsRepo.GetAppointmentViewReviewCount(context.Background(), shop1.ID)
	s.Require().NoError(err)
	s.Require().Equal(analytics["appointmentCount"], 2)
	s.Require().Equal(analytics["viewCount"], 4)
	s.Require().Equal(analytics["reviewCount"], 1)
	analytics, err = analyticsRepo.GetAppointmentViewReviewCount(context.Background(), shop2.ID)
	s.Require().NoError(err)
	s.Require().Equal(analytics["appointmentCount"], 0)
	s.Require().Equal(analytics["viewCount"], 1)
	s.Require().Equal(analytics["reviewCount"], 2)
}
