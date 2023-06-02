package app

import (
	"fmt"

	"github.com/just-hms/large-scale-multistructure-db/be/config"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/auth"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/tokenapi"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/webapi"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/redis"
)

func Run(cfg *config.Config) {

	mongo, err := mongo.New(cfg.Mongo.Host, cfg.Mongo.Port, "barbershop")
	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}
	if err := repo.AddIndexes(mongo); err != nil {
		fmt.Printf("index-error: %s", err.Error())
		return
	}

	redis, err := redis.New(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password)
	if err != nil {
		fmt.Printf("redis-error: %s", err.Error())
		return
	}

	ucs, err := BuildRequirements(mongo, redis, cfg)
	if err != nil {
		fmt.Printf("build-error: %s", err.Error())
		return
	}

	router := controller.Router(ucs, true)

	router.Run()
}

func BuildRequirements(m *mongo.Mongo, r *redis.Redis, cfg *config.Config) (map[byte]usecase.Usecase, error) {

	userRepo := repo.NewUserRepo(m)
	barberShopRepo := repo.NewBarberShopRepo(m)
	viewShopRepo := repo.NewShopViewRepo(m)
	appointmentRepo := repo.NewAppointmentRepo(m)
	reviewRepo := repo.NewReviewRepo(m)
	voteRepo := repo.NewVoteRepo(m)
	slotRepo := repo.NewSlotRepo(r)

	barberAnalyticsRepo := repo.NewBarberAnalyticsRepo(m)

	password := auth.NewPasswordAuth()
	tokenapi := tokenapi.New(cfg.TokenLifespan, cfg.ApiSecret)

	// TODO : move check before and pass conf matrix
	search_api, err := webapi.NewGeocodingWebAPI(cfg.Geocoding.Apikey)
	if err != nil {
		return nil, err
	}
	ucs := map[byte]usecase.Usecase{}
	ucs[usecase.USER] = usecase.NewUserUseCase(userRepo, password, tokenapi)
	ucs[usecase.BARBER_SHOP] = usecase.NewBarberShopUseCase(barberShopRepo, viewShopRepo, barberAnalyticsRepo, slotRepo)
	ucs[usecase.APPOINTMENT] = usecase.NewAppoinmentUseCase(
		appointmentRepo,
		slotRepo,
		userRepo,
		barberShopRepo,
	)
	ucs[usecase.REVIEW] = usecase.NewReviewUseCase(reviewRepo, voteRepo)
	ucs[usecase.CALENDAR] = usecase.NewCalendarUseCase(slotRepo)
	ucs[usecase.GEOCODING] = usecase.NewGeocodingUseCase(search_api)

	ucs[usecase.TOKEN] = usecase.NewTokenUsecase(tokenapi)

	return ucs, nil
}
