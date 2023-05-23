package usecase

import (
	"context"
	"errors"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

type AppointmentUseCase struct {
	appointmentRepo AppointmentRepo
	userRepo        UserRepo
	shopRepo        BarberShopRepo
	cache           SlotRepo
}

// New -.
func NewAppoinmentUseCase(r AppointmentRepo, c SlotRepo, u UserRepo, s BarberShopRepo) *AppointmentUseCase {
	return &AppointmentUseCase{
		appointmentRepo: r,
		cache:           c,
		userRepo:        u,
		shopRepo:        s,
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

	shop, err := uc.shopRepo.GetByID(ctx, appointment.BarbershopID)
	if err != nil {
		return err
	}

	minus := 0
	slot, err := uc.cache.Get(ctx, appointment.BarbershopID, appointment.StartDate)
	if err == nil {
		minus = slot.BookedAppointments + slot.UnavailableEmployees
	}

	if shop.Employees-minus <= 0 {
		return errors.New("cannot book because this slot is full")
	}

	err = uc.cache.Book(ctx, appointment)
	if err != nil {
		return err
	}

	err = uc.appointmentRepo.Book(ctx, appointment)
	return err
}

func (uc *AppointmentUseCase) Cancel(ctx context.Context, appointment *entity.Appointment) error {
	err := uc.appointmentRepo.Cancel(ctx, appointment)
	if err != nil {
		return err
	}

	return uc.cache.Cancel(ctx, appointment)
}

func (uc *AppointmentUseCase) GetByIDs(ctx context.Context, shopID, ID string) (*entity.Appointment, error) {
	return uc.appointmentRepo.GetByIDs(ctx, shopID, ID)
}
