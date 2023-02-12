package app

import (
	"context"
	"fmt"
	"large-scale-multistructure-db/be/internal/controller"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/internal/usecase/auth"
	"large-scale-multistructure-db/be/internal/usecase/repo"
	"large-scale-multistructure-db/be/pkg/mongo"
	"large-scale-multistructure-db/be/pkg/redis"
)

// TODO : fix this to devide the router from the rest, and put it in controllers

func Run() {
	// Repository

	mongo, err := mongo.New(&mongo.Options{
		DB_NAME: "barber-deploy",
	})

	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}

	redis := redis.New(&redis.RedisOptions{})

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
