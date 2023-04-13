package repo_test

import (
	"context"
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

func (s *RepoSuite) TestUserStore() {

	testCases := []struct {
		name      string
		user      *entity.User
		expectErr bool
	}{
		{
			name:      "store new user",
			user:      &entity.User{Email: "test@example.com", Username: "John Doe"},
			expectErr: false,
		},
		{
			name:      "store existing user",
			user:      &entity.User{Email: "test@example.com", Username: "Jane Doe"},
			expectErr: true,
		},
	}

	userRepo := repo.NewUserRepo(s.db)

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {
			err := userRepo.Store(context.Background(), tc.user)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Empty(tc.user.ID)
			} else {
				s.Require().NoError(err)
				s.Require().NotEmpty(tc.user.ID)
			}
		})
	}
}

func (s *RepoSuite) TestUserGetByID() {
	userRepo := repo.NewUserRepo(s.db)

	user := &entity.User{
		Email: "existinguser@gmail.com",
	}
	userRepo.Store(context.Background(), user)

	testCases := []struct {
		name      string
		ID        string
		expectErr bool
	}{
		{
			name:      "get existing user",
			ID:        user.ID,
			expectErr: false,
		},
		{
			name:      "get non-existent user",
			ID:        "non_existent_user_id",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			res, err := userRepo.GetByID(context.Background(), tc.ID)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Nil(res)
				return
			}

			s.Require().NoError(err)
			s.Require().NotNil(res)
		})
	}
}

func (s *RepoSuite) TestUserDeleteByID() {

	userRepo := repo.NewUserRepo(s.db)
	user := &entity.User{
		Email: "existinguser@gmail.com",
	}
	userRepo.Store(context.Background(), user)

	testCases := []struct {
		name      string
		ID        string
		expectErr bool
	}{
		{
			name:      "delete existing user",
			ID:        user.ID,
			expectErr: false,
		},
		{
			name:      "delete non-existent user",
			ID:        user.ID,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			err := userRepo.DeleteByID(context.Background(), tc.ID)
			if tc.expectErr {
				s.Require().Error(err)
				return
			}

			s.Require().NoError(err)

			// check if the user is correctly deleted
			res, err := userRepo.GetByID(context.Background(), user.ID)
			s.Require().NotNil(err)
			s.Require().Nil(res)
		})
	}
}

func (s *RepoSuite) TestUserList() {
	userRepo := repo.NewUserRepo(s.db)

	users := []*entity.User{
		{
			Email: "existinguser@gmail.com",
		},
		{
			Email: "kek@gmail.com",
		},
	}
	for _, user := range users {
		userRepo.Store(context.Background(), user)
	}

	testCases := []struct {
		name      string
		email     string
		expectLen int
	}{
		{
			name:      "list all users",
			email:     "",
			expectLen: 2,
		},
		{
			name:      "list users by email",
			email:     "kek",
			expectLen: 1,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			users, err := userRepo.List(context.Background(), tc.email)
			s.Require().NoError(err)
			s.Require().Len(users, tc.expectLen)
		})
	}
}

func (s *RepoSuite) TestUserGetByEmail() {
	userRepo := repo.NewUserRepo(s.db)

	user := &entity.User{
		Email: "existinguser@gmail.com",
	}
	userRepo.Store(context.Background(), user)

	testCases := []struct {
		name      string
		email     string
		expectErr bool
	}{
		{
			name:      "get existing user by email",
			email:     user.Email,
			expectErr: false,
		},
		{
			name:      "get non-existent user by email",
			email:     "non_existent_email@example.com",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			_, err := userRepo.GetByEmail(context.Background(), tc.email)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

func (s *RepoSuite) TestEditShopsByIDs() {
	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)

	shop := &entity.BarberShop{
		Name: "shop",
	}
	shopRepo.Store(context.Background(), shop)

	user := &entity.User{
		Username: "barbers",
	}
	userRepo.Store(context.Background(), user)

	testCases := []struct {
		name          string
		user          *entity.User
		baberbshopIDs []string
	}{
		{
			name:          "add a barbershop",
			user:          user,
			baberbshopIDs: []string{shop.ID},
		},
		{
			name:          "remove a barbershop",
			user:          user,
			baberbshopIDs: []string{},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			err := userRepo.EditShopsByIDs(context.Background(), tc.user, tc.baberbshopIDs)
			s.Require().NoError(err)

			modifiedUser, err := userRepo.GetByID(context.Background(), user.ID)
			s.Require().Equal(len(modifiedUser.OwnedShops), len(tc.baberbshopIDs))
		})
	}
}

func (s *RepoSuite) TestModifyByID() {
	userRepo := repo.NewUserRepo(s.db)

	user := &entity.User{
		Email:    "existinguser@gmail.com",
		Password: "old_password",
	}
	userRepo.Store(context.Background(), user)

	testCases := []struct {
		name      string
		ID        string
		expectErr bool
		mods      *entity.User
	}{
		{
			name:      "edit non-existent user",
			ID:        "non_existent_user_id",
			expectErr: true,
		},
		{
			name:      "edit email",
			ID:        user.ID,
			expectErr: false,
			mods: &entity.User{
				Email: "new_email",
			},
		},
		{
			name:      "edit password",
			ID:        user.ID,
			expectErr: false,
			mods: &entity.User{
				Password: "new_password",
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			err := userRepo.ModifyByID(context.Background(), tc.ID, tc.mods)

			if tc.expectErr {
				s.Require().Error(err)
				return
			}

			modifiedUser, err := userRepo.GetByID(context.Background(), user.ID)

			if tc.mods.Email != "" {
				s.Require().Equal(tc.mods.Email, modifiedUser.Email)
			} else {
				s.Require().NotEqual(tc.mods.Email, modifiedUser.Email)
			}

			if tc.mods.Password != "" {
				s.Require().Equal(tc.mods.Password, modifiedUser.Password)
			} else {
				s.Require().NotEqual(tc.mods.Email, modifiedUser.Password)
			}
		})
	}
}
