package repo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/config"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
)

func BenchmarkIndexes(b *testing.B) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println("cannot find config")
		return
	}
	mongo, err := mongo.New(cfg.Mongo.Host, cfg.Mongo.Port, "barbershop")
	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}
	userRepo := repo.NewUserRepo(mongo)
	barberShopRepo := repo.NewBarberShopRepo(mongo)

	b.Run("user-GetByEmail-with-index", func(b *testing.B) {
		userRepo.GetByEmail(ctx, "admin@admin.com")
		users, _ := userRepo.List(ctx, "")
		fmt.Println(len(users))
	})

	b.Run("barbershop-Find-with-index", func(b *testing.B) {
		barberShopRepo.Find(ctx, 0.200, 0.200, "", 10000)
	})

	// TODO remove the index

	b.Run("user-GetByEmail", func(b *testing.B) {
		userRepo.GetByEmail(ctx, "admin@admin.com")
	})

	b.Run("barbershop-Find-with-index", func(b *testing.B) {
		barberShopRepo.Find(ctx, 0.200, 0.200, "", 10000)
	})

	// TODO add the index back

}
