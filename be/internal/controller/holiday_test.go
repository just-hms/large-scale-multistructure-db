package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
)

func (s *ControllerSuite) TestHolidaySet() {

	fakeTime := time.Now().Add(1 * time.Hour)

	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
		input  *controller.SetHolidaysInput
	}{
		{
			name:   "Require Login",
			status: http.StatusUnauthorized,
			ID:     s.fixture[SHOP2_ID],
		},
		{
			name:   "Require barber",
			status: http.StatusUnauthorized,
			token:  s.fixture[USER1_TOKEN],
			ID:     s.fixture[SHOP2_ID],
		},
		{
			name:   "Cannot set an holiday in the past",
			status: http.StatusBadRequest,
			token:  s.fixture[BARBER2_TOKEN],
			ID:     s.fixture[SHOP2_ID],
			input: &controller.SetHolidaysInput{
				Date:                 time.Now().Add(-1 * time.Hour),
				UnavailableEmployees: 1,
			},
		},
		{
			name:   "Correctly set",
			token:  s.fixture[BARBER2_TOKEN],
			ID:     s.fixture[SHOP2_ID],
			status: http.StatusAccepted,
			input: &controller.SetHolidaysInput{
				Date:                 fakeTime,
				UnavailableEmployees: 2,
			},
		},
		{
			name:   "Correctly unset",
			token:  s.fixture[BARBER2_TOKEN],
			ID:     s.fixture[SHOP2_ID],
			status: http.StatusAccepted,
			input: &controller.SetHolidaysInput{
				Date:                 fakeTime,
				UnavailableEmployees: 2,
			},
		},
		{
			name:   "Cannot set an holiday if overbooked",
			status: http.StatusBadRequest,
			token:  s.fixture[BARBER2_TOKEN],
			ID:     s.fixture[EMPTY_SHOP],
			input: &controller.SetHolidaysInput{
				Date:                 fakeTime,
				UnavailableEmployees: 1,
			},
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			holidayJSON, _ := json.Marshal(tc.input)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/barber_shop/"+tc.ID+"/holiday", bytes.NewBuffer(holidayJSON))
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
