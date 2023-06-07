package repo_test

import (
	"context"
	"fmt"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

func (s *RepoSuite) TestGetAppointmentCount() {

	s.SetupAnalyticsTestSuite()

	year, month, _ := time.Now().Date()
	monthKey := fmt.Sprintf("%02d-%02d", year, month)
	oldYear, oldMonth, _ := time.Now().AddDate(0, -4, 0).Date()
	oldMonthKey := fmt.Sprintf("%02d-%02d", oldYear, oldMonth)

	analyticsRepo := repo.NewAdminAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetAppointmentCount(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 3)
	s.Require().Equal(analytics[oldMonthKey], 2)
}

func (s *RepoSuite) TestGetViewCount() {

	s.SetupAnalyticsTestSuite()

	year, month, _ := time.Now().Date()
	monthKey := fmt.Sprintf("%02d-%02d", year, month)

	analyticsRepo := repo.NewAdminAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetViewCount(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 8)

}

func (s *RepoSuite) TestGetReviewCount() {

	s.SetupAnalyticsTestSuite()

	year, month, _ := time.Now().Date()
	monthKey := fmt.Sprintf("%02d-%02d", year, month)
	oldYear, oldMonth, _ := time.Now().AddDate(0, -2, 0).Date()
	oldMonthKey := fmt.Sprintf("%02d-%02d", oldYear, oldMonth)
	olderYear, olderMonth, _ := time.Now().AddDate(-1, -2, 0).Date()
	olderMonthKey := fmt.Sprintf("%02d-%02d", olderYear, olderMonth)

	analyticsRepo := repo.NewAdminAnalyticsRepo(s.db)

	analytics, err := analyticsRepo.GetReviewCount(context.Background(), fixture[SHOP1_ID])
	s.Require().NoError(err)
	s.Require().Equal(analytics[monthKey], 1)
	s.Require().Equal(analytics[oldMonthKey], 2)
	s.Require().Equal(analytics[olderMonthKey], 2)
}
