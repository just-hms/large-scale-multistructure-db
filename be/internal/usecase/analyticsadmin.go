package usecase

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

type AdminAnalyticsUseCase struct {
	analyticsRepo AdminAnalyticsRepo
}

func NewAdminAnalyticsUseCase(r AdminAnalyticsRepo) *AdminAnalyticsUseCase {
	return &AdminAnalyticsUseCase{
		analyticsRepo: r,
	}
}

func (uc *AdminAnalyticsUseCase) GetAdminAnalytics(ctx context.Context) (*entity.AdminAnalytics, error) {

	var err error
	analytics := &entity.AdminAnalytics{}

	analytics.AppointmentsByMonth, err = uc.analyticsRepo.GetAppointmentCount(ctx)
	if err != nil {
		return nil, err
	}
	analytics.ViewsByMonth, err = uc.analyticsRepo.GetViewCount(ctx)
	if err != nil {
		return nil, err
	}
	analytics.ReviewsByMonth, err = uc.analyticsRepo.GetReviewCount(ctx)
	if err != nil {
		return nil, err
	}
	analytics.NewUsersByMonth, err = uc.analyticsRepo.GetNewUsersCount(ctx)
	if err != nil {
		return nil, err
	}
	analytics.AppointmentCancellationUserRanking, err = uc.analyticsRepo.GetAppointmentCancellationUserRanking(ctx)
	if err != nil {
		return nil, err
	}
	analytics.GetAppointmentCancellationShopRanking, err = uc.analyticsRepo.GetAppointmentCancellationShopRanking(ctx)
	if err != nil {
		return nil, err
	}
	analytics.GetEngagementShopRanking, err = uc.analyticsRepo.GetEngagementShopRanking(ctx)
	if err != nil {
		return nil, err
	}

	return analytics, err

}
