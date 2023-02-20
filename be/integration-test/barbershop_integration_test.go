package integration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"large-scale-multistructure-db/be/internal/controller"
	"large-scale-multistructure-db/be/internal/entity"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func (s *IntegrationSuite) TestBarberShopFind() {

	testCases := []struct {
		name           string
		lat            string
		lon            string
		radius         string
		token          string
		status         int
		barberShopName string
		resultCount    int
		res            interface{}
	}{
		{
			name:   "Unauthorized",
			status: http.StatusUnauthorized,
		},
		{
			name:        "All barbershops",
			status:      http.StatusOK,
			token:       s.params["authToken"],
			resultCount: 3,
		},
		{
			name:        "No barbershop near where you are",
			lat:         "1.1",
			lon:         "1.1",
			radius:      "1",
			token:       s.params["authToken"],
			status:      http.StatusOK,
			resultCount: 0,
		},
		{
			name:           "No barbershop for this name",
			barberShopName: "not_existing_shop",
			token:          s.params["authToken"],
			status:         http.StatusOK,
			resultCount:    0,
		},
		{
			name:           "Two barbershop found for this name",
			barberShopName: "not_existing_shop",
			token:          s.params["authToken"],
			status:         http.StatusOK,
			resultCount:    2,
		},
		{
			name:        "Found 1 barbershop near you",
			lat:         "1.1",
			lon:         "1.1",
			radius:      "100",
			token:       s.params["authToken"],
			status:      http.StatusOK,
			resultCount: 1,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the self endpoint
			var req *http.Request

			query := url.Values{
				"name":   {tc.barberShopName},
				"lat":    {tc.lat},
				"lon":    {tc.lon},
				"radius": {tc.radius},
			}

			req, _ = http.NewRequest("GET", "/api/barber_shop?"+query.Encode(), nil)

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
			token:  s.params["authToken"],
			ID:     s.params["barberShop1ID"],
			status: http.StatusOK,
		},
		{
			name:   "BarberShop doesn't exist",
			token:  s.params["authToken"],
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
		name          string
		token         string
		newBarberShop *controller.CreateBarbershopInput
		status        int
	}{
		{
			name: "Require login",
			newBarberShop: &controller.CreateBarbershopInput{
				Name:            "barberShop7",
				Latitude:        1,
				Longitude:       1,
				EmployeesNumber: 2,
			},
			status: http.StatusUnauthorized,
		},
		{
			name:  "Require admin",
			token: s.params["authToken"],
			newBarberShop: &controller.CreateBarbershopInput{
				Name:            "barberShop7",
				Latitude:        1,
				Longitude:       1,
				EmployeesNumber: 2,
			},
			status: http.StatusUnauthorized,
		},
		{
			name:  "Already exists",
			token: s.params["adminToken"],
			newBarberShop: &controller.CreateBarbershopInput{
				Name:            "barberShop1",
				Latitude:        1,
				Longitude:       1,
				EmployeesNumber: 2,
			},
			status: http.StatusBadRequest,
		},
		{
			name:   "Invalid input",
			token:  s.params["adminToken"],
			status: http.StatusBadRequest,
		},
		{
			name:  "Correctly Created",
			token: s.params["adminToken"],
			newBarberShop: &controller.CreateBarbershopInput{
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

			newBarberShopJson, _ := json.Marshal(tc.newBarberShop)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/admin/barber_shop/", bytes.NewBuffer(newBarberShopJson))
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
func (s *IntegrationSuite) TestBarberShopModifyByID() {
	testCases := []struct {
		name           string
		token          string
		ID             string
		editBarberShop *controller.ModifyBarberShopInput
		status         int
	}{
		{
			name:           "Require login",
			editBarberShop: &controller.ModifyBarberShopInput{Employees: 2},
			ID:             "genericID",
			status:         http.StatusUnauthorized,
		},
		{
			name:           "Require barber",
			token:          s.params["authToken"],
			editBarberShop: &controller.ModifyBarberShopInput{Employees: 3},
			status:         http.StatusUnauthorized,
			ID:             "genericID",
		},
		{
			name:           "Require to be a barber in that shop",
			token:          s.params["barber1Auth"],
			editBarberShop: &controller.ModifyBarberShopInput{Employees: 1},
			status:         http.StatusUnauthorized,
			ID:             s.params["barberShop2ID"],
		},
		{
			name:           "Correctly Modified",
			token:          s.params["barber1Auth"],
			editBarberShop: &controller.ModifyBarberShopInput{Employees: 2},
			status:         http.StatusAccepted,
			ID:             s.params["barberShop1ID"],
		},

		// TODO: test more bulky editing
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			editBarberShopJson, _ := json.Marshal(tc.editBarberShop)

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
			ID:     s.params["barberShop2ID"],
		},
		{
			name:   "Require admin",
			token:  s.params["authToken"],
			ID:     s.params["barberShop2ID"],
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly Eliminated",
			token:  s.params["adminToken"],
			ID:     s.params["barberShop2ID"],
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
