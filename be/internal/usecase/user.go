package usecase

import (
	"context"
	"fmt"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase/auth"
)

// TranslationUseCase -.
type UserUseCase struct {
	repo UserRepo
}

// New -.
func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) Store(ctx context.Context, user *entity.User) error {

	if _, err := uc.repo.GetByEmail(ctx, user.Email); err != nil {
		return fmt.Errorf("Email already exists")
	}

	password, err := auth.HashAndSalt(user.Password)

	if err != nil {
		return err
	}

	user.Password = password

	return uc.repo.Store(ctx, user)
}

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (*entity.User, string, error) {

	var (
		user *entity.User
		err  error
	)

	if user, err = uc.repo.GetByEmail(ctx, email); err != nil {
		return nil, "", err
	}

	if !auth.Verify(user.Password, password) {
		return nil, "", fmt.Errorf("Wrong password")
	}

	token, err := auth.CreateToken(user.ID)

	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (uc *UserUseCase) GetByToken(ctx context.Context, token string) (*entity.User, error) {

	userID, err := auth.ExtractTokenID(token)

	if err != nil {
		return nil, err
	}

	return uc.repo.GetByID(ctx, userID)
}
