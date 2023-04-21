package repo_test

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

func (s *RepoSuite) TestShopViewStore() {

	user := &entity.User{Username: "giovanni"}
	shop := &entity.BarberShop{Name: "brownies"}

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)

	err := userRepo.Store(context.Background(), user)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop)
	s.Require().NoError(err)

	testCases := []struct {
		name      string
		input     *entity.ShopView
		expectErr bool
	}{
		{
			name: "store new slotview",
			input: &entity.ShopView{
				ViewerID:     user.ID,
				BarberShopID: shop.ID,
			},
			expectErr: false,
		},
	}

	shopViewRepo := repo.NewShopViewRepo(s.db)

	for _, tc := range testCases {

		s.Run(tc.name, func() {
			err := shopViewRepo.Store(context.Background(), tc.input)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}
