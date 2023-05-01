package usecase

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

type GeocodingUseCase struct {
	geocodinapi GeocodingWebAPI
}

func NewGeocodingUseCase(g GeocodingWebAPI) *GeocodingUseCase {
	return &GeocodingUseCase{
		geocodinapi: g,
	}
}

func (uc *GeocodingUseCase) Search(ctx context.Context, area string) ([]entity.GeocodingInfo, error) {
	return uc.geocodinapi.Search(ctx, area)
}
