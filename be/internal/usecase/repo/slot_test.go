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

func (s *RepoSuite) TestSlotBook() {
	slotRepo := repo.NewSlotRepo(s.cache)
	app := &entity.Appointment{
		BarbershopID: mockShop1ID,
		UserID:       mockUser1ID,
		StartDate:    mockDate1,
	}
	slot, err := slotRepo.Get(context.Background(), app.BarbershopID, app.StartDate)
	s.Require().NoError(err)
	s.Require().Equal(slot.Employees, -1)
	slot.Employees = 3
	err = slotRepo.Book(context.Background(), app, slot)
	s.Require().NoError(err)

	slot, err = slotRepo.Get(context.Background(), app.BarbershopID, app.StartDate)
	s.Require().NoError(err)
	s.Require().Equal(slot.Employees, 3)
	s.Require().Equal(slot.BookedAppointments, 1)

	slot.Employees = 0
	err = slotRepo.Book(context.Background(), app, slot)
	s.Require().Error(err)
}

func (s *RepoSuite) TestSlotGetByBarberShopID() {
	slotRepo := repo.NewSlotRepo(s.cache)
	appointments := []*entity.Appointment{
		{BarbershopID: mockShop1ID, UserID: mockUser1ID, StartDate: mockDate1},
		{BarbershopID: mockShop1ID, UserID: mockUser2ID, StartDate: mockDate1},
		{BarbershopID: mockShop1ID, UserID: mockUser3ID, StartDate: mockDate2},
	}

	for _, app := range appointments {
		slot, err := slotRepo.Get(context.Background(), app.BarbershopID, app.StartDate)
		s.Require().NoError(err)
		if slot.Employees == -1 {
			slot.Employees = 3
		}
		err = slotRepo.Book(context.Background(), app, slot)
		s.Require().NoError(err)
	}

	_, slots, err := slotRepo.GetByBarberShopID(context.Background(), mockShop1ID)
	s.Require().NoError(err)
	s.Require().Len(slots, 2)

	s.Require().Equal(2, slots[0].BookedAppointments)
	s.Require().Equal(1, slots[1].BookedAppointments)
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

	slot, err := slotRepo.Get(context.Background(), app.BarbershopID, app.StartDate)
	s.Require().NoError(err)
	s.Require().Equal(slot.Employees, -1)
	slot.Employees = 3
	err = slotRepo.Book(context.Background(), app, slot)
	s.Require().NoError(err)

	err = slotRepo.Cancel(context.Background(), app)
	s.Require().NoError(err)

	keys, slots, err := slotRepo.GetByBarberShopID(context.Background(), mockShop1ID)
	s.Require().NoError(err)
	s.Require().Equal(len(slots), len(keys))
	s.Require().Len(slots, 1)
	s.Require().Equal(slots[0].BookedAppointments, 0)

}

func (s *RepoSuite) TestSlotSetHoliday() {
	slotRepo := repo.NewSlotRepo(s.cache)

	app := &entity.Appointment{
		BarbershopID: mockShop1ID,
		UserID:       mockUser1ID,
		StartDate:    mockDate1,
	}

	slot, err := slotRepo.Get(context.Background(), app.BarbershopID, app.StartDate)
	s.Require().NoError(err)
	s.Require().Equal(slot.Employees, -1)
	slot.Employees = 3
	err = slotRepo.Book(context.Background(), app, slot)
	s.Require().NoError(err)

	err = slotRepo.SetEmployees(context.Background(), mockShop1ID, 5)
	s.Require().NoError(err)

	_, slots, err := slotRepo.GetByBarberShopID(context.Background(), mockShop1ID)
	s.Require().NoError(err)
	s.Require().Len(slots, 1)
	s.Require().Equal(slots[0].Employees, 5)
}
