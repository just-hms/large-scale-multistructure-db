package repo_test

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

func (s *RepoSuite) TestReviewStore() {

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)

	user := &entity.User{Email: "giovanni"}
	shop := &entity.BarberShop{Name: "brownies"}

	err := userRepo.Store(context.Background(), user)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop)
	s.Require().NoError(err)

	testCases := []struct {
		name        string
		input       *entity.Review
		shopID      string
		expectedErr bool
	}{
		{
			name:        "Correctly inserted",
			expectedErr: false,
			shopID:      shop.ID,
			input: &entity.Review{
				Rating:  4,
				Content: "test",
				UserID:  user.ID,
			},
		},
		{
			name:        "The shop does not exists",
			expectedErr: true,
			shopID:      "test123",
			input: &entity.Review{
				Rating:  4,
				Content: "test",
				UserID:  user.ID,
			},
		},
		{
			name:        "The user does not exists",
			expectedErr: true,
			shopID:      shop.ID,
			input: &entity.Review{
				Rating:  4,
				Content: "test",
				UserID:  "test123",
			},
		},
	}

	reviewRepo := repo.NewReviewRepo(s.db)
	for _, tc := range testCases {

		s.Run(tc.name, func() {
			err = reviewRepo.Store(context.Background(), tc.input, tc.shopID)
			if tc.expectedErr {
				s.Require().Error(err)
				return
			}

			s.Require().NoError(err)
			// check that the review was correctly created
			// in the barbershop collection
			shop, err = shopRepo.GetByID(context.Background(), shop.ID)
			s.Require().NoError(err)
			s.Require().Len(shop.Reviews, 1)
		})
	}

}

func (s *RepoSuite) TestGetByBarberShopID() {

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)
	reviewRepo := repo.NewReviewRepo(s.db)

	user1 := &entity.User{Email: "giovanni"}
	user2 := &entity.User{Email: "mario"}
	shop := &entity.BarberShop{Name: "brownies"}

	err := userRepo.Store(context.Background(), user1)
	s.Require().NoError(err)
	err = userRepo.Store(context.Background(), user2)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop)
	s.Require().NoError(err)

	review1 := &entity.Review{
		Rating:  4,
		Content: "test1",
		UserID:  user1.ID,
	}
	review2 := &entity.Review{
		Rating:  5,
		Content: "test2",
		UserID:  user2.ID,
	}

	err = reviewRepo.Store(context.Background(), review1, shop.ID)
	s.Require().NoError(err)
	err = reviewRepo.Store(context.Background(), review2, shop.ID)
	s.Require().NoError(err)

	reviews, err := reviewRepo.GetByBarberShopID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(reviews, 2)

}

func (s *RepoSuite) TestReviewDelete() {

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)
	reviewRepo := repo.NewReviewRepo(s.db)

	user := &entity.User{Email: "giovanni"}
	shop := &entity.BarberShop{Name: "brownies"}

	err := userRepo.Store(context.Background(), user)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop)
	s.Require().NoError(err)

	review := &entity.Review{
		Rating:  4,
		Content: "test",
		UserID:  user.ID,
	}

	err = reviewRepo.Store(context.Background(), review, shop.ID)
	s.Require().NoError(err)

	// check that the review was correctly created
	// in the barbershop collection
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews, 1)

	err = reviewRepo.DeleteByID(context.Background(), shop.ID, review.ReviewID)
	s.Require().NoError(err)

	// check that the review was correctly deleted
	// in the barbershop collection
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews, 0)

}
