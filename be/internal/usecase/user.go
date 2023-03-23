package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/jwt"
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

// TODO : add standard the error codes

func (uc *UserUseCase) Store(ctx context.Context, user *entity.User) error {

	if _, err := uc.repo.GetByEmail(ctx, user.Email); err == nil {
		return UserAlreadyExists
	}

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

func (uc *UserUseCase) DeleteByID(ctx context.Context, ID string) error {
	return uc.repo.DeleteByID(ctx, ID)
}

// TODO: retrieve also the ID and barberShopsIDs
func (uc *UserUseCase) GetByID(ctx context.Context, ID string) (*entity.User, error) {
	return uc.repo.GetByID(ctx, ID)
}

func (uc *UserUseCase) ModifyByID(ctx context.Context, ID string, user *entity.User) error {

	if _, err := uc.repo.GetByEmail(ctx, user.Email); err == nil {
		return UserAlreadyExists
	}

	if user.Password != "" {
		return fmt.Errorf("The password field cannot be edited here")
	}

	return uc.repo.ModifyByID(ctx, ID, user)
}

// TODO: retrieve also the ID and barberShopsIDs
func (uc *UserUseCase) List(ctx context.Context, email string) ([]*entity.User, error) {
	return uc.repo.List(ctx, email)
}

func (uc *UserUseCase) LostPassword(ctx context.Context, email string) (string, error) {
	user, err := uc.repo.GetByEmail(ctx, email)

	if err != nil {
		return "", err
	}

	resetToken, err := jwt.CreateToken(user.ID)

	if err != nil {
		return "", err
	}

	return resetToken, nil
}

func (uc *UserUseCase) ResetPassword(ctx context.Context, ID string, password string) error {

	hashed, err := uc.password.HashAndSalt(password)

	if err != nil {
		return err
	}

	return uc.repo.ModifyByID(ctx, ID, &entity.User{
		Password: hashed,
	})
}
