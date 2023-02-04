package integration_test

import (
	"large-scale-multistructure-db/be/internal/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
	srv *gin.Engine
}

func (s *IntegrationSuite) SetupTest() {
	s.srv = app.Router()
}

func (s *IntegrationSuite) TestHealth(t *testing.T) {

	req, _ := http.NewRequest("GET", "/health", nil)

	w := httptest.NewRecorder()
	s.srv.ServeHTTP(w, req)

	s.Equal(w.Code, http.StatusOK)
}
