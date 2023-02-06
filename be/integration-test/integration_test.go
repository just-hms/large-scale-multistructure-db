package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"large-scale-multistructure-db/be/internal/controller"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/internal/usecase/auth"
	"large-scale-multistructure-db/be/internal/usecase/repo"
	"large-scale-multistructure-db/be/pkg/jwt"
	"large-scale-multistructure-db/be/pkg/mongo"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
	srv   *gin.Engine
	mongo *mongo.Mongo

	resetDB func()
	params  map[string]string
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) SetupTest() {
	s.resetDB()
}

func (s *IntegrationSuite) SetupSuite() {

	fmt.Println(">>> From SetupSuite")

	mongo, err := mongo.New(&mongo.Options{
		DB_NAME: "test",
	})

	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}

	// create repos and usecases
	userRepo := repo.NewUserRepo(mongo)
	password := auth.NewPasswordAuth()
	usecases := []usecase.Usecase{
		usecase.NewUserUseCase(
			userRepo,
			password,
		),
	}

	// Fill the test DB

	s.resetDB = func() {
		mongo.DB.Drop(context.TODO())

		p, _ := password.HashAndSalt("password")

		var ID string

		ID, _ = userRepo.Store(context.TODO(), &entity.User{
			Email:    "correct@example.com",
			Password: p,
			IsAdmin:  false,
		})

		s.params["authID"] = ID
		s.params["authToken"], _ = jwt.CreateToken(ID)

		ID, _ = userRepo.Store(context.TODO(), &entity.User{
			Email:    "admin@example.com",
			Password: p,
			IsAdmin:  true,
		})

		s.params["adminID"] = ID
		s.params["adminToken"], _ = jwt.CreateToken(ID)

		userRepo.Store(context.TODO(), &entity.User{
			Email:    "to.filter@example.com",
			Password: p,
			IsAdmin:  true,
		})

	}

	// serv the mock server and db
	s.params = make(map[string]string)
	s.srv = controller.Router(usecases)
	s.mongo = mongo
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
			loginUser: &entity.User{Email: "correct@example.com", Password: "password"},
			status:    http.StatusOK,
		},
		{
			name:      "Invalid input",
			loginUser: &entity.User{Email: "not_an_email", Password: "password"},
			status:    http.StatusBadRequest,
		},
		{
			name:      "Wrong password",
			loginUser: &entity.User{Email: "correct@example.com", Password: "invalid"},
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

func (s *IntegrationSuite) TestRegister() {

	testCases := []struct {
		name         string
		creationUser *entity.User
		status       int
	}{
		{
			name:         "Already exists",
			creationUser: &entity.User{Email: "correct@example.com", Password: "password"},
			status:       http.StatusUnauthorized,
		},
		{
			name:         "Invalid input",
			creationUser: &entity.User{Email: "not_an_email", Password: "password"},
			status:       http.StatusBadRequest,
		},
		{
			name:         "Correctly Created",
			creationUser: &entity.User{Email: "new@example.com", Password: "password"},
			status:       http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		loginUserJson, _ := json.Marshal(tc.creationUser)

		// create a request for the register endpoint
		req, _ := http.NewRequest("POST", "/user/", bytes.NewBuffer(loginUserJson))
		req.Header.Set("Content-Type", "application/json")

		// serve the request to the test server
		w := httptest.NewRecorder()
		s.srv.ServeHTTP(w, req)

		// assert that the response status code is as expected
		s.Require().Equal(tc.status, w.Code)
	}
}

func (s *IntegrationSuite) TestShowSelf() {

	testCases := []struct {
		name   string
		token  string
		status int
	}{
		{
			name:   "Wrongly formatted token",
			token:  "wrong_token",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correct token",
			token:  s.params["authToken"],
			status: http.StatusOK,
		},
	}

	for _, tc := range testCases {

		// create a request for the self endpoint
		req, _ := http.NewRequest("GET", "/user/self", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+tc.token)

		w := httptest.NewRecorder()
		s.srv.ServeHTTP(w, req)

		// assert that the response status code is as expected
		s.Require().Equal(tc.status, w.Code)

		// check for user len if the request was accepted
		if w.Code == http.StatusAccepted {

			body, err := io.ReadAll(w.Body)

			// require no error in reading the response
			s.Require().Nil(err)

			type response struct {
				User entity.User `json:"user"`
			}

			var res response
			err = json.Unmarshal(body, &res)
			s.Require().Nil(err)

		}
	}
}

func (s *IntegrationSuite) TestDeleteSelf() {

	testCases := []struct {
		name   string
		token  string
		status int
	}{
		{
			name:   "Wrongly formatted token",
			token:  "wrong_token",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly deleted",
			token:  s.params["authToken"],
			status: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {

		fmt.Println("lol", s.params["authToken"])
		// create a request for the self endpoint
		req, _ := http.NewRequest("DELETE", "/user/self", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+tc.token)

		// serve the request to the test server
		w := httptest.NewRecorder()
		s.srv.ServeHTTP(w, req)

		// assert that the response status code is as expected
		s.Require().Equal(tc.status, w.Code)
	}
}

func (s *IntegrationSuite) TestShowAll() {

	testCases := []struct {
		name        string
		token       string
		filter      string
		resultCount int
		status      int
	}{
		{
			name:   "Wrongly formatted token",
			token:  "wrong_token",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Not an admin",
			token:  s.params["authToken"],
			status: http.StatusUnauthorized,
		},
		{
			name:        "Correctly shown without filter",
			token:       s.params["adminToken"],
			status:      http.StatusOK,
			resultCount: 3,
		},
		{
			name:        "Correctly shown with filter",
			filter:      "filter",
			token:       s.params["adminToken"],
			status:      http.StatusOK,
			resultCount: 1,
		},
	}

	for _, tc := range testCases {

		// create a request for the self endpoint
		var req *http.Request

		if tc.filter != "" {
			req, _ = http.NewRequest("GET", "/admin/user?email="+tc.filter, nil)
		} else {
			req, _ = http.NewRequest("GET", "/admin/user?email=", nil)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+tc.token)

		// serve the request to the test server
		w := httptest.NewRecorder()
		s.srv.ServeHTTP(w, req)

		// assert that the response status code is as expected
		s.Require().Equal(tc.status, w.Code)

		// check for user len if the request was accepted
		if w.Code == http.StatusAccepted {

			body, err := io.ReadAll(w.Body)

			// require no error in reading the response
			s.Require().Nil(err)

			type response struct {
				Users []entity.User `json:"users"`
			}

			var res response

			err = json.Unmarshal(body, &res)
			s.Require().Nil(err)

			// assert that the number of returned user is as expected
			s.Require().Equal(tc.resultCount, len(res.Users))
		}

	}
}

func (s *IntegrationSuite) TestShow() {

	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
	}{
		{
			name:   "Wrongly formatted token",
			token:  "wrong_token",
			ID:     "generic_ID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Not an admin",
			token:  s.params["authToken"],
			ID:     "generic_ID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly shown user",
			token:  s.params["adminToken"],
			ID:     s.params["authID"],
			status: http.StatusOK,
		},
		{
			name:   "User not exists",
			token:  s.params["adminToken"],
			ID:     "wrong_ID",
			status: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {

		// create a request for the self endpoint
		var req *http.Request

		req, _ = http.NewRequest("GET", "/admin/user/"+tc.ID, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+tc.token)

		// serve the request to the test server
		w := httptest.NewRecorder()
		s.srv.ServeHTTP(w, req)

		// assert that the response status code is as expected
		s.Require().Equal(tc.status, w.Code)

		// check for user len if the request was accepted
		if w.Code == http.StatusAccepted {

			body, err := io.ReadAll(w.Body)

			// require no error in reading the response
			s.Require().Nil(err)

			type response struct {
				User entity.User `json:"user"`
			}

			var res response
			err = json.Unmarshal(body, &res)
			s.Require().Nil(err)

		}

	}
}

func (s *IntegrationSuite) TestDelete() {

	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
	}{
		{
			name:   "Wrongly formatted token",
			token:  "wrong_token",
			ID:     "generic_ID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Not an admin",
			token:  s.params["authToken"],
			ID:     "generic_ID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly deleted user",
			token:  s.params["adminToken"],
			ID:     s.params["authID"],
			status: http.StatusAccepted,
		},
		{
			name:   "User not exists",
			token:  s.params["adminToken"],
			ID:     "wrong_ID",
			status: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {

		// create a request for the self endpoint
		var req *http.Request

		req, _ = http.NewRequest("DELETE", "/admin/user/"+tc.ID, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+tc.token)

		// serve the request to the test server
		w := httptest.NewRecorder()
		s.srv.ServeHTTP(w, req)

		// assert that the response status code is as expected
		s.Require().Equal(tc.status, w.Code)

	}
}
