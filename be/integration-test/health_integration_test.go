package integration_test

import (
	"net/http"
	"net/http/httptest"
)

func (s *IntegrationSuite) TestHealth() {

	req, _ := http.NewRequest("GET", "/api/health/", nil)

	w := httptest.NewRecorder()
	s.srv.ServeHTTP(w, req)

	s.Require().Equal(http.StatusOK, w.Code)
}
