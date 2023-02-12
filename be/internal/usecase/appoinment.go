package usecase

import (
	"context"
	"large-scale-multistructure-db/be/internal/entity"
)

type AppoinmentUseCase struct {
	repo  AppointmentRepo
	cache SlotRepo
}

// New -.
func NewAppoinmentUseCase(r AppointmentRepo, c SlotRepo) *AppoinmentUseCase {
	return &AppoinmentUseCase{
		repo:  r,
		cache: c,
	}
}

func (uc *AppoinmentUseCase) Book(ctx context.Context, appointment *entity.Appointment) error {
	err := uc.repo.Book(ctx, appointment)

	if err != nil {
		return err
	}

	err = uc.cache.Book(ctx, appointment)

	return err
}

func (uc *AppoinmentUseCase) Cancel(ctx context.Context, appointment *entity.Appointment) error {

	err := uc.repo.Cancel(ctx, appointment)

	if err != nil {
		return err
	}

	err = uc.cache.Cancel(ctx, appointment)

	return err
}
