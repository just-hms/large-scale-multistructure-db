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

func (uc *HolidayUseCase) Set(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) error {
	return uc.cache.SetHoliday(ctx, shopID, date, unavailableEmployees)
}
