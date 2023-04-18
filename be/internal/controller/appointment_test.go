package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
)

func (s *ControllerSuite) TestBook() {

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
			ID:     s.fixture[SHOP2_ID],
			input: controller.BookAppointmentInput{
				DateTime: time.Now().Add(time.Hour),
			},
		},
		{
			name:   "Correctly booked",
			token:  s.fixture[USER2_TOKEN],
			ID:     s.fixture[BARBER1_ID],
			status: http.StatusCreated,
			input: controller.BookAppointmentInput{
				DateTime: time.Now().Add(time.Hour),
			},
		},
		{
			name:   "Cannot book two appontments",
			token:  s.fixture[USER1_TOKEN],
			ID:     s.fixture[BARBER1_ID],
			status: http.StatusBadRequest,
			input: controller.BookAppointmentInput{
				DateTime: time.Now().Add(time.Hour),
			},
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			bookingJson, _ := json.Marshal(tc.input)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/barber_shop/"+tc.ID+"/appointment", bytes.NewBuffer(bookingJson))
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

func (s *ControllerSuite) TestCancelSelfAppointment() {

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
			token:  s.fixture[USER1_TOKEN],
			status: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

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

func (s *ControllerSuite) TestCancelAppointment() {

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
			ID:      s.fixture[USER1_SHOP1_APPOINTMENT_ID],
			SHOP_ID: s.fixture[SHOP1_ID],
		},
		{
			name:    "Require barber",
			token:   s.fixture[USER1_TOKEN],
			status:  http.StatusUnauthorized,
			ID:      s.fixture[USER1_SHOP1_APPOINTMENT_ID],
			SHOP_ID: s.fixture[SHOP1_ID],
		},
		{
			name:    "Correctly deleted",
			token:   s.fixture[BARBER1_TOKEN],
			status:  http.StatusAccepted,
			ID:      s.fixture[USER1_SHOP1_APPOINTMENT_ID],
			SHOP_ID: s.fixture[SHOP1_ID],
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

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

func (s *ControllerSuite) TestBookAfterCancel() {

	// delete an appointment from USER1
	req, _ := http.NewRequest("DELETE", "/api/user/self/appointment", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.fixture[USER1_TOKEN])

	// serve the request to the test server
	w := httptest.NewRecorder()
	s.srv.ServeHTTP(w, req)

	body, err := io.ReadAll(w.Body)
	s.Require().NoError(err)
	fmt.Println(string(body))

	s.Require().Equal(http.StatusAccepted, w.Code)

	// test if now he can book

	BookingJson, err := json.Marshal(controller.BookAppointmentInput{
		DateTime: time.Now().Add(time.Hour),
	})
	s.Require().NoError(err)

	// create a request for the register endpoint
	req, err = http.NewRequest(
		"POST", "/api/barber_shop/"+s.fixture[SHOP2_ID]+"/appointment",
		bytes.NewBuffer(BookingJson),
	)
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.fixture[USER1_TOKEN])

	// serve the request to the test server
	w = httptest.NewRecorder()
	s.srv.ServeHTTP(w, req)

	// assert that the response status code is as expected
	s.Require().Equal(http.StatusCreated, w.Code)
}
