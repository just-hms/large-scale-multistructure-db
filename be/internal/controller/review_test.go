package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

func (s *ControllerSuite) TestStore() {

	testCases := []struct {
		name   string
		token  string
		ID     string
		status int
		input  controller.StoreReviewInput
	}{
		{
			name:   "Require Login",
			status: http.StatusUnauthorized,
			ID:     s.fixture[SHOP2_ID],
			input: controller.StoreReviewInput{
				Rating:  4,
				Content: "test",
			},
		},
		{
			name:   "Cannot review: the shop does not exists",
			token:  s.fixture[USER2_TOKEN],
			ID:     "fake_shop",
			status: http.StatusBadRequest,
			input: controller.StoreReviewInput{
				Rating:  4,
				Content: "test",
			},
		},
		{
			name:   "Correctly reviewed",
			token:  s.fixture[USER2_TOKEN],
			ID:     s.fixture[SHOP1_ID],
			status: http.StatusCreated,
			input: controller.StoreReviewInput{
				Rating:  4,
				Content: "test",
			},
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			reviewJson, _ := json.Marshal(tc.input)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/barbershop/"+tc.ID+"/review", bytes.NewBuffer(reviewJson))
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

func (s *ControllerSuite) TestShow() {

	testCases := []struct {
		name        string
		token       string
		ID          string
		status      int
		input       controller.StoreReviewInput
		resultCount int
	}{
		{
			name:        "Require Login",
			status:      http.StatusUnauthorized,
			ID:          s.fixture[SHOP2_ID],
			resultCount: 0,
		},
		{
			name:        "Cannot review: the shop does not exists",
			token:       s.fixture[USER2_TOKEN],
			ID:          "fake_shop",
			status:      http.StatusBadRequest,
			resultCount: 0,
		},
		{
			name:        "No reviews for the specified shop",
			token:       s.fixture[USER1_TOKEN],
			ID:          s.fixture[SHOP2_ID],
			status:      http.StatusOK,
			resultCount: 0,
		},
		{
			name:        "2 reviews for the specified shop",
			token:       s.fixture[USER1_TOKEN],
			ID:          s.fixture[SHOP1_ID],
			status:      http.StatusOK,
			resultCount: 2,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			// create a request for the register endpoint
			req, _ := http.NewRequest("GET", "/api/barbershop/"+tc.ID+"/review", nil)
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
					Reviews []entity.Review `json:"reviews"`
				}

				var res response

				err = json.Unmarshal(body, &res)
				s.Require().Nil(err)

				// assert that the number of returned user is as expected
				s.Require().Equal(tc.resultCount, len(res.Reviews))
			}

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)
		})

	}
}

func (s *ControllerSuite) TestDelete() {

	testCases := []struct {
		name     string
		token    string
		shopId   string
		reviewId string
		status   int
	}{
		{
			name:     "Require Login",
			shopId:   s.fixture[SHOP1_ID],
			reviewId: s.fixture[USER1_SHOP1_REVIEW1_ID],
			status:   http.StatusUnauthorized,
		},
		{
			name:     "Require Barber",
			token:    s.fixture[USER1_TOKEN],
			shopId:   s.fixture[SHOP1_ID],
			reviewId: s.fixture[USER1_SHOP1_REVIEW1_ID],
			status:   http.StatusUnauthorized,
		},
		{
			name:     "Correctly deleted",
			token:    s.fixture[BARBER1_TOKEN],
			shopId:   s.fixture[SHOP1_ID],
			reviewId: s.fixture[USER1_SHOP1_REVIEW1_ID],
			status:   http.StatusAccepted,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			// create a request for the register endpoint
			req, _ := http.NewRequest("DELETE", "/api/barbershop/"+tc.shopId+"/review/"+tc.reviewId, nil)
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

func (s *ControllerSuite) TestVote() {

	testCases := []struct {
		name     string
		token    string
		shopId   string
		reviewId string
		status   int
		input    controller.ReviewVoteInput
	}{
		{
			name:     "Require Login",
			shopId:   s.fixture[SHOP1_ID],
			reviewId: s.fixture[USER1_SHOP1_REVIEW1_ID],
			status:   http.StatusUnauthorized,
		},
		{
			name:     "Cannot review: the shop does not exists",
			token:    s.fixture[USER1_TOKEN],
			shopId:   "fake_shop",
			reviewId: s.fixture[USER1_SHOP1_REVIEW1_ID],
			status:   http.StatusBadRequest,
			input: controller.ReviewVoteInput{
				Upvote: true,
			},
		},
		{
			name:     "Correctly upvoted",
			token:    s.fixture[USER1_TOKEN],
			shopId:   s.fixture[SHOP1_ID],
			reviewId: s.fixture[USER1_SHOP1_REVIEW1_ID],
			status:   http.StatusCreated,
			input: controller.ReviewVoteInput{
				Upvote: true,
			},
		},
		{
			name:     "Cannot upvote twice",
			token:    s.fixture[USER1_TOKEN],
			shopId:   s.fixture[SHOP1_ID],
			reviewId: s.fixture[USER1_SHOP1_REVIEW1_ID],
			status:   http.StatusBadRequest,
			input: controller.ReviewVoteInput{
				Upvote: true,
			},
		},
		{
			name:     "Correctly downvoted",
			token:    s.fixture[USER1_TOKEN],
			shopId:   s.fixture[SHOP1_ID],
			reviewId: s.fixture[USER1_SHOP1_REVIEW1_ID],
			status:   http.StatusCreated,
			input: controller.ReviewVoteInput{
				Upvote: false,
			},
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			reviewJson, _ := json.Marshal(tc.input)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/barbershop/"+tc.shopId+"/review/"+tc.reviewId+"/vote", bytes.NewBuffer(reviewJson))
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

func (s *ControllerSuite) TestDeleteVote() {

	testCases := []struct {
		name     string
		token    string
		shopId   string
		reviewId string
		status   int
	}{
		{
			name:     "Require Login",
			shopId:   s.fixture[SHOP1_ID],
			reviewId: s.fixture[USER1_SHOP1_REVIEW1_ID],
			status:   http.StatusUnauthorized,
		},
		{
			name:     "Review does not exist",
			token:    s.fixture[BARBER1_TOKEN],
			shopId:   s.fixture[SHOP1_ID],
			reviewId: "fake_review",
			status:   http.StatusBadRequest,
		},
		{
			name:     "Correctly deleted",
			token:    s.fixture[BARBER1_TOKEN],
			shopId:   s.fixture[SHOP1_ID],
			reviewId: s.fixture[USER1_SHOP1_REVIEW1_ID],
			status:   http.StatusAccepted,
		},
	}

	for _, tc := range testCases {

		s.Run(tc.name, func() {

			// create a request for the register endpoint
			req, _ := http.NewRequest("DELETE", "/api/barbershop/"+tc.shopId+"/review/"+tc.reviewId, nil)
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
