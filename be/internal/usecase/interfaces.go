package usecase

import (
	"context"
	"large-scale-multistructure-db/be/internal/entity"
	"time"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (

	// INTERFACES

	User interface {
		Login(ctx context.Context, user *entity.User) (*entity.User, error)     // IMPLEMENTED | TESTED
		Store(ctx context.Context, user *entity.User) (string, error)           // IMPLEMENTED | TESTED
		GetByID(ctx context.Context, ID string) (*entity.User, error)           // IMPLEMENTED | TESTED
		ModifyByID(ctx context.Context, ID string, user *entity.User) error     // IMPLEMENTED | TESTED
		DeleteByID(ctx context.Context, ID string) error                        // IMPLEMENTED | TESTED
		List(ctx context.Context, email string) ([]*entity.User, error)         // IMPLEMENTED | TESTED
		LostPassword(ctx context.Context, email string) (string, error)         // IMPLEMENTED
		ResetPassword(ctx context.Context, ID string, newPassword string) error // IMPLEMENTED
	}

	PasswordAuth interface {
		Verify(hashed string, password string) bool  // IMPLEMENTED
		HashAndSalt(password string) (string, error) // IMPLEMENTED
	}

	BarberShop interface {
		Find(ctx context.Context, lat string, lon string, name string, radius string) ([]*entity.BarberShop, error) // IMPLEMENTED
		GetByID(ctx context.Context, viewerID string, ID string) (*entity.BarberShop, error)                        // IMPLEMENTED
		Store(ctx context.Context, shop *entity.BarberShop) (string, error)                                         // IMPLEMENTED
		ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error                                   // IMPLEMENTED
		DeleteByID(ctx context.Context, ID string) error                                                            // IMPLEMENTED
	}

	Calendar interface {
		GetByBarberShopID(ctx context.Context, ID string) (*entity.Calendar, error)
	}

	Appointment interface {
		Book(ctx context.Context, appointment *entity.Appointment) (string, error)
		Cancel(ctx context.Context, appointment *entity.Appointment) (string, error)
		DeleteByID(ctx context.Context, ID string) (string, error)
	}

	Holiday interface {
		Set(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) (string, error)
	}

	Review interface {
		Store(ctx context.Context, userID string, shopID string) (string, error)
		GetByBarberShop(ctx context.Context, shopID string) ([]*entity.Review, error)
		DeleteByID(ctx context.Context, ID string) error
		VoteByID(ctx context.Context, ID string) error
	}

	// TODO : add analytics, maybe raw access to db using custom store for each one

	// REPOS

	// UserRepo -.
	UserRepo interface {
		Store(ctx context.Context, user *entity.User) (string, error)
		GetByID(ctx context.Context, ID string) (*entity.User, error)
		DeleteByID(ctx context.Context, ID string) error
		ModifyByID(ctx context.Context, ID string, user *entity.User) error
		GetByEmail(ctx context.Context, email string) (*entity.User, error)
		List(ctx context.Context, email string) ([]*entity.User, error)
	}

	BarberShopRepo interface {
		Store(ctx context.Context, shop *entity.BarberShop) (string, error)
		Find(ctx context.Context, lat string, lon string, name string, radius string) ([]*entity.BarberShop, error)
		GetByID(ctx context.Context, ID string) (*entity.BarberShop, error)
		ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error
		DeleteByID(ctx context.Context, ID string) error
	}

	// TODO: add something that refresh the slots every day
	SlotRepo interface {
		GetByBarberShopID(ctx context.Context, ID string) ([]*entity.Slot, error)
		Book(ctx context.Context, appointment *entity.Appointment) (string, error)
		Cancel(ctx context.Context, appointment *entity.Appointment) (string, error)
		SetHoliday(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) (string, error)
	}

	ShopViewRepo interface {
		Store(ctx context.Context, view *entity.ShopView) (string, error)
	}

	AppointmentRepo interface {
		Book(ctx context.Context, appointment *entity.Appointment) (string, error)
		Cancel(ctx context.Context, appointment *entity.Appointment) (string, error)
		DeleteByID(ctx context.Context, ID string) (string, error)
	}

	ReviewRepo interface {
		Store(ctx context.Context, userID string, shopID string) (string, error)
		GetByBarberShop(ctx context.Context, shopID string) ([]*entity.Review, error)
		DeleteByID(ctx context.Context, ID string) error
		VoteByID(ctx context.Context, ID string) error
	}

	Usecase interface{}
)
