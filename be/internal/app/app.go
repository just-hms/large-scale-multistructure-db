package app

import (
	"fmt"
	"large-scale-multistructure-db/be/internal/controller"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/internal/usecase/auth"
	"large-scale-multistructure-db/be/internal/usecase/repo"
	"large-scale-multistructure-db/be/pkg/mongo"
)

// TODO : fix this to devide the router from the rest, and put it in controllers

func Run() {
	// Repository

	mongo, err := mongo.New()
	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}

	// redis := redis.New()

	// UseCase

	usecases := []usecase.Usecase{
		usecase.NewUserUseCase(
			repo.NewUserRepo(mongo),
			auth.NewPasswordAuth(),
		),
	}

	router := controller.Router(usecases)

	router.Run()
}
