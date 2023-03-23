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
	})

	// TODO create an account admin

	router := controller.Router(usecases)

	router.Run()
}
