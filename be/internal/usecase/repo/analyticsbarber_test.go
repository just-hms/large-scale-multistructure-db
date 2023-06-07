package repo_test

import (
	"context"
	"fmt"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

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

func (s *RepoSuite) TestGetAppointmentCancellationRatioByShop() {

	s.SetupAnalyticsTestSuite()

	year, month, _ := time.Now().Date()
	monthKey := fmt.Sprintf("%02d-%02d", year, month)
	oldYear, oldMonth, _ := time.Now().AddDate(0, -4, 0).Date()
	oldMonthKey := fmt.Sprintf("%02d-%02d", oldYear, oldMonth)

	analyticsRepo := repo.NewBarberAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetAppointmentCancellationRatioByShop(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 0.0)
	s.Require().Equal(analytics[oldMonthKey], 0.33)

	analytics, err = analyticsRepo.GetAppointmentCancellationRatioByShop(context.Background(), fixture[SHOP2_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 0.0)
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
	s.Require().Equal(analytics[0], fixture[USER2_USERNAME])

	analytics, err = analyticsRepo.GetInactiveUsersByShop(context.Background(), fixture[SHOP2_ID])
	s.Require().NoError(err)
	s.Require().Equal(len(analytics), 0)

}
