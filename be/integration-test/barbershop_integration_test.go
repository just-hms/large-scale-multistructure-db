package integration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

func (s *IntegrationSuite) TestBarberShopFind() {

	testCases := []struct {
		input       controller.FindBarbershopInput
		name        string
		radius      string
		token       string
		status      int
		resultCount int
	}{
		{
			name:   "Unauthorized",
			status: http.StatusUnauthorized,
		},
		{
			name:        "All barbershops",
			status:      http.StatusOK,
			token:       s.params[USER1_TOKEN],
			resultCount: 3,
		},
		{
			name: "No barbershop near where you are",
			input: controller.FindBarbershopInput{
				Latitude:  1.1,
				Longitude: 1.1,
				Radius:    1,
			},
			token:       s.params[USER1_TOKEN],
			status:      http.StatusOK,
			resultCount: 0,
		},
		{
			name: "No barbershop for this name",
			input: controller.FindBarbershopInput{
				Name: "not_existing_shop",
			},
			token:       s.params[USER1_TOKEN],
			status:      http.StatusOK,
			resultCount: 0,
		},
		// TODO:FIX THIS
		{
			name: "Two barbershop found for this name",
			input: controller.FindBarbershopInput{
				Name: "boh",
			},
			token:       s.params[USER1_TOKEN],
			status:      http.StatusOK,
			resultCount: 2,
		},
		{
			name: "Found 1 barbershop near you",
			input: controller.FindBarbershopInput{
				Latitude:  1.1,
				Longitude: 1.1,
				Radius:    100,
			},
			token:       s.params[USER1_TOKEN],
			status:      http.StatusOK,
			resultCount: 1,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			newBarberShopJson, _ := json.Marshal(tc.input)

			// create a request for the self endpoint
			var req *http.Request
			req, _ = http.NewRequest("POTS", "/api/barber_shop", bytes.NewBuffer(newBarberShopJson))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+tc.token)

			// serve the request to the test server
			w := httptest.NewRecorder()
			s.srv.ServeHTTP(w, req)

			if w.Code == http.StatusAccepted {

				body, err := io.ReadAll(w.Body)

				// require no error in reading the response
				s.Require().Nil(err)

				type response struct {
					BarberShops []entity.User `json:"barberShops"`
				}

				var res response

				err = json.Unmarshal(body, &res)
				s.Require().Nil(err)

				// assert that the number of returned user is as expected
				s.Require().Equal(tc.resultCount, len(res.BarberShops))
			}

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)
		})

	}
}
func (s *IntegrationSuite) TestBarberShopShow() {
	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
	}{
		{
			name:   "Wrongly formatted token",
			token:  "wrong_token",
			ID:     "wrongID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly shown Barbershop",
			token:  s.params[USER1_TOKEN],
			ID:     s.params[SHOP1_ID],
			status: http.StatusOK,
		},
		{
			name:   "BarberShop doesn't exist",
			token:  s.params[USER1_TOKEN],
			ID:     "wrong_ID",
			status: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the self endpoint
			var req *http.Request

			req, _ = http.NewRequest("GET", "/api/barber_shop/"+tc.ID, nil)
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
					Barber entity.BarberShop `json:"barber"`
				}

				var res response
				err = json.Unmarshal(body, &res)
				s.Require().Nil(err)
			}
		})
	}
}
func (s *IntegrationSuite) TestBarberShopStore() {

	testCases := []struct {
		name   string
		token  string
		input  *controller.CreateBarbershopInput
		status int
	}{
		{
			name: "Require login",
			input: &controller.CreateBarbershopInput{
				Name:            "barberShop7",
				Latitude:        1,
				Longitude:       1,
				EmployeesNumber: 2,
			},
			status: http.StatusUnauthorized,
		},
		{
			name:  "Require admin",
			token: s.params[USER1_ID],
			input: &controller.CreateBarbershopInput{
				Name:            "barberShop7",
				Latitude:        1,
				Longitude:       1,
				EmployeesNumber: 2,
			},
			status: http.StatusUnauthorized,
		},
		{
			name:  "Already exists",
			token: s.params[ADMIN_TOKEN],
			input: &controller.CreateBarbershopInput{
				Name:            "barberShop1",
				Latitude:        1,
				Longitude:       1,
				EmployeesNumber: 2,
			},
			status: http.StatusBadRequest,
		},
		{
			name:   "Invalid input",
			token:  s.params[ADMIN_TOKEN],
			status: http.StatusBadRequest,
		},
		{
			name:  "Correctly Created",
			token: s.params[ADMIN_TOKEN],
			input: &controller.CreateBarbershopInput{
				Name:            "barberShop7",
				Latitude:        1.1,
				Longitude:       1,
				EmployeesNumber: 2,
			},
			status: http.StatusCreated,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			newBarberShopJson, _ := json.Marshal(tc.input)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/admin/barber_shop", bytes.NewBuffer(newBarberShopJson))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+tc.token)

			// serve the request to the test server
			w := httptest.NewRecorder()
			s.srv.ServeHTTP(w, req)

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)

			// TODO: do a request to check that the creation went successfully
		})
	}
}
func (s *IntegrationSuite) TestBarberShopModifyByID() {
	testCases := []struct {
		name   string
		token  string
		ID     string
		input  *controller.ModifyBarberShopInput
		status int
	}{
		{
			name:   "Require login",
			input:  &controller.ModifyBarberShopInput{Employees: 2},
			ID:     "genericID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Require barber",
			token:  s.params[USER1_ID],
			input:  &controller.ModifyBarberShopInput{Employees: 3},
			status: http.StatusUnauthorized,
			ID:     "genericID",
		},
		{
			name:   "Require to be a barber in that shop",
			token:  s.params[BARBER1_TOKEN],
			input:  &controller.ModifyBarberShopInput{Employees: 1},
			status: http.StatusUnauthorized,
			ID:     s.params[SHOP2_ID],
		},
		{
			name:   "Correctly Modified",
			token:  s.params[BARBER1_TOKEN],
			input:  &controller.ModifyBarberShopInput{Employees: 2},
			status: http.StatusAccepted,
			ID:     s.params[SHOP1_ID],
		},

		// TODO: test more bulky editing
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			editBarberShopJson, _ := json.Marshal(tc.input)

			// create a request for the register endpoint
			req, _ := http.NewRequest("PUT", "/api/barber_shop/"+tc.ID, bytes.NewBuffer(editBarberShopJson))
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
func (s *IntegrationSuite) TestBarberShopDeleteByID() {

	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
	}{
		{
			name:   "Require Login",
			status: http.StatusUnauthorized,
			ID:     s.params[SHOP2_ID],
		},
		{
			name:   "Require admin",
			token:  s.params[USER1_TOKEN],
			ID:     s.params[SHOP2_ID],
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly Eliminated",
			token:  s.params[ADMIN_TOKEN],
			ID:     s.params[SHOP2_ID],
			status: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the register endpoint
			req, _ := http.NewRequest("DELETE", "/api/admin/barber_shop/"+tc.ID, nil)

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
