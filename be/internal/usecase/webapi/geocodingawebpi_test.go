package webapi_test

import (
	"context"
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/config"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/webapi"
	"github.com/stretchr/testify/suite"
)

type GeocodingAPISuite struct {
	suite.Suite
}

func TestGeocodingAPISuite(t *testing.T) {
	suite.Run(t, new(GeocodingAPISuite))
}

func (s *GeocodingAPISuite) TestSearch() {
	cfg, err := config.NewConfig()
	s.Require().NoError(err)

	api, err := webapi.NewGeocodingWebAPI(cfg.Geocoding.Apikey)
	s.Require().NoError(err)
	res, err := api.Search(context.Background(), "via brombeis naples")
	s.Require().NoError(err)

	s.Require().Len(res, 5)

	s.Require().Equal(res[0], entity.GeocodingInfo{
		Country:   "Italy",
		Region:    "Campania",
		City:      "Naples",
		Latitude:  40.8506395,
		Longitude: 14.2487991,
		Address:   "Via Giovanni Brombeis, 80135 Naples NA, Italy",
	})
}
