package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/jwt"
)

func (s *ControllerSuite) TestLogin() {

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

		s.Run(tc.name, func() {

			loginUserJson, _ := json.Marshal(tc.loginUser)

			// create a request for the login endpoint
			req, _ := http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(loginUserJson))
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

func (s *ControllerSuite) TestRegister() {

	testCases := []struct {
		name         string
		creationUser *controller.RegisterInput
		status       int
	}{
		{
			name:         "Already exists",
			creationUser: &controller.RegisterInput{Email: "correct@example.com", Password: "password"},
			status:       http.StatusUnauthorized,
		},
		{
			name:         "Invalid input",
			creationUser: &controller.RegisterInput{Email: "not_an_email", Password: "password"},
			status:       http.StatusBadRequest,
		},
		{
			name:         "Correctly Created",
			creationUser: &controller.RegisterInput{Email: "new@example.com", Password: "password"},
			status:       http.StatusCreated,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {
			loginUserJson, _ := json.Marshal(tc.creationUser)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(loginUserJson))
			req.Header.Set("Content-Type", "application/json")

			// serve the request to the test server
			w := httptest.NewRecorder()
			s.srv.ServeHTTP(w, req)

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)
		})
	}
}

func (s *ControllerSuite) TestShowSelf() {

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
			token:  s.fixture[USER1_TOKEN],
			status: http.StatusOK,
		},
		{
			name:           "Show barber",
			token:          s.fixture[BARBER1_TOKEN],
			status:         http.StatusOK,
			barberShopsLen: 1,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			// create a request for the self endpoint
			req, _ := http.NewRequest("GET", "/api/user/self", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+tc.token)

			w := httptest.NewRecorder()
			s.srv.ServeHTTP(w, req)

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)

			// check for user barbershops len if the request was accepted
			if w.Code == http.StatusOK {

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

func (s *ControllerSuite) TestDeleteSelf() {

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
			token:  s.fixture[USER1_TOKEN],
			status: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			// create a request for the self endpoint
			req, _ := http.NewRequest("DELETE", "/api/user/self", nil)
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

func (s *ControllerSuite) TestUserShowAll() {

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
			token:  s.fixture[USER1_TOKEN],
			status: http.StatusUnauthorized,
		},
		{
			name:        "Correctly shown without filter",
			token:       s.fixture[ADMIN_TOKEN],
			status:      http.StatusOK,
			resultCount: 3,
		},
		{
			name:        "Correctly shown with filter",
			filter:      "filter",
			token:       s.fixture[ADMIN_TOKEN],
			status:      http.StatusOK,
			resultCount: 1,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

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

func (s *ControllerSuite) TestUserShow() {

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
			token:  s.fixture[USER1_TOKEN],
			ID:     "genericID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly shown user",
			token:  s.fixture[ADMIN_TOKEN],
			ID:     s.fixture[USER1_ID],
			status: http.StatusOK,
		},
		{
			name:   "User not exists",
			token:  s.fixture[ADMIN_TOKEN],
			ID:     "wrong_ID",
			status: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

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

func (s *ControllerSuite) TestUserDelete() {

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
			token:  s.fixture[USER1_TOKEN],
			ID:     "genericID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly deleted user",
			token:  s.fixture[ADMIN_TOKEN],
			ID:     s.fixture[USER1_ID],
			status: http.StatusAccepted,
		},
		{
			name:   "User not exists",
			token:  s.fixture[ADMIN_TOKEN],
			ID:     "wrong_ID",
			status: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

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

func (s *ControllerSuite) TestUserModify() {

	testCases := []struct {
		name           string
		token          string
		ID             string
		status         int
		barberShopsLen int
		input          controller.ModifyUserInput
	}{
		{
			name:   "Wrongly formatted token",
			token:  "wrong_token",
			ID:     "genericID",
			status: http.StatusUnauthorized,
			input:  controller.ModifyUserInput{},
		},
		{
			name:   "Not an admin",
			token:  s.fixture[USER1_TOKEN],
			ID:     "genericID",
			status: http.StatusUnauthorized,
			input:  controller.ModifyUserInput{},
		},
		{
			name:  "Set a shop",
			token: s.fixture[ADMIN_TOKEN],
			ID:    s.fixture[USER1_ID],
			input: controller.ModifyUserInput{
				BarbershopsID: []string{
					s.fixture[SHOP1_ID],
				},
			},
			barberShopsLen: 1,
			status:         http.StatusAccepted,
		},
		{
			name:  "Set two shops",
			token: s.fixture[ADMIN_TOKEN],
			ID:    s.fixture[USER1_ID],
			input: controller.ModifyUserInput{
				BarbershopsID: []string{
					s.fixture[SHOP1_ID],
					s.fixture[SHOP2_ID],
				},
			},
			barberShopsLen: 2,
			status:         http.StatusAccepted,
		},
		{
			name:   "Empty the user shops",
			token:  s.fixture[ADMIN_TOKEN],
			ID:     s.fixture[USER1_ID],
			status: http.StatusAccepted,
			input: controller.ModifyUserInput{
				BarbershopsID: []string{},
			},
			barberShopsLen: 0,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			modifyUserJson, _ := json.Marshal(tc.input)

			// create a request for the self endpoint
			var req *http.Request

			req, _ = http.NewRequest("PUT", "/api/admin/user/"+tc.ID, bytes.NewBuffer(modifyUserJson))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+tc.token)

			// serve the request to the test server
			w := httptest.NewRecorder()
			s.srv.ServeHTTP(w, req)

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)

			if w.Code == http.StatusAccepted {

				req, _ = http.NewRequest("GET", "/api/admin/user/"+tc.ID, nil)
				req.Header.Set("Content-Type", "application/json")
				req.Header.Add("Authorization", "Bearer "+tc.token)

				w2 := httptest.NewRecorder()
				s.srv.ServeHTTP(w2, req)

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
