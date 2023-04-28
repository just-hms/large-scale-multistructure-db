package usecase

import (
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

type PlaceUseCase struct {
	geocodinapi GeocodingWebAPI
}

func (uc *PlaceUseCase) Search(area string) ([]entity.GeocodingInfo, error) {
	return uc.geocodinapi.Search(area)
}
