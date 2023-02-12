package usecase

import (
	"context"
	"large-scale-multistructure-db/be/internal/entity"
)

type CalendarUseCase struct {
	repo SlotRepo
}

func NewCalendarUseCase(r SlotRepo) *CalendarUseCase {
	return &CalendarUseCase{
		repo: r,
	}
}

func (uc *CalendarUseCase) GetByBarberShopID(ctx context.Context, ID string) (*entity.Calendar, error) {

	slots, err := uc.repo.GetByBarberShopID(ctx, ID)

	if err != nil {
		return nil, err
	}

	return &entity.Calendar{
		Slots: slots,
	}, nil

}
