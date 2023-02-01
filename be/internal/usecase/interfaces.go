package usecase

import (
	"context"
	"large-scale-multistructure-db/be/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// User -.
	User interface {
		Store(context.Context, *entity.User) error
		Login(context.Context, string, string) (*entity.User, string, error)
		GetByToken(context.Context, string) (*entity.User, error)
	}

	// UserRepo -.
	UserRepo interface {
		Store(context.Context, *entity.User) error
		GetByID(context.Context, uint) (*entity.User, error)
		GetByEmail(context.Context, string) (*entity.User, error)
	}
)
