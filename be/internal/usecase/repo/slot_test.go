package repo_test

import (
	"context"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

var (
	mockShop1ID = "1"
	mockUser1ID = "user1"
	mockUser2ID = "user2"
	mockUser3ID = "user3"

	mockDate1 = time.Now().Add(time.Hour)
	mockDate2 = time.Now().Add(time.Hour * 2)
)

func (s *RepoSuite) TestSlotGetByBarberShopID() {
	slotRepo := repo.NewSlotRepo(s.cache)
	appointments := []*entity.Appointment{
		{BarbershopID: mockShop1ID, UserID: mockUser1ID, StartDate: mockDate1},
		{BarbershopID: mockShop1ID, UserID: mockUser2ID, StartDate: mockDate1},
		{BarbershopID: mockShop1ID, UserID: mockUser3ID, StartDate: mockDate2},
	}

	for _, app := range appointments {
		err := slotRepo.Book(context.Background(), app)
		s.Require().NoError(err)
	}

	slots, err := slotRepo.GetByBarberShopID(context.Background(), mockShop1ID)
	s.Require().NoError(err)
	s.Require().Len(slots, 2)

	s.Require().Equal(2, slots[0].BookedAppointments)
	s.Require().Equal(1, slots[1].BookedAppointments)
}

func (s *RepoSuite) TestSlotBook() {
	slotRepo := repo.NewSlotRepo(s.cache)
	app := &entity.Appointment{
		BarbershopID: mockShop1ID,
		UserID:       mockUser1ID,
		StartDate:    mockDate1,
	}
	slotRepo.Book(context.Background(), app)

	slots, err := slotRepo.GetByBarberShopID(context.Background(), mockShop1ID)
	s.Require().NoError(err)
	s.Require().Len(slots, 1)
}

func (s *RepoSuite) TestSlotCancel() {
	slotRepo := repo.NewSlotRepo(s.cache)

	app := &entity.Appointment{
		BarbershopID: mockShop1ID,
		UserID:       mockUser1ID,
		StartDate:    mockDate1,
	}

	// the cache is empty so the cancel should return an error
	err := slotRepo.Cancel(context.Background(), app)
	s.Require().Error(err)

	err = slotRepo.Book(context.Background(), app)
	s.Require().NoError(err)

	err = slotRepo.Cancel(context.Background(), app)
	s.Require().NoError(err)

	slots, err := slotRepo.GetByBarberShopID(context.Background(), mockShop1ID)
	s.Require().NoError(err)
	s.Require().Len(slots, 1)
	s.Require().Equal(slots[0].BookedAppointments, 0)

}

func (s *RepoSuite) TestSlotSetHoliday() {
	slotRepo := repo.NewSlotRepo(s.cache)

	err := slotRepo.SetHoliday(context.Background(), mockShop1ID, mockDate1, 5)
	s.Require().NoError(err)

	slots, err := slotRepo.GetByBarberShopID(context.Background(), mockShop1ID)
	s.Require().NoError(err)
	s.Require().Len(slots, 1)
	s.Require().Equal(slots[0].AvailableEmployees, 5)
}

// TODO
// - test expiration in some way ???
