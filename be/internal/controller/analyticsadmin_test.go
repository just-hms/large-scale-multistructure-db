package controller_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

func (s *ControllerSuite) TestGetAdminAnalytics() {
	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
	}{
		{
			name:   "Wrongly formatted token",
			token:  "wrong_token",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Not an admin",
			token:  s.fixture[BARBER1_TOKEN],
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly got Admin's Analytics",
			token:  s.fixture[ADMIN_TOKEN],
			status: http.StatusOK,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			var req *http.Request

			req, _ = http.NewRequest("GET", "/api/admin/analytics", nil)
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
					AdminAnalytics entity.AdminAnalytics `json:"adminAnalytics"`
				}

				var res response
				err = json.Unmarshal(body, &res)
				s.Require().Nil(err)
			}
		})
	}
}
