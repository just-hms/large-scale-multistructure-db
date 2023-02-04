package usecase

import (
	"context"
	"errors"
	"fmt"
	"large-scale-multistructure-db/be/internal/entity"
)

// TranslationUseCase -.
type UserUseCase struct {
	repo     UserRepo
	password PasswordAuth
}

var UserAlreadyExists = errors.New("Email already exists")

// New -.
func NewUserUseCase(r UserRepo, p PasswordAuth) *UserUseCase {
	return &UserUseCase{
		repo:     r,
		password: p,
	}
}

func (uc *UserUseCase) Store(ctx context.Context, user *entity.User) error {

	// TODO: if there is any other error it creates the user

	// TODO : standard the error codes

	if _, err := uc.repo.GetByEmail(ctx, user.Email); err == nil {
		return UserAlreadyExists
	}

	// TODO : add check in pwd len

	hashedPassword, err := uc.password.HashAndSalt(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return uc.repo.Store(ctx, user)
}

func (uc *UserUseCase) Login(ctx context.Context, loginUser *entity.User) (*entity.User, error) {

	user, err := uc.repo.GetByEmail(ctx, loginUser.Email)

	if err != nil {
		return nil, err
	}

	if !uc.password.Verify(user.Password, loginUser.Password) {
		return nil, fmt.Errorf("Wrong password")
	}

	return user, nil
}

func (uc *UserUseCase) DeleteByID(ctx context.Context, ID uint) error {
	return uc.repo.DeleteByID(ctx, ID)
}

func (uc *UserUseCase) GetByID(ctx context.Context, ID uint) (*entity.User, error) {
	return uc.repo.GetByID(ctx, ID)
}

func (uc *UserUseCase) ModifyByID(ctx context.Context, ID uint, user *entity.User) error {

	// TODO: don't edit the password here

	if _, err := uc.repo.GetByEmail(ctx, user.Email); err == nil {
		return UserAlreadyExists
	}

	return uc.repo.ModifyByID(ctx, ID, user)
}

func (uc *UserUseCase) List(ctx context.Context, email string) ([]*entity.User, error) {
	return uc.repo.List(ctx, email)
}
