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

	//Add the Username to the Appointment
	appointment.Username = us.Username

	slot, err := uc.cache.Get(ctx, appointment.BarbershopID, appointment.StartDate)
	if err != nil {
		return err
	}
	//If the slot didn't exist, fill the number of available employees
	if slot.Employees == -1 {
		shop, err := uc.shopRepo.GetByID(ctx, appointment.BarbershopID)
		if err != nil {
			return err
		}
		slot.Employees = shop.Employees
	}

	err = uc.cache.Book(ctx, appointment, slot)
	if err != nil {
		return err
	}

	err = uc.appointmentRepo.Book(ctx, appointment)
	return err
}

func (uc *AppointmentUseCase) CancelFromUser(ctx context.Context, userID string, appointment *entity.Appointment) error {
	appointment.Status = "canceled"
	err := uc.appointmentRepo.SetStatusFromUser(ctx, userID, appointment)
	if err != nil {
		return err
	}

	return uc.cache.Cancel(ctx, appointment)
}

func (uc *AppointmentUseCase) CancelFromShop(ctx context.Context, shopID string, appointment *entity.Appointment) error {
	appointment.Status = "canceled"
	err := uc.appointmentRepo.SetStatusFromShop(ctx, shopID, appointment)
	if err != nil {
		return err
	}

	return uc.cache.Cancel(ctx, appointment)
}

func (uc *AppointmentUseCase) SetCompletedFromShop(ctx context.Context, shopID string, appointment *entity.Appointment) error {
	appointment.Status = "completed"
	err := uc.appointmentRepo.SetStatusFromShop(ctx, shopID, appointment)
	if err != nil {
		return err
	}

	return uc.cache.Cancel(ctx, appointment)
}

func (uc *AppointmentUseCase) GetByID(ctx context.Context, ID string) (*entity.Appointment, error) {
	return uc.appointmentRepo.GetByID(ctx, ID)
}
