package usecase

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

type BarberShopUseCase struct {
	shopRepo BarberShopRepo
	viewRepo ShopViewRepo
}

func NewBarberShopUseCase(shopRepo BarberShopRepo, viewRepo ShopViewRepo) *BarberShopUseCase {
	return &BarberShopUseCase{
		shopRepo: shopRepo,
		viewRepo: viewRepo,
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
		ViewerID:     viewerID,
		BarberShopID: ID,
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
	if shop.Employees != -1 {

	}

	return err

}
func (uc *BarberShopUseCase) DeleteByID(ctx context.Context, ID string) error {
	return uc.shopRepo.DeleteByID(ctx, ID)
}
