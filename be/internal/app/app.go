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

// TODO : fix this to devide the router from the rest, and put it in controllers

func Run() {
	// Repository

	mongo, err := mongo.New(&mongo.Options{
		DBName: "barber-deploy",
	})

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

	userUseCase := usecase.NewUserUseCase(
		userRepo,
		password,
	)

	barberUseCase := usecase.NewBarberShopUseCase(
		barberShopRepo,
		viewShopRepo,
	)

	appointmentUseCase := usecase.NewAppoinmentUseCase(
		appintmentRepo,
		slotRepo,
	)
	calendarUseCase := usecase.NewCalendarUseCase(
		slotRepo,
	)

	userUseCase.Store(context.TODO(), &entity.User{
		Email:    "admin@admin.com",
		Password: "super_secret",
		Type:     entity.ADMIN,
	})

	router := controller.Router(
		userUseCase,
		barberUseCase,
		appointmentUseCase,
		calendarUseCase,
	)

	router.Run()
}
