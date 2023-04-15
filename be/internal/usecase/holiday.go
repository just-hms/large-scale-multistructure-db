package usecase

import (
	"context"
	"time"
)

type HolidayUseCase struct {
	cache SlotRepo
}

// New -.
func NewHolidayUseCase(c SlotRepo) *HolidayUseCase {
	return &HolidayUseCase{
		cache: c,
	}
}

// TODO
// - add check for not enough workes in that day
// - unavailableEmployees cannot be higher then the actual employes
// - check that the date is not in the past
func (uc *HolidayUseCase) Set(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) error {
	return uc.cache.SetHoliday(ctx, shopID, date, unavailableEmployees)
}
