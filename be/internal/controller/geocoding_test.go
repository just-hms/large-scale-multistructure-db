package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

func (s *ControllerSuite) TestGeocodingSearch() {

	searchJSON, _ := json.Marshal(controller.SearchInput{
		Area: "via giovanni Roma",
	})

	req, _ := http.NewRequest("POST", "/api/geocoding/search", bytes.NewBuffer(searchJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.fixture[USER1_TOKEN])

	w := httptest.NewRecorder()
	s.srv.ServeHTTP(w, req)

	s.Require().Equal(http.StatusOK, w.Code)

	body, err := io.ReadAll(w.Body)
	s.Require().Nil(err)

	type response struct {
		Geocodes []entity.GeocodingInfo `json:"geocodes"`
	}

	var res response
	err = json.Unmarshal(body, &res)
	s.Require().Nil(err)

	s.Require().Len(res.Geocodes, 5)
}
