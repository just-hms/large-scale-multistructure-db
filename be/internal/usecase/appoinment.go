package usecase

import (
	"context"
	"errors"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

type AppointmentUseCase struct {
	appointmentRepo AppointmentRepo
	userRepo        UserRepo
	cache           SlotRepo
}

// New -.
func NewAppoinmentUseCase(r AppointmentRepo, c SlotRepo, u UserRepo) *AppointmentUseCase {
	return &AppointmentUseCase{
		appointmentRepo: r,
		cache:           c,
		userRepo:        u,
	}
}

func (uc *AppointmentUseCase) Book(ctx context.Context, appointment *entity.Appointment) error {

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

// TODO:
// - get the appointment from the userID?
// - get the appointment from the ID?
// - add the logit to retrieve information if not all data is provided

func (uc *AppointmentUseCase) Cancel(ctx context.Context, appointment *entity.Appointment) error {
	if appointment.UserID == "" {
		// TODO add something like this
		appointment.UserID = ""
	}
	err := uc.appointmentRepo.Cancel(ctx, appointment)
	if err != nil {
		return err
	}

	return uc.cache.Cancel(ctx, appointment)
}
