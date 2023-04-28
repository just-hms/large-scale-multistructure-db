package geocodinapi_test

import (
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/geocodinapi"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/env"
	"github.com/stretchr/testify/suite"
)

type GeocodingAPISuite struct {
	suite.Suite
}

func TestGeocodingAPISuite(t *testing.T) {
	suite.Run(t, new(GeocodingAPISuite))
}

func (s *GeocodingAPISuite) TestSearch() {
	apikey, err := env.GetString("GEOCODE_API_SECRET")
	s.Require().NoError(err)
	a := geocodinapi.New(apikey)
	_, err = a.Search("firenze")
	s.Require().NoError(err)
}
