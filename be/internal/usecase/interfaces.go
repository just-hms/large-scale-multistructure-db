package usecase

import (
	"context"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type Usecase interface{}

const (
	USER int = iota
	PASSWORD_AUTH
	BARBER_SHOP
	CALENDAR
	APPOINTMENT
	HOLIDAY
	REVIEW
	LEN
)

type (

	// INTERFACES

	User interface {
		Login(ctx context.Context, user *entity.User) (*entity.User, error) // IMPLEMENTED | TESTED
		Store(ctx context.Context, user *entity.User) error                 // IMPLEMENTED | TESTED
		GetByID(ctx context.Context, ID string) (*entity.User, error)       // IMPLEMENTED | TESTED
		ModifyByID(ctx context.Context, ID string, user *entity.User) error // IMPLEMENTED | TESTED
		DeleteByID(ctx context.Context, ID string) error                    // IMPLEMENTED | TESTED
		List(ctx context.Context, email string) ([]*entity.User, error)     // IMPLEMENTED | TESTED
		EditShopsByIDs(ctx context.Context, ID string, IDs []string) error

		LostPassword(ctx context.Context, email string) (string, error)         // IMPLEMENTED
		ResetPassword(ctx context.Context, ID string, newPassword string) error // IMPLEMENTED
	}

	PasswordAuth interface {
		Verify(hashed string, password string) bool  // IMPLEMENTED
		HashAndSalt(password string) (string, error) // IMPLEMENTED
	}

	BarberShop interface {
		Find(ctx context.Context, lat float64, lon float64, name string, radius float64) ([]*entity.BarberShop, error) // IMPLEMENTED
		GetByID(ctx context.Context, viewerID string, ID string) (*entity.BarberShop, error)                           // IMPLEMENTED
		Store(ctx context.Context, shop *entity.BarberShop) error                                                      // IMPLEMENTED
		ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error                                      // IMPLEMENTED
		DeleteByID(ctx context.Context, ID string) error                                                               // IMPLEMENTED
	}

	Calendar interface {
		GetByBarberShopID(ctx context.Context, ID string) (*entity.Calendar, error) // IMPLEMENTED
	}

	Appointment interface {
		Book(ctx context.Context, appointment *entity.Appointment) error   // IMPLEMENTED
		Cancel(ctx context.Context, appointment *entity.Appointment) error // IMPLEMENTED
	}

	Holiday interface {
		Set(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) error // IMPLEMENTED
	}

	Review interface {
		Store(ctx context.Context, userID string, shopID string) error
		GetByBarberShop(ctx context.Context, shopID string) ([]*entity.Review, error)
		DeleteByID(ctx context.Context, ID string) error

		UpVoteByID(ctx context.Context, voterID, ID string) error
		DownVoteByID(ctx context.Context, voterID, ID string) error
		RemoveVoteByID(ctx context.Context, voterID, ID string) error
	}

	// TODO : add analytics, maybe raw access to db using custom store for each one

	// REPOS

	// UserRepo -.
	UserRepo interface {
		Store(ctx context.Context, user *entity.User) error
		GetByID(ctx context.Context, ID string) (*entity.User, error)
		DeleteByID(ctx context.Context, ID string) error
		ModifyByID(ctx context.Context, ID string, user *entity.User) error
		GetByEmail(ctx context.Context, email string) (*entity.User, error)
		List(ctx context.Context, email string) ([]*entity.User, error)
		EditShopsByIDs(ctx context.Context, user *entity.User, IDs []string) error
	}

	BarberShopRepo interface {
		Store(ctx context.Context, shop *entity.BarberShop) error
		Find(ctx context.Context, lat float64, lon float64, name string, radius float64) ([]*entity.BarberShop, error)
		GetByID(ctx context.Context, ID string) (*entity.BarberShop, error)
		ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error
		DeleteByID(ctx context.Context, ID string) error
	}

	SlotRepo interface {
		GetByBarberShopID(ctx context.Context, ID string) ([]*entity.Slot, error)
		Book(ctx context.Context, appointment *entity.Appointment) error
		Cancel(ctx context.Context, appointment *entity.Appointment) error
		SetHoliday(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) error
	}

	ShopViewRepo interface {
		Store(ctx context.Context, view *entity.ShopView) error
	}

	AppointmentRepo interface {
		Book(ctx context.Context, appointment *entity.Appointment) error
		Cancel(ctx context.Context, appointment *entity.Appointment) error
	}

	ReviewRepo interface {
		Store(ctx context.Context, userID string, shopID string) error
		GetByBarberShop(ctx context.Context, shopID string) ([]*entity.Review, error)
		DeleteByID(ctx context.Context, ID string) error
	}

	VoteRepo interface {
		DownVote(ctx context.Context, voterID string, shopID string) error
		UpVote(ctx context.Context, voterID string, shopID string) error
		RemoveVote(ctx context.Context, voterID string, shopID string) error
	}
)
