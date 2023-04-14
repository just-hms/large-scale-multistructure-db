package repo_test

import (
	"context"
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
)

func (s *RepoSuite) TestBarberShopStore() {

	testCases := []struct {
		name      string
		shop      *entity.BarberShop
		expectErr bool
	}{
		{
			name:      "store new barbershop",
			shop:      &entity.BarberShop{Name: "brownies", Location: entity.FAKE_LOCATION},
			expectErr: false,
		},
		{
			name:      "store existing barbershop",
			shop:      &entity.BarberShop{Name: "brownies", Location: entity.FAKE_LOCATION},
			expectErr: true,
		},
	}

	shopRepo := repo.NewBarberShopRepo(s.db)

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {
			err := shopRepo.Store(context.Background(), tc.shop)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Empty(tc.shop.ID)
			} else {
				s.Require().NoError(err)
				s.Require().NotEmpty(tc.shop.ID)
			}
		})
	}
}

func (s *RepoSuite) TestBarberShopFind() {
	shopRepo := repo.NewBarberShopRepo(s.db)

	shops := []*entity.BarberShop{
		{Name: "hair brownies", Location: entity.NewLocation(14.1234, 24.5678)},
		{Name: "haircut place", Location: entity.NewLocation(11.1234, 22.5678)},
		{Name: "cut and shave", Location: entity.NewLocation(11.1334, 22.5679)},
	}

	for _, shop := range shops {
		err := shopRepo.Store(context.Background(), shop)
		s.Require().NoError(err)
	}

	// Test cases
	testCases := []struct {
		name         string
		lat          float64
		lon          float64
		nameFilter   string
		radiusFilter float64
		expectedLen  int
	}{
		{
			name:        "find all",
			expectedLen: 3,
		},
		{
			name:        "find barbershops with name filter",
			nameFilter:  "hair",
			expectedLen: 2,
		},
		{
			name:         "find barbershops within radius",
			lat:          14.1234,
			lon:          24.5678,
			radiusFilter: 1000,
			expectedLen:  1,
		},
		{
			name:         "find barbershops with both radius and name filters",
			nameFilter:   "hair",
			lat:          11.1234,
			lon:          22.5678,
			radiusFilter: 100000,
			expectedLen:  1,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			shops, err := shopRepo.Find(context.Background(), tc.lat, tc.lon, tc.nameFilter, tc.radiusFilter)
			s.Require().NoError(err)
			s.Require().Len(shops, tc.expectedLen)
		})
	}
}

func (s *RepoSuite) TestBarberShopGetByID() {

	shopRepo := repo.NewBarberShopRepo(s.db)
	shop := &entity.BarberShop{
		Name: "Test Barber Shop", Phone: "555-555-5555",
		Location: entity.FAKE_LOCATION,
	}
	shopRepo.Store(context.Background(), shop)

	testCases := []struct {
		name      string
		ID        string
		expectErr bool
	}{
		{
			name:      "get existing barbershop",
			ID:        shop.ID,
			expectErr: false,
		},
		{
			name:      "get non-existent barbershop",
			ID:        "non-existent-id",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			barber, err := shopRepo.GetByID(context.Background(), tc.ID)

			if tc.expectErr {
				s.Require().Error(err)
				s.Require().Nil(barber)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(barber)
				s.Require().Equal(shop.ID, barber.ID)
				s.Require().Equal(shop.Name, barber.Name)
				s.Require().Equal(shop.Phone, barber.Phone)
			}
		})
	}
}

func (s *RepoSuite) TestBarberShopDeleteByID() {

	shopRepo := repo.NewBarberShopRepo(s.db)
	shop := &entity.BarberShop{
		Name: "brownies",
	}
	shopRepo.Store(context.Background(), shop)

	testCases := []struct {
		name      string
		ID        string
		expectErr bool
	}{
		{
			name:      "delete existing shop",
			ID:        shop.ID,
			expectErr: false,
		},
		{
			name:      "delete non-existent shop",
			ID:        shop.ID,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			err := shopRepo.DeleteByID(context.Background(), tc.ID)
			if tc.expectErr {
				s.Require().Error(err)
				return
			}

			s.Require().NoError(err)

			// check if the barbershop is correctly deleted
			res, err := shopRepo.GetByID(context.Background(), shop.ID)
			s.Require().NotNil(err)
			s.Require().Nil(res)
		})
	}
}

func (s *RepoSuite) TestBarberShopModifyByID() {
	barberRepo := repo.NewBarberShopRepo(s.db)

	shop := &entity.BarberShop{
		Name: "brownies",
	}
	barberRepo.Store(context.Background(), shop)

	testCases := []struct {
		name      string
		ID        string
		expectErr bool
		mods      *entity.BarberShop
	}{
		{
			name:      "edit non-existent barbershop",
			ID:        "non_existent_shop_id",
			expectErr: true,
		},
		{
			name:      "edit name",
			ID:        shop.ID,
			expectErr: false,
			mods: &entity.BarberShop{
				Name: "brownies2.0",
			},
		},
		{
			name:      "edit location",
			ID:        shop.ID,
			expectErr: false,
			mods: &entity.BarberShop{
				Location: entity.NewLocation(-2, -3),
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			err := barberRepo.ModifyByID(context.Background(), tc.ID, tc.mods)

			if tc.expectErr {
				s.Require().Error(err)
				return
			}

			modifiedShop, err := barberRepo.GetByID(context.Background(), shop.ID)
			s.Require().NoError(err)

			if tc.mods.Name != "" {
				s.Require().Equal(tc.mods.Name, modifiedShop.Name)
			} else {
				s.Require().NotEqual(tc.mods.Name, modifiedShop.Name)
			}

			if tc.mods.Location != nil {
				s.Require().Equal(tc.mods.Location, modifiedShop.Location)
			} else {
				s.Require().NotEqual(tc.mods.Location, modifiedShop.Location)
			}
		})
	}
}
