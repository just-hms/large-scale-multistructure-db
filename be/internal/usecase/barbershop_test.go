package usecase_test

import (
	"context"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func shopHelper(t *testing.T) (*usecase.BarberShopUseCase, *MockBarberShopRepo, *MockShopViewRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockBarberShopRepo(mockCtl)
	view := NewMockShopViewRepo(mockCtl)

	shop := usecase.NewBarberShopUseCase(repo, view)

	return shop, repo, view
}

func TestBarberShopGetByID(t *testing.T) {
	t.Parallel()

	shop, repo, view := shopHelper(t)

	tests := []struct {
		name     string
		mock     func()
		ID       string
		viewerID string
		res      *entity.BarberShop
		err      error
	}{
		{
			name:     "barbeshop not found",
			ID:       "1",
			viewerID: "2",
			mock: func() {
				repo.EXPECT().GetByID(context.Background(), "1").Return(nil, errInternalServErr)
			},
			res: nil,
			err: errInternalServErr,
		},
		{
			name:     "barbeshop found",
			ID:       "1",
			viewerID: "3",
			mock: func() {
				repo.EXPECT().GetByID(context.Background(), "1").Return(&entity.BarberShop{}, nil)
				view.EXPECT().Store(context.Background(), &entity.ShopView{
					ViewerID:     "3",
					BarberShopID: "1",
				}).Return("", nil)
			},
			res: &entity.BarberShop{},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := shop.GetByID(context.Background(), tc.viewerID, tc.ID)

			if err == nil {
				require.NotNil(t, res)
			}

			require.ErrorIs(t, tc.err, err)
		})
	}
}
