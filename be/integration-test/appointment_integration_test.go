package integration_test

import (
	"bytes"
	"encoding/json"
	"large-scale-multistructure-db/be/internal/controller"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
			ID:     s.params["barberShop2ID"],
			input: controller.BookAppointmentInput{
				DateTime: time.Now().Add(time.Hour),
			},
		},
		{
			name:   "Correctly booked",
			token:  s.params["authToken"],
			ID:     s.params["barberShop2ID"],
			status: http.StatusCreated,
			input: controller.BookAppointmentInput{
				DateTime: time.Now().Add(time.Hour),
			},
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			BookingJson, _ := json.Marshal(tc.input)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/barber_shop/"+tc.ID+"/appointment/", bytes.NewBuffer(BookingJson))
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
			token:  s.params["auth2Token"],
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
		name   string
		token  string
		status int
		ID     string
	}{
		{
			name:   "Require Login",
			status: http.StatusUnauthorized,
			ID:     s.params["appointmentID"],
		},
		{
			name:   "Correctly deleted",
			token:  s.params["barber2Auth"],
			status: http.StatusAccepted,
			ID:     s.params["appointmentID"],
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the register endpoint
			req, _ := http.NewRequest("DELETE", "/barber_shop/"+tc.ID+"/appointment", nil)
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