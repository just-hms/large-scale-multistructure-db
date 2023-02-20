package integration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"large-scale-multistructure-db/be/internal/controller"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/pkg/jwt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func (s *IntegrationSuite) TestLogin() {

	testCases := []struct {
		name      string
		loginUser *controller.RegisterInput
		status    int
	}{
		{
			name:      "Correct Login",
			loginUser: &controller.RegisterInput{Email: "correct@example.com", Password: "password"},
			status:    http.StatusOK,
		},
		{
			name:      "Invalid input",
			loginUser: &controller.RegisterInput{Email: "not_an_email", Password: "password"},
			status:    http.StatusBadRequest,
		},
		{
			name:      "Wrong password",
			loginUser: &controller.RegisterInput{Email: "correct@example.com", Password: "invalid"},
			status:    http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			loginUserJson, _ := json.Marshal(tc.loginUser)

			// create a request for the login endpoint
			req, _ := http.NewRequest("POST", "/api/user/login/", bytes.NewBuffer(loginUserJson))
			req.Header.Set("Content-Type", "application/json")

			// serve the request to the test server
			w := httptest.NewRecorder()
			s.srv.ServeHTTP(w, req)

			if w.Code == http.StatusOK {

				body, err := io.ReadAll(w.Body)

				// require no error in reading the response
				s.Require().Nil(err)

				type response struct {
					Token string `json:"token"`
				}

				var res response

				err = json.Unmarshal(body, &res)
				s.Require().Nil(err)

				// check if the tokenID is in the jwt token
				tokenID, err := jwt.ExtractTokenID(res.Token)

				s.Require().Nil(err)
				s.Require().NotEmpty(tokenID)

			}

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)
		})
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

		s.T().Run(tc.name, func(t *testing.T) {
			loginUserJson, _ := json.Marshal(tc.creationUser)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/user/", bytes.NewBuffer(loginUserJson))
			req.Header.Set("Content-Type", "application/json")

			// serve the request to the test server
			w := httptest.NewRecorder()
			s.srv.ServeHTTP(w, req)

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)
		})
	}
}

func (s *IntegrationSuite) TestShowSelf() {

	testCases := []struct {
		name           string
		token          string
		status         int
		barberShopsLen int
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
		{
			name:           "Show barber",
			token:          s.params["barber1Auth"],
			status:         http.StatusOK,
			barberShopsLen: 1,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the self endpoint
			req, _ := http.NewRequest("GET", "/api/user/self/", nil)
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

				s.Require().Len(res.User.OwnedShops, tc.barberShopsLen)

			}
		})

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

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the self endpoint
			req, _ := http.NewRequest("DELETE", "/api/user/self/", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+tc.token)

			// serve the request to the test server
			w := httptest.NewRecorder()
			s.srv.ServeHTTP(w, req)

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)
		})

	}
}

func (s *IntegrationSuite) TestUserShowAll() {

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

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the self endpoint
			var req *http.Request

			query := url.Values{
				"email": {tc.filter},
			}

			req, _ = http.NewRequest("GET", "/api/admin/user?email="+query.Encode(), nil)
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
		})

	}
}

func (s *IntegrationSuite) TestUserShow() {

	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
	}{
		{
			name:   "Wrongly formatted token",
			token:  "wrong_token",
			ID:     "genericID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Not an admin",
			token:  s.params["authToken"],
			ID:     "genericID",
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

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the self endpoint
			var req *http.Request

			req, _ = http.NewRequest("GET", "/api/admin/user/"+tc.ID, nil)
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
		})

	}
}

func (s *IntegrationSuite) TestUserDelete() {

	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
	}{
		{
			name:   "Wrongly formatted token",
			token:  "wrong_token",
			ID:     "genericID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Not an admin",
			token:  s.params["authToken"],
			ID:     "genericID",
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

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the self endpoint
			var req *http.Request

			req, _ = http.NewRequest("DELETE", "/api/admin/user/"+tc.ID, nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+tc.token)

			// serve the request to the test server
			w := httptest.NewRecorder()
			s.srv.ServeHTTP(w, req)

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)
		})

	}
}
