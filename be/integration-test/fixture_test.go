package integration_test

import (
	"context"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/jwt"
)

const (
	USER1_ID byte = iota
	USER2_ID
	ADMIN_ID
	BARBER1_ID
	BARBER2_ID

	USER1_TOKEN
	USER2_TOKEN
	ADMIN_TOKEN

	BARBER1_TOKEN
	BARBER2_TOKEN

	SHOP1_ID
	SHOP2_ID

	USER1_SHOP1_APPOINTMENT_ID
)

func InitFixture(ucs map[byte]usecase.Usecase) (map[byte]string, error) {

	fixture := map[byte]string{}

	// create barbershops
	shops := []*entity.BarberShop{
		{Name: "barberShop1", Employees: 2, Location: entity.NewLocation(1, 1)},
		{Name: "barberShop2", Employees: 2, Location: entity.NewLocation(1, 2)},
	}
	barberShopUsecase := ucs[usecase.BARBER_SHOP].(usecase.BarberShop)
	for _, s := range shops {
		err := barberShopUsecase.Store(context.Background(), s)
		if err != nil {
			return nil, err
		}
	}
	fixture[SHOP1_ID] = shops[0].ID
	fixture[SHOP2_ID] = shops[1].ID

	// create users
	users := []*entity.User{
		{Email: "correct@example.com", Password: "password", Type: entity.USER},
		{Email: "another@example.com", Password: "password", Type: entity.USER},
		{Email: "admin@example.com", Password: "password", Type: entity.ADMIN},
		{Email: "to.filter@example.com", Password: "password", Type: entity.USER},

		{
			Email: "barber1@example.com", Password: "password", Type: entity.BARBER,
			OwnedShops: []*entity.BarberShop{{Name: shops[0].Name, ID: shops[0].ID}},
		},
		{
			Email: "barber2@example.com", Password: "password", Type: entity.BARBER,
			OwnedShops: []*entity.BarberShop{{Name: shops[1].Name, ID: shops[1].ID}},
		},
	}
	userUsecase := ucs[usecase.USER].(usecase.User)
	for _, u := range users {
		err := userUsecase.Store(context.Background(), u)
		if err != nil {
			return nil, err
		}
	}

	fixture[USER1_ID] = users[0].ID
	fixture[USER1_TOKEN], _ = jwt.CreateToken(users[0].ID)

	fixture[USER2_ID] = users[1].ID
	fixture[USER2_TOKEN], _ = jwt.CreateToken(users[1].ID)

	fixture[ADMIN_ID] = users[2].ID
	fixture[ADMIN_TOKEN], _ = jwt.CreateToken(users[2].ID)

	fixture[BARBER1_ID] = users[4].ID
	fixture[BARBER1_TOKEN], _ = jwt.CreateToken(users[4].ID)
	fixture[BARBER2_ID] = users[5].ID
	fixture[BARBER2_TOKEN], _ = jwt.CreateToken(users[5].ID)

	// appointments

	appointments := []*entity.Appointment{
		{
			Start: time.Now().Add(time.Hour * 2), UserID: fixture[USER1_ID],
			BarbershopID: fixture[SHOP1_ID],
		},
	}
	appointmentUsecase := ucs[usecase.APPOINTMENT].(usecase.Appointment)
	for _, a := range appointments {
		err := appointmentUsecase.Book(context.Background(), a)
		if err != nil {
			return nil, err
		}
	}

	fixture[USER1_SHOP1_APPOINTMENT_ID] = appointments[0].ID

	return fixture, nil
}
