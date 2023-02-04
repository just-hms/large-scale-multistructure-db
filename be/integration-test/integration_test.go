package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"large-scale-multistructure-db/be/internal/app"
	"large-scale-multistructure-db/be/internal/entity"
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

// listen for 'go test' command --> run test methods
func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) SetupSuite() {

	// TODO : add a test DB

	fmt.Println(">>> From SetupSuite")
	s.srv = app.Router()
}

func (s *IntegrationSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

func (s *IntegrationSuite) TestHealth() {

	req, _ := http.NewRequest("GET", "/health", nil)

	w := httptest.NewRecorder()
	s.srv.ServeHTTP(w, req)

	s.Require().Equal(http.StatusOK, w.Code)
}

func (s *IntegrationSuite) TestLogin() {

	loginUser := &entity.User{Email: "test@example.com", Password: "password"}
	loginUserJson, _ := json.Marshal(loginUser)

	// create a request for the login endpoint
	req, _ := http.NewRequest("POST", "/user/login", bytes.NewBuffer(loginUserJson))
	req.Header.Set("Content-Type", "application/json")

	// serve the request to the test server
	w := httptest.NewRecorder()
	s.srv.ServeHTTP(w, req)

	// assert that the response status code is 200 OK
	s.Require().NotEqual(http.StatusBadRequest, w.Code)
}

func (s *IntegrationSuite) TestCreate() {

	createUser := &entity.User{Email: "test@example.com", Password: "password"}
	createUserJson, _ := json.Marshal(createUser)

	// create a request for the login endpoint
	req, _ := http.NewRequest("POST", "/user/", bytes.NewBuffer(createUserJson))
	req.Header.Set("Content-Type", "application/json")

	// serve the request to the test server
	w := httptest.NewRecorder()
	s.srv.ServeHTTP(w, req)

	// assert that the response status code is 200 OK
	s.Require().Equal(http.StatusCreated, w.Code)
}
