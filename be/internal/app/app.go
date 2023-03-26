package app

import (
	"context"
	"fmt"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/auth"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/redis"
)

func Run() {

	mongo, err := mongo.New(&mongo.Options{DBName: "barber-deploy"})
	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}

	redis, err := redis.New()
	if err != nil {
		fmt.Printf("redis-error: %s", err.Error())
		return
	}

	userRepo := repo.NewUserRepo(mongo)
	barberShopRepo := repo.NewBarberShopRepo(mongo)
	viewShopRepo := repo.NewShopViewRepo(mongo)
	appintmentRepo := repo.NewAppointmentRepo(mongo)
	slotRepo := repo.NewSlotRepo(redis)

	password := auth.NewPasswordAuth()

	ucs := make([]usecase.Usecase, usecase.LEN)
	ucs[usecase.USER] = usecase.NewUserUseCase(userRepo, password)
	ucs[usecase.BARBER_SHOP] = usecase.NewBarberShopUseCase(barberShopRepo, viewShopRepo)
	ucs[usecase.APPOINTMENT] = usecase.NewAppoinmentUseCase(appintmentRepo, slotRepo)
	ucs[usecase.CALENDAR] = usecase.NewCalendarUseCase(slotRepo)

	userRepo.Store(context.TODO(), &entity.User{
		Email:    "admin@admin.com",
		Password: "super_secret",
		Type:     entity.ADMIN,
	})

	// TODO: get the production env
	router := controller.Router(ucs, true)

	router.Run()
}
