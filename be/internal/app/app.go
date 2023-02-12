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

	// redis := redis.New()

	// UseCase

	userUsecase := usecase.NewUserUseCase(
		repo.NewUserRepo(mongo),
		auth.NewPasswordAuth(),
	)

	usecases := []usecase.Usecase{
		userUsecase,
		usecase.NewBarberShopUseCase(
			repo.NewBarberShopRepo(mongo),
			repo.NewShopViewRepo(mongo),
		),
	}

	userUsecase.Store(context.TODO(), &entity.User{
		Email:    "admin@admin.com",
		Password: "super_secret",
		Type:     entity.ADMIN,
	})

	// TODO create an account admin

	router := controller.Router(usecases)

	router.Run()
}
