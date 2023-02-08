package usecase

import "large-scale-multistructure-db/be/internal/entity"

type BarberShopUseCase struct {
	repo BarberShopRepo
}

func NewBarberShopUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *BarberShopUseCase) Find(lat string, lon string, name string, radius string) ([]*entity.BarberShop, error) {
	return uc.repo.Find(lat, lon, name, radius)
}
