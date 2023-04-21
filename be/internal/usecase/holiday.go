package usecase

import (
	"context"
	"errors"
	"time"
)

type HolidayUseCase struct {
	cache SlotRepo
	shop  BarberShopRepo
}

// New -.
func NewHolidayUseCase(c SlotRepo, s BarberShopRepo) *HolidayUseCase {
	return &HolidayUseCase{
		cache: c,
		shop:  s,
	}
}

// TODO
// - add check for not enough workes in that day
// - unavailableEmployees cannot be higher then the actual employes
// - check that the date is not in the past
func (uc *HolidayUseCase) Set(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) error {

	if date.Before(time.Now()) {
		return errors.New("cannot set an holiday in the past")
	}

	shop, err := uc.shop.GetByID(ctx, shopID)

	if err != nil {
		return errors.New("the specified shop does not exists")
	}

	slot, err := uc.cache.Get(ctx, shopID, date)
	if err != nil {
		return err
	}

	if shop.Employees-slot.BookedAppoIntments < unavailableEmployees {
		return errors.New("cannot add more unavailableEmployees then the number of available Employees")
	}
	return uc.cache.SetHoliday(ctx, shopID, date, unavailableEmployees)
}
