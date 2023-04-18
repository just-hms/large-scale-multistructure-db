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

	ucs := BuildRequirements(mongo, redis)

	userUsecase := ucs[usecase.USER].(usecase.User)

	userUsecase.Store(context.Background(), &entity.User{
		Email:    "admin@admin.com",
		Password: "super_secret",
		Type:     entity.ADMIN,
	})

	// TODO: get the production env
	router := controller.Router(ucs, true)

	router.Run()
}

func BuildRequirements(m *mongo.Mongo, r *redis.Redis) map[byte]usecase.Usecase {

	userRepo := repo.NewUserRepo(m)
	barberShopRepo := repo.NewBarberShopRepo(m)
	viewShopRepo := repo.NewShopViewRepo(m)
	appintmentRepo := repo.NewAppointmentRepo(m)
	slotRepo := repo.NewSlotRepo(r)

	password := auth.NewPasswordAuth()

	ucs := map[byte]usecase.Usecase{}
	ucs[usecase.USER] = usecase.NewUserUseCase(userRepo, password)
	ucs[usecase.BARBER_SHOP] = usecase.NewBarberShopUseCase(barberShopRepo, viewShopRepo)
	ucs[usecase.APPOINTMENT] = usecase.NewAppoinmentUseCase(appintmentRepo, slotRepo, userRepo)
	ucs[usecase.CALENDAR] = usecase.NewCalendarUseCase(slotRepo)
	ucs[usecase.HOLIDAY] = usecase.NewHolidayUseCase(slotRepo)

	return ucs
}
