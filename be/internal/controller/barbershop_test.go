package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

func (s *ControllerSuite) TestBarberShopFind() {

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
			token:       s.fixture[USER1_TOKEN],
			resultCount: 3,
		},
		{
			name: "No barbershop near where you are",
			input: controller.FindBarbershopInput{
				Latitude:  1.1,
				Longitude: 1.1,
				Radius:    1,
			},
			token:       s.fixture[USER1_TOKEN],
			status:      http.StatusOK,
			resultCount: 0,
		},
		{
			name: "No barbershop for this name",
			input: controller.FindBarbershopInput{
				Name: "not_existing_shop",
			},
			token:       s.fixture[USER1_TOKEN],
			status:      http.StatusOK,
			resultCount: 0,
		},
		{
			name: "Two barbershop found for this name",
			input: controller.FindBarbershopInput{
				Name: "boh",
			},
			token:       s.fixture[USER1_TOKEN],
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
			token:       s.fixture[USER1_TOKEN],
			status:      http.StatusOK,
			resultCount: 1,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			newBarberShopJson, _ := json.Marshal(tc.input)

			// create a request for the self endpoint
			req, _ := http.NewRequest("POST", "/api/barbershop", bytes.NewBuffer(newBarberShopJson))
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

// TODO:
// - check that the appointments are returned

func (s *ControllerSuite) TestBarberShopShow() {
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
			token:  s.fixture[USER1_TOKEN],
			ID:     s.fixture[SHOP1_ID],
			status: http.StatusOK,
		},
		{
			name:   "BarberShop doesn't exist",
			token:  s.fixture[USER1_TOKEN],
			ID:     "wrong_ID",
			status: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			// create a request for the self endpoint
			var req *http.Request

			req, _ = http.NewRequest("GET", "/api/barbershop/"+tc.ID, nil)
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
func (s *ControllerSuite) TestBarberShopStore() {

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
				Description:     "kek",
			},
			status: http.StatusUnauthorized,
		},
		{
			name:  "Require admin",
			token: s.fixture[USER1_ID],
			input: &controller.CreateBarbershopInput{
				Name:            "barberShop7",
				Latitude:        1,
				Longitude:       1,
				EmployeesNumber: 2,
				Description:     "kek",
			},
			status: http.StatusUnauthorized,
		},
		{
			name:  "Already exists",
			token: s.fixture[ADMIN_TOKEN],
			input: &controller.CreateBarbershopInput{
				Name:            "barberShop1",
				Latitude:        1,
				Longitude:       1,
				EmployeesNumber: 2,
				Description:     "kek",
			},
			status: http.StatusBadRequest,
		},
		{
			name:   "Invalid input",
			token:  s.fixture[ADMIN_TOKEN],
			status: http.StatusBadRequest,
		},
		{
			name:  "Correctly Created",
			token: s.fixture[ADMIN_TOKEN],
			input: &controller.CreateBarbershopInput{
				Name:            "barberShop7",
				Latitude:        1.1,
				Longitude:       1,
				EmployeesNumber: 2,
				Description:     "kek",
			},
			status: http.StatusCreated,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			newBarberShopJson, _ := json.Marshal(tc.input)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/admin/barbershop", bytes.NewBuffer(newBarberShopJson))
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
func (s *ControllerSuite) TestBarberShopModifyByID() {
	testCases := []struct {
		name           string
		token          string
		ID             string
		input          *controller.ModifyBarberShopInput
		shouldTestSlot bool
		status         int
	}{
		{
			name:           "Require login",
			input:          &controller.ModifyBarberShopInput{Employees: 2},
			ID:             "genericID",
			shouldTestSlot: false,
			status:         http.StatusUnauthorized,
		},
		{
			name:           "Require barber",
			token:          s.fixture[USER1_ID],
			input:          &controller.ModifyBarberShopInput{Employees: 3},
			status:         http.StatusUnauthorized,
			shouldTestSlot: false,
			ID:             "genericID",
		},
		{
			name:           "Require to be a barber in that shop",
			token:          s.fixture[BARBER1_TOKEN],
			input:          &controller.ModifyBarberShopInput{Employees: 1},
			status:         http.StatusUnauthorized,
			shouldTestSlot: false,
			ID:             s.fixture[SHOP2_ID],
		},
		{
			name:           "Correctly Modified",
			token:          s.fixture[BARBER1_TOKEN],
			input:          &controller.ModifyBarberShopInput{Employees: 2},
			status:         http.StatusAccepted,
			shouldTestSlot: false,
			ID:             s.fixture[SHOP1_ID],
		},
		{
			name:           "Correctly Modified Slot",
			token:          s.fixture[BARBER1_TOKEN],
			input:          &controller.ModifyBarberShopInput{Employees: 5},
			status:         http.StatusAccepted,
			shouldTestSlot: true,
			ID:             s.fixture[SHOP1_ID],
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			// Check that Slot Employees modifying also works
			if tc.shouldTestSlot {

				app := controller.BookAppointmentInput{
					DateTime: time.Now().Add(time.Hour),
				}

				bookingJson, _ := json.Marshal(app)

				// create a request for the register endpoint
				req, _ := http.NewRequest("POST", "/api/barbershop/"+tc.ID+"/appointment", bytes.NewBuffer(bookingJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Add("Authorization", "Bearer "+s.fixture[USER2_TOKEN])

				// serve the request to the test server
				w := httptest.NewRecorder()
				s.srv.ServeHTTP(w, req)

				s.Require().Equal(http.StatusCreated, w.Code)
			}

			editBarberShopJson, _ := json.Marshal(tc.input)

			// create a request for the register endpoint
			req, _ := http.NewRequest("PUT", "/api/barbershop/"+tc.ID, bytes.NewBuffer(editBarberShopJson))
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
func (s *ControllerSuite) TestBarberShopDeleteByID() {

	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
	}{
		{
			name:   "Require Login",
			status: http.StatusUnauthorized,
			ID:     s.fixture[SHOP2_ID],
		},
		{
			name:   "Require admin",
			token:  s.fixture[USER1_TOKEN],
			ID:     s.fixture[SHOP2_ID],
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly Eliminated",
			token:  s.fixture[ADMIN_TOKEN],
			ID:     s.fixture[SHOP2_ID],
			status: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			// create a request for the register endpoint
			req, _ := http.NewRequest("DELETE", "/api/admin/barbershop/"+tc.ID, nil)

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
