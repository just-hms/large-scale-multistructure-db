package usecase

import (
	"context"
	"errors"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

type AppoinmentUseCase struct {
	appointmentRepo AppointmentRepo
	userRepo        UserRepo
	cache           SlotRepo
}

// New -.
func NewAppoinmentUseCase(r AppointmentRepo, c SlotRepo, u UserRepo) *AppoinmentUseCase {
	return &AppoinmentUseCase{
		appointmentRepo: r,
		cache:           c,
		userRepo:        u,
	}
}

func (uc *AppoinmentUseCase) Book(ctx context.Context, appointment *entity.Appointment) error {

	us, err := uc.userRepo.GetByID(ctx, appointment.UserID)
	if err != nil {
		return err
	}

	if us.CurrentAppointment != nil {
		return errors.New("cannot book two appointments")
	}

	err = uc.appointmentRepo.Book(ctx, appointment)
	if err != nil {
		return err
	}

	err = uc.cache.Book(ctx, appointment)
	return err
}

func (uc *AppoinmentUseCase) Cancel(ctx context.Context, appointment *entity.Appointment) error {
	err := uc.appointmentRepo.Cancel(ctx, appointment)
	if err != nil {
		return err
	}

	return uc.cache.Cancel(ctx, appointment)
}
