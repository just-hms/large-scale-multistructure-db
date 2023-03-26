package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
)

func (s *IntegrationSuite) TestBook() {

	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
		input  controller.BookAppointmentInput
	}{
		{
			name:   "Require Login",
			status: http.StatusUnauthorized,
			ID:     s.params[SHOP2_ID],
			input: controller.BookAppointmentInput{
				DateTime: time.Now().Add(time.Hour),
			},
		},
		{
			name:   "Correctly booked",
			token:  s.params[USER2_TOKEN],
			ID:     s.params[BARBER1_ID],
			status: http.StatusCreated,
			input: controller.BookAppointmentInput{
				DateTime: time.Now().Add(time.Hour),
			},
		},
		{
			name:   "Cannot book two appontments",
			token:  s.params[USER1_TOKEN],
			ID:     s.params[BARBER1_ID],
			status: http.StatusUnauthorized,
			input: controller.BookAppointmentInput{
				DateTime: time.Now().Add(time.Hour),
			},
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			BookingJson, _ := json.Marshal(tc.input)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/barber_shop/"+tc.ID+"/appointment", bytes.NewBuffer(BookingJson))
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

func (s *IntegrationSuite) TestCancelSelfAppointment() {

	testCases := []struct {
		name   string
		token  string
		status int
	}{
		{
			name:   "Require Login",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly deleted",
			token:  s.params[USER1_TOKEN],
			status: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the register endpoint
			req, _ := http.NewRequest("DELETE", "/api/user/self/appointment", nil)
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

func (s *IntegrationSuite) TestCancelAppointment() {

	testCases := []struct {
		name    string
		token   string
		status  int
		ID      string
		SHOP_ID string
	}{
		{
			name:    "Require Login",
			status:  http.StatusUnauthorized,
			ID:      s.params[USER1_SHOP1_APPOINTMENT],
			SHOP_ID: s.params[SHOP1_ID],
		},
		{
			name:    "Require barber",
			token:   s.params[USER1_TOKEN],
			status:  http.StatusUnauthorized,
			ID:      s.params[USER1_SHOP1_APPOINTMENT],
			SHOP_ID: s.params[SHOP1_ID],
		},
		{
			name:    "Correctly deleted",
			token:   s.params[BARBER1_TOKEN],
			status:  http.StatusAccepted,
			ID:      s.params[USER1_SHOP1_APPOINTMENT],
			SHOP_ID: s.params[SHOP1_ID],
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the register endpoint
			req, _ := http.NewRequest("DELETE", "/api/barber_shop/"+tc.SHOP_ID+"/appointment/"+tc.ID, nil)
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
