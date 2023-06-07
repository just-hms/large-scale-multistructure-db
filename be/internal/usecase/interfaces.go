package usecase

import (
	"context"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

type Usecase interface{}

const (
	USER byte = iota
	PASSWORD_AUTH
	BARBER_SHOP
	CALENDAR
	APPOINTMENT
	HOLIDAY
	REVIEW
	GEOCODING
	TOKEN
)

// Usecase interfaces

type (
	User interface {
		Login(ctx context.Context, user *entity.User) (*entity.User, error)
		Store(ctx context.Context, user *entity.User) error
		GetByID(ctx context.Context, ID string) (*entity.User, error)
		ModifyByID(ctx context.Context, ID string, user *entity.User) error
		DeleteByID(ctx context.Context, ID string) error
		List(ctx context.Context, email string) ([]*entity.User, error)
		EditShopsByIDs(ctx context.Context, ID string, IDs []string) error

		LostPassword(ctx context.Context, email string) (string, error)
		ResetPassword(ctx context.Context, ID string, newPassword string) error
	}

	BarberShop interface {
		Find(ctx context.Context, lat float64, lon float64, name string, radius float64) ([]*entity.BarberShop, error)
		GetByID(ctx context.Context, viewerID string, ID string) (*entity.BarberShop, error)
		GetOwnedShops(ctx context.Context, user *entity.User) ([]*entity.BarberShop, error)
		Store(ctx context.Context, shop *entity.BarberShop) error
		ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error
		DeleteByID(ctx context.Context, ID string) error
		GetShopAnalytics(ctx context.Context, ID string) (*entity.BarberAnalytics, error)
	}

	Calendar interface {
		GetByBarberShopID(ctx context.Context, ID string) (*entity.Calendar, error)
	}

	Appointment interface {
		Book(ctx context.Context, appointment *entity.Appointment) error
		CancelFromUser(ctx context.Context, userID string, appointment *entity.Appointment) error
		CancelFromShop(ctx context.Context, shopID string, appointment *entity.Appointment) error
		SetCompletedFromShop(ctx context.Context, shopID string, appointment *entity.Appointment) error
		GetByID(ctx context.Context, ID string) (*entity.Appointment, error)
	}

	Review interface {
		Store(ctx context.Context, review *entity.Review, shopID string) error
		GetByBarberShopID(ctx context.Context, shopID string) ([]*entity.Review, error)
		DeleteByID(ctx context.Context, reviewID string) error

		UpVoteByID(ctx context.Context, userID, shopID, reviewID string) error
		DownVoteByID(ctx context.Context, userID, shopID, reviewID string) error
		RemoveVoteByID(ctx context.Context, userID, shopID, reviewID string) error
	}

	Geocoding interface {
		Search(ctx context.Context, area string) ([]entity.GeocodingInfo, error)
	}

	Token interface {
		CreateToken(userID string) (string, error)
		ExtractTokenID(tokenString string) (string, error)
	}
)

// Utility interfaces
type (
	GeocodingWebAPI interface {
		Search(ctx context.Context, area string) ([]entity.GeocodingInfo, error)
	}
	PasswordAuth interface {
		Verify(hashed string, password string) bool
		HashAndSalt(password string) (string, error)
	}
	TokenApi interface {
		CreateToken(id string) (string, error)
		ExtractTokenID(token string) (string, error)
	}
)

// Repositories
type (
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
		GetOwnedShops(ctx context.Context, user *entity.User) ([]*entity.BarberShop, error)
		ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error
		DeleteByID(ctx context.Context, ID string) error
	}

	SlotRepo interface {
		GetByBarberShopID(ctx context.Context, ID string) ([]string, []*entity.Slot, error)
		Book(ctx context.Context, appointment *entity.Appointment, slot *entity.Slot) error
		Get(ctx context.Context, shopID string, date time.Time) (*entity.Slot, error)
		Cancel(ctx context.Context, appointment *entity.Appointment) error
		SetEmployees(ctx context.Context, shopID string, availableEmployees int) error
	}

	ShopViewRepo interface {
		Store(ctx context.Context, view *entity.ShopView) error
	}

	AppointmentRepo interface {
		Book(ctx context.Context, appointment *entity.Appointment) error
		SetStatusFromUser(ctx context.Context, userID string, appointment *entity.Appointment) error
		SetStatusFromShop(ctx context.Context, shopID string, appointment *entity.Appointment) error
		GetByID(ctx context.Context, ID string) (*entity.Appointment, error)
	}

	ReviewRepo interface {
		Store(ctx context.Context, review *entity.Review, shopID string) error
		GetByBarberShopID(ctx context.Context, shopID string) ([]*entity.Review, error)
		DeleteByID(ctx context.Context, reviewID string) error
	}

	VoteRepo interface {
		UpVoteByID(ctx context.Context, userID, shopID, reviewID string) error
		DownVoteByID(ctx context.Context, userID, shopID, reviewID string) error
		RemoveVoteByID(ctx context.Context, userID, shopID, reviewID string) error
	}

	BarberAnalyticsRepo interface {
		GetAppointmentCountByShop(ctx context.Context, shopID string) (map[string]int, error)
		GetViewCountByShop(ctx context.Context, shopID string) (map[string]int, error)
		GetReviewCountByShop(ctx context.Context, shopID string) (map[string]int, error)
		GetAppointmentCancellationRatioByShop(ctx context.Context, shopID string) (map[string]float64, error)
		GetAppointmentViewRatioByShop(ctx context.Context, shopID string) (map[string]float64, error)
		GetUpDownVoteCountByShop(ctx context.Context, shopID string) (map[string]map[string]int, error)
		GetReviewWeightedRatingByShop(ctx context.Context, shopID string) (float64, error)
		GetInactiveUsersByShop(ctx context.Context, shopID string) ([]string, error)
	}

	AdminAnalyticsRepo interface {
		GetAppointmentCount(ctx context.Context, shopID string) (map[string]int, error)
		GetViewCount(ctx context.Context, shopID string) (map[string]int, error)
		GetReviewCount(ctx context.Context, shopID string) (map[string]int, error)
	}
)
