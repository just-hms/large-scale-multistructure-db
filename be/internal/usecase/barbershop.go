package usecase

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

type BarberShopUseCase struct {
	shopRepo      BarberShopRepo
	viewRepo      ShopViewRepo
	analyticsRepo BarberAnalyticsRepo
	cache         SlotRepo
}

func NewBarberShopUseCase(shopRepo BarberShopRepo, viewRepo ShopViewRepo, analyticsRepo BarberAnalyticsRepo, cache SlotRepo) *BarberShopUseCase {
	return &BarberShopUseCase{
		shopRepo:      shopRepo,
		viewRepo:      viewRepo,
		analyticsRepo: analyticsRepo,
		cache:         cache,
	}
}

func (uc *BarberShopUseCase) Find(ctx context.Context, lat float64, lon float64, name string, radius float64) ([]*entity.BarberShop, error) {
	return uc.shopRepo.Find(ctx, lat, lon, name, radius)
}

func (uc *BarberShopUseCase) GetByID(ctx context.Context, viewerID string, ID string) (*entity.BarberShop, error) {

	// return the shop
	shop, err := uc.shopRepo.GetByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	// save the view
	err = uc.viewRepo.Store(ctx, &entity.ShopView{
		UserID:       viewerID,
		BarbershopID: ID,
	})

	if err != nil {
		return nil, err
	}

	return shop, nil

}

func (uc *BarberShopUseCase) GetOwnedShops(ctx context.Context, user *entity.User) ([]*entity.BarberShop, error) {
	return uc.shopRepo.GetOwnedShops(ctx, user)
}

func (uc *BarberShopUseCase) Store(ctx context.Context, shop *entity.BarberShop) error {
	return uc.shopRepo.Store(ctx, shop)
}

func (uc *BarberShopUseCase) ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error {

	//Update the Shop info
	err := uc.shopRepo.ModifyByID(ctx, ID, shop)
	if err != nil {
		return err
	}

	//If the Employees number got updated, check if any existing Redis Slot needs to be updated
	if shop.Employees > -1 {
		err = uc.cache.SetEmployees(ctx, ID, shop.Employees)
		if err != nil {
			return err
		}
	}

	return err

}

func (uc *BarberShopUseCase) DeleteByID(ctx context.Context, ID string) error {
	return uc.shopRepo.DeleteByID(ctx, ID)
}

func (uc *BarberShopUseCase) GetShopAnalytics(ctx context.Context, ID string) (*entity.BarberAnalytics, error) {

	var err error
	analytics := &entity.BarberAnalytics{}

	analytics.AppointmentsByMonth, err = uc.analyticsRepo.GetAppointmentCountByShop(ctx, ID)
	if err != nil {
		return nil, err
	}
	analytics.ViewsByMonth, err = uc.analyticsRepo.GetViewCountByShop(ctx, ID)
	if err != nil {
		return nil, err
	}
	analytics.ReviewsByMonth, err = uc.analyticsRepo.GetReviewCountByShop(ctx, ID)
	if err != nil {
		return nil, err
	}
	analytics.AppointmentCancellationRatioByMonth, err = uc.analyticsRepo.GetAppointmentCancellationRatioByShop(ctx, ID)
	if err != nil {
		return nil, err
	}
	analytics.AppointmentViewRatioByMonth, err = uc.analyticsRepo.GetAppointmentViewRatioByShop(ctx, ID)
	if err != nil {
		return nil, err
	}
	analytics.UpDownVoteCountByMonth, err = uc.analyticsRepo.GetUpDownVoteCountByShop(ctx, ID)
	if err != nil {
		return nil, err
	}
	analytics.ReviewWeightedRating, err = uc.analyticsRepo.GetReviewWeightedRatingByShop(ctx, ID)
	if err != nil {
		return nil, err
	}
	analytics.InactiveUsersList, err = uc.analyticsRepo.GetInactiveUsersByShop(ctx, ID)
	if err != nil {
		return nil, err
	}

	return analytics, err

}
