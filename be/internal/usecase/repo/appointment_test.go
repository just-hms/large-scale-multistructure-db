package repo_test

import (
	"context"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

func (s *RepoSuite) TestAppointmentBook() {

	user := &entity.User{Email: "giovanni"}
	user2 := &entity.User{Email: "banana"}
	shop := &entity.BarberShop{Name: "brownies"}

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)

	err := userRepo.Store(context.Background(), user)
	s.Require().NoError(err)
	err = userRepo.Store(context.Background(), user2)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop)
	s.Require().NoError(err)

	testCases := []struct {
		name        string
		input       *entity.Appointment
		expectedErr bool
	}{
		{
			name:        "Correctly inserted",
			expectedErr: false,
			input: &entity.Appointment{
				StartDate:    time.Now().Add(1 * time.Hour),
				UserID:       user.ID,
				BarbershopID: shop.ID,
			},
		},
		{
			name:        "The shop does not exists",
			expectedErr: true,
			input: &entity.Appointment{
				StartDate:    time.Now().Add(1 * time.Hour),
				UserID:       user2.ID,
				BarbershopID: "fake_id",
			},
		},
	}

	appointmentRepo := repo.NewAppointmentRepo(s.db)
	for _, tc := range testCases {

		s.Run(tc.name, func() {
			err = appointmentRepo.Book(context.Background(), tc.input)
			if tc.expectedErr {
				s.Require().Error(err)
				return
			}

			s.Require().NoError(err)
			// check that the appointment was correctly created

			user, err = userRepo.GetByID(context.Background(), user.ID)
			s.Require().NoError(err)
			// in the user collection
			s.Require().NotNil(user.CurrentAppointment)
			// in the barbershop collection
			shop, err = shopRepo.GetByID(context.Background(), shop.ID)
			s.Require().NoError(err)
			s.Require().Len(shop.Appointments, 1)
		})
	}

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

	appointment.Status = "canceled"
	err = appointmentRepo.SetStatusFromUser(context.Background(), user.ID, appointment)
	s.Require().NoError(err)

	// check that the appointment was correctly cancelled

	// in the user collection
	user, err = userRepo.GetByID(context.Background(), user.ID)
	s.Require().NoError(err)
	s.Require().Nil(user.CurrentAppointment)

	// in the barbershop collection
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Appointments, 1)
	s.Require().Equal(shop.Appointments[0].Status, "canceled")

	// Test that Cancel also works from the Shop

	err = appointmentRepo.Book(context.Background(), appointment)
	s.Require().NoError(err)

	appointment.Status = "canceled"
	err = appointmentRepo.SetStatusFromShop(context.Background(), shop.ID, appointment)
	s.Require().NoError(err)

	// check that the appointment was correctly cancelled

	// in the user collection
	user, err = userRepo.GetByID(context.Background(), user.ID)
	s.Require().NoError(err)
	s.Require().Nil(user.CurrentAppointment)

	// in the barbershop collection
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Appointments, 2)
	s.Require().Equal(shop.Appointments[1].Status, "canceled")
}

func (s *RepoSuite) TestAppointmentComplete() {

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

	appointment.Status = "completed"
	err = appointmentRepo.SetStatusFromShop(context.Background(), shop.ID, appointment)
	s.Require().NoError(err)

	// check that the appointment was correctly cancelled

	// in the user collection
	user, err = userRepo.GetByID(context.Background(), user.ID)
	s.Require().NoError(err)
	s.Require().Nil(user.CurrentAppointment)

	// in the barbershop collection
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Appointments, 1)
	s.Require().Equal(shop.Appointments[0].Status, "completed")
}
