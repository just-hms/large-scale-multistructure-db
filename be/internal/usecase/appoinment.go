package usecase

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
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

	// TODO: check if the hour has meaning
	// TODO: get the user repo and add the current appoinment from there
	err := uc.repo.Book(ctx, appointment)

	if err != nil {
		return err
	}

	err = uc.cache.Book(ctx, appointment)

	return err
}

func (uc *AppoinmentUseCase) Cancel(ctx context.Context, appointment *entity.Appointment) error {

	// TODO: get the user repo and remove the current appointment from there
	err := uc.repo.Cancel(ctx, appointment)

	if err != nil {
		return err
	}

	err = uc.cache.Cancel(ctx, appointment)

	return err
}
