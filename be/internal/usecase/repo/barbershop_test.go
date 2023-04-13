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
			shop:      &entity.BarberShop{Name: "brownies"},
			expectErr: false,
		},
		{
			name:      "store existing barbershop",
			shop:      &entity.BarberShop{Name: "brownies"},
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
		{Name: "brownies", Latitude: "10.1234", Longitude: "20.5678"},
		{Name: "haircut place", Latitude: "11.1234", Longitude: "22.5678"},
		{Name: "cut and shave", Latitude: "12.1234", Longitude: "24.5678"},
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
			name:         "find barbershops with name filter",
			lat:          10.1234,
			lon:          20.5678,
			radiusFilter: -1,
			nameFilter:   "brownies",
			expectedLen:  1,
		},
		{
			name:         "find barbershops within radius",
			lat:          10.1234,
			lon:          20.5678,
			radiusFilter: -1,
			expectedLen:  1,
		},
		{
			name:         "find barbershops with both radius and name filters",
			lat:          10.1234,
			lon:          20.5678,
			radiusFilter: -1,
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
	shop := &entity.BarberShop{Name: "Test Barber Shop", Phone: "555-555-5555"}
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

			// check if the user is correctly deleted
			res, err := shopRepo.GetByID(context.Background(), shop.ID)
			s.Require().NotNil(err)
			s.Require().Nil(res)
		})
	}
}

func (s *RepoSuite) TestBarberShopModifyByID() {
	s.Fail("not implemented")
}
