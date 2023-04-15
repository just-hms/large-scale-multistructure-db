package repo_test

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

func (s *RepoSuite) TestAppointmentBook() {

	user := &entity.User{Username: "giovanni"}
	shop := &entity.BarberShop{Name: "brownies"}

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)

	err := userRepo.Store(context.Background(), user)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop)
	s.Require().NoError(err)

	appointment := &entity.Appointment{
		UserID:       user.ID,
		BarbershopID: shop.ID,
	}

	appointmentRepo := repo.NewAppointmentRepo(s.db)

	err = appointmentRepo.Book(context.Background(), appointment)
	s.Require().NoError(err)

	// check that the appointment was correctly created

	// in the user collection
	user, err = userRepo.GetByID(context.Background(), user.ID)
	s.Require().NoError(err)
	s.Require().NotNil(user.CurrentAppointment)

	// in the barbershop collection
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Appointments, 1)

}

func (s *RepoSuite) TestAppointmentCancel() {

	user := &entity.User{Username: "giovanni"}
	shop := &entity.BarberShop{Name: "brownies"}

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)

	err := userRepo.Store(context.Background(), user)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop)
	s.Require().NoError(err)

	appointment := &entity.Appointment{
		UserID:       user.ID,
		BarbershopID: shop.ID,
	}

	appointmentRepo := repo.NewAppointmentRepo(s.db)

	err = appointmentRepo.Book(context.Background(), appointment)
	s.Require().NoError(err)

	err = appointmentRepo.Cancel(context.Background(), appointment)
	s.Require().NoError(err)

	// check that the appointment was correctly cancelled

	// in the user collection
	user, err = userRepo.GetByID(context.Background(), user.ID)
	s.Require().NoError(err)
	s.Require().Nil(user.CurrentAppointment)

	// in the barbershop collection
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Appointments, 0)

}
