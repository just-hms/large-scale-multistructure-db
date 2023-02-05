package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"large-scale-multistructure-db/be/internal/controller"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/internal/usecase/auth"
	"large-scale-multistructure-db/be/internal/usecase/repo"
	"large-scale-multistructure-db/be/pkg/mongo"
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

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) SetupSuite() {

	fmt.Println(">>> From SetupSuite")

	// TODO : add a test DB
	mongo, err := mongo.New()
	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}

	usecases := []usecase.Usecase{
		usecase.NewUserUseCase(
			repo.NewUserRepo(mongo),
			auth.NewPasswordAuth(),
		),
	}

	s.srv = controller.Router(usecases)
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

	testCases := []struct {
		name      string
		loginUser *entity.User
		status    int
	}{
		{
			name:      "Correct Login",
			loginUser: &entity.User{Email: "test@example.com", Password: "password"},
			status:    http.StatusOK,
		},
		{
			name:      "Invalid input",
			loginUser: &entity.User{Email: "not_an_email", Password: "password"},
			status:    http.StatusBadRequest,
		},
		{
			name:      "Wrong password",
			loginUser: &entity.User{Email: "test@example.com", Password: "invalid"},
			status:    http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		loginUserJson, _ := json.Marshal(tc.loginUser)

		// create a request for the login endpoint
		req, _ := http.NewRequest("POST", "/user/login", bytes.NewBuffer(loginUserJson))
		req.Header.Set("Content-Type", "application/json")

		// serve the request to the test server
		w := httptest.NewRecorder()
		s.srv.ServeHTTP(w, req)

		// assert that the response status code is as expected
		s.Require().Equal(tc.status, w.Code)
	}
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

	// assert that the response status code is 201 OK
	s.Require().Equal(http.StatusCreated, w.Code)
}
