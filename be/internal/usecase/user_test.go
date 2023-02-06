package usecase_test

import (
	"context"
	"errors"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func userHelper(t *testing.T) (*usecase.UserUseCase, *MockUserRepo, *MockPasswordAuth) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	password := NewMockPasswordAuth(mockCtl)
	repo := NewMockUserRepo(mockCtl)

	user := usecase.NewUserUseCase(repo, password)

	return user, repo, password
}

var errInternalServErr = errors.New("internal server error")

func TestLogin(t *testing.T) {
	t.Parallel()

	user, repo, password := userHelper(t)

	tests := []struct {
		name string
		mock func()

		res interface{}

		input *entity.User
		err   error
	}{
		{
			name: "user not found",
			input: &entity.User{
				Email: "wrong_email",
			},
			mock: func() {
				repo.EXPECT().GetByEmail(context.Background(), "wrong_email").Return(nil, errInternalServErr)
			},
			res: nil,
			err: errInternalServErr,
		},
		{
			name: "user found but wrong password",
			input: &entity.User{
				Email:    "correct_email",
				Password: "c",
			},
			mock: func() {
				repo.EXPECT().GetByEmail(context.Background(), "correct_email").
					Return(&entity.User{Email: "correct_email", Password: "hashed_password"}, errInternalServErr)

				password.EXPECT().Verify("hashed_password", "wrong_password").Return(false)
			},
			res: nil,
			err: errInternalServErr,
		},
		{
			name: "user found and right password",
			input: &entity.User{
				Email:    "correct_email",
				Password: "correct_password",
			},
			mock: func() {
				repo.EXPECT().GetByEmail(context.Background(), "correct_email").
					Return(&entity.User{Email: "correct_email", Password: "hashed_password"}, errInternalServErr)

				password.EXPECT().Verify("hashed_password", "correct_password").Return(true)
			},
			res: &entity.User{},

			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := user.Login(context.Background(), tc.input)

			if err == nil {
				require.NotNil(t, res)
			}

			require.ErrorIs(t, tc.err, err)
		})
	}
}
func TestStore(t *testing.T) {
	t.Parallel()

	user, repo, password := userHelper(t)

	tests := []struct {
		name string
		mock func()

		res interface{}

		input *entity.User
		err   error
	}{
		{
			name: "user already exists",
			input: &entity.User{
				Email:    "existing_email",
				Password: "password",
			},
			mock: func() {
				repo.EXPECT().GetByEmail(context.Background(), "existing_email").Return(&entity.User{}, nil)
			},
			err: usecase.UserAlreadyExists,
		},
		{
			name: "correctly saved",
			input: &entity.User{
				Email:    "new_email",
				Password: "password",
			},
			mock: func() {

				repo.EXPECT().GetByEmail(context.Background(), "new_email").Return(nil, errInternalServErr)
				password.EXPECT().HashAndSalt("password").Return("hashed_password", nil)

				repo.EXPECT().Store(context.Background(),
					&entity.User{
						Email:    "new_email",
						Password: "hashed_password",
					},
				).Return("", nil)
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			_, err := user.Store(context.Background(), tc.input)

			require.ErrorIs(t, tc.err, err)
		})
	}
}

func TestModifyByID(t *testing.T) {
	t.Parallel()

	user, repo, _ := userHelper(t)

	var tests = []struct {
		name   string
		input  string
		user   *entity.User
		mock   func()
		expect error
	}{
		{
			name:  "new email already exists",
			input: "1",
			user: &entity.User{
				Email:    "existing_email",
				Password: "password",
			},
			mock: func() {
				repo.EXPECT().GetByEmail(context.Background(), "existing_email").Return(&entity.User{}, nil)
			},
			expect: usecase.UserAlreadyExists,
		},
		{
			name:  "modify success",
			input: "2",
			user: &entity.User{
				Email:    "new_email",
				Password: "password",
			},
			mock: func() {
				repo.EXPECT().GetByEmail(context.Background(), "new_email").Return(nil, errInternalServErr)
				repo.EXPECT().ModifyByID(context.Background(), "2", &entity.User{
					Email:    "new_email",
					Password: "password",
				}).Return(nil)
			},
			expect: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			err := user.ModifyByID(context.Background(), tc.input, tc.user)

			require.ErrorIs(t, tc.expect, err)
		})
	}
}
