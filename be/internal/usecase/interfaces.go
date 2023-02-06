package usecase

import (
	"context"
	"large-scale-multistructure-db/be/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// User -.
	User interface {
		Login(context.Context, *entity.User) (*entity.User, error)
		Store(context.Context, *entity.User) (string, error)

		GetByID(ctx context.Context, ID string) (*entity.User, error)
		ModifyByID(ctx context.Context, ID string, user *entity.User) error
		DeleteByID(ctx context.Context, ID string) error
		List(ctx context.Context, email string) ([]*entity.User, error)
	}

	// UserRepo -.
	UserRepo interface {
		Store(context.Context, *entity.User) (string, error)
		GetByID(context.Context, string) (*entity.User, error)
		DeleteByID(context.Context, string) error
		ModifyByID(context.Context, string, *entity.User) error
		GetByEmail(context.Context, string) (*entity.User, error)
		List(context.Context, string) ([]*entity.User, error)
	}

	PasswordAuth interface {
		Verify(string, string) bool
		HashAndSalt(string) (string, error)
	}

	Usecase interface{}
)
