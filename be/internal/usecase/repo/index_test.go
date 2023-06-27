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

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
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
	adminAnalyticsRepo := repo.NewAdminAnalyticsRepo(mongo)

	err = repo.CreateTestIndexes(mongo, ctx)
	if err != nil {
		b.Fail()
	}
	users, err := userRepo.List(ctx, "")
	if err != nil {
		b.Fail()
	}
	barbershops, err := barberShopRepo.Find(ctx, -1, -1, "", 0)
	if err != nil {
		b.Fail()
	}

	fmt.Println("Number of users: ", len(users))
	fmt.Println("Number of barbershops: ", len(barbershops))

	b.Run("user-GetByEmail-with-index", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := userRepo.GetByEmail(ctx, users[len(users)-1].Email)
			if err != nil {
				fmt.Println(err)
				b.Fail()
			}
		}
	})

	b.Run("barbershop-Find-with-index", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := barberShopRepo.Find(ctx, 41.9027835, 12.4963655, "", 10000)
			if err != nil {
				fmt.Println(err)
				b.Fail()
			}
		}
	})

	b.Run("shopEgagementRanking-index", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := adminAnalyticsRepo.GetEngagementShopRanking(ctx)
			if err != nil {
				b.Fail()
			}
		}
	})

	err = repo.DropTestIndexes(mongo, ctx)
	if err != nil {
		b.Fail()
	}

	b.Run("user-GetByEmail", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := userRepo.GetByEmail(ctx, users[len(users)-1].Email)
			if err != nil {
				b.Fail()
			}
		}
	})

	b.Run("shopEgagementRanking", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := adminAnalyticsRepo.GetEngagementShopRanking(ctx)
			if err != nil {
				b.Fail()
			}
		}
	})
}
