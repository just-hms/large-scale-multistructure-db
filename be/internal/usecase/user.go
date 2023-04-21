package usecase

import (
	"context"
	"errors"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/jwt"
)

// TranslationUseCase -.
type UserUseCase struct {
	repo     UserRepo
	password PasswordAuth
}

var ErrUserAlreadyExists = errors.New("email already exists")

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
		return ErrUserAlreadyExists
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
		return nil, errors.New("wrong password")
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

func (uc *UserUseCase) EditShopsByIDs(ctx context.Context, ID string, IDs []string) error {

	user, err := uc.repo.GetByID(ctx, ID)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Type == entity.ADMIN {
		if len(IDs) > 0 {
			return errors.New("cannot edit shops if the user is admin")
		}
		return nil
	}

	return uc.repo.EditShopsByIDs(ctx, user, IDs)
}

func (uc *UserUseCase) ModifyByID(ctx context.Context, ID string, update *entity.User) error {
	if _, err := uc.repo.GetByEmail(ctx, update.Email); err == nil {
		return ErrUserAlreadyExists
	}

	if update.Password != "" {
		return errors.New("the password field cannot be edited here")
	}

	if update.Type == entity.ADMIN && len(update.OwnedShops) > 0 {
		return errors.New("cannot become an admin if you own shops")
	}

	return uc.repo.ModifyByID(ctx, ID, update)
}
