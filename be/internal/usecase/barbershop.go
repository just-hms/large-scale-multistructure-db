package usecase

import (
	"context"
	"large-scale-multistructure-db/be/internal/entity"
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

func (uc *BarberShopUseCase) Find(ctx context.Context, lat string, lon string, name string, radius string) ([]*entity.BarberShop, error) {
	return uc.shopRepo.Find(ctx, lat, lon, name, radius)
}

func (uc *BarberShopUseCase) GetByID(ctx context.Context, viewerID string, ID string) (*entity.BarberShop, error) {

	// return the shop
	shop, err := uc.shopRepo.GetByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	// save the view
	_, err = uc.viewRepo.Store(ctx, &entity.ShopView{
		ViewerID:     viewerID,
		BarberShopID: ID,
	})

	if err != nil {
		return nil, err
	}

	return shop, nil

}

func (uc *BarberShopUseCase) Store(ctx context.Context, shop *entity.BarberShop) (string, error) {
	return uc.shopRepo.Store(ctx, shop)
}
func (uc *BarberShopUseCase) ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error {
	return uc.shopRepo.ModifyByID(ctx, ID, shop)
}
func (uc *BarberShopUseCase) DeleteByID(ctx context.Context, ID string) error {
	return uc.shopRepo.DeleteByID(ctx, ID)
}
