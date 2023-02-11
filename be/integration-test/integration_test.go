package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"large-scale-multistructure-db/be/internal/controller"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/internal/usecase/auth"
	"large-scale-multistructure-db/be/internal/usecase/repo"
	"large-scale-multistructure-db/be/pkg/jwt"
	"large-scale-multistructure-db/be/pkg/mongo"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
	srv   *gin.Engine
	mongo *mongo.Mongo

	resetDB func()
	params  map[string]string
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) SetupTest() {
	s.resetDB()
}

func (s *IntegrationSuite) SetupSuite() {

	fmt.Println(">>> From SetupSuite")

	mongo, err := mongo.New(&mongo.Options{
		DB_NAME: "test",
	})

	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}

	// create repos and usecases
	// TODO add barbershop ones
	userRepo := repo.NewUserRepo(mongo)
	barberShopRepo := repo.NewBarberShopRepo(mongo)
	viewShopRepo := repo.NewShopViewRepo(mongo)

	password := auth.NewPasswordAuth()

	usecases := []usecase.Usecase{
		usecase.NewUserUseCase(
			userRepo,
			password,
		),
		usecase.NewBarberShopUseCase(
			barberShopRepo,
			viewShopRepo,
		),
	}

	// Fill the test DB

	s.resetDB = func() {

		mongo.DB.Drop(context.TODO())

		p, _ := password.HashAndSalt("password")

		// create users
		us := &entity.User{
			Email:    "correct@example.com",
			Password: p,
			IsAdmin:  false,
		}

		userRepo.Store(context.TODO(), us)

		s.params["authID"] = us.ID
		s.params["authToken"], _ = jwt.CreateToken(us.ID)

		admin := &entity.User{
			Email:    "admin@example.com",
			Password: p,
			IsAdmin:  true,
		}

		userRepo.Store(context.TODO(), admin)

		s.params["adminID"] = admin.ID
		s.params["adminToken"], _ = jwt.CreateToken(admin.ID)

		userRepo.Store(context.TODO(), &entity.User{
			Email:    "to.filter@example.com",
			Password: p,
			IsAdmin:  true,
		})

		// create barberShops

		barberShop1 := &entity.BarberShop{
			Name:            "barberShop1",
			Latitude:        "1",
			Longitude:       "1",
			EmployeesNumber: 2,
		}
		barberShopRepo.Store(context.TODO(), barberShop1)

		barber1 := &entity.User{
			Email:    "barber1@example.com",
			Password: p,
			IsAdmin:  false,
			BarberShopIDs: []string{
				barberShop1.ID,
			},
		}

		userRepo.Store(context.TODO(), barber1)

		s.params["bnonarberShop1ID"] = barberShop1.ID
		s.params["barber1Auth"], _ = jwt.CreateToken(barber1.ID)

		barberShop2 := &entity.BarberShop{
			Name:            "barberShop1",
			Latitude:        "1",
			Longitude:       "1",
			EmployeesNumber: 2,
		}

		barberShopRepo.Store(context.TODO(), barberShop2)

		barber2 := &entity.User{
			Email:    "barber2@example.com",
			Password: p,
			IsAdmin:  false,
			BarberShopIDs: []string{
				barberShop2.ID,
			},
		}

		userRepo.Store(context.TODO(), barber2)

		s.params["barberShop2ID"] = barber2.ID
		s.params["barber2Auth"], _ = jwt.CreateToken(barber2.ID)

	}

	// serv the mock server and db
	s.params = make(map[string]string)
	s.srv = controller.Router(usecases)
	s.mongo = mongo
}

func (s *IntegrationSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

func (s *IntegrationSuite) TestHealth() {

	req, _ := http.NewRequest("GET", "/api/health/", nil)

	w := httptest.NewRecorder()
	s.srv.ServeHTTP(w, req)

	s.Require().Equal(http.StatusOK, w.Code)
}

// USERS

func (s *IntegrationSuite) TestLogin() {

	testCases := []struct {
		name      string
		loginUser *entity.User
		status    int
	}{
		{
			name:      "Correct Login",
			loginUser: &entity.User{Email: "correct@example.com", Password: "password"},
			status:    http.StatusOK,
		},
		{
			name:      "Invalid input",
			loginUser: &entity.User{Email: "not_an_email", Password: "password"},
			status:    http.StatusBadRequest,
		},
		{
			name:      "Wrong password",
			loginUser: &entity.User{Email: "correct@example.com", Password: "invalid"},
			status:    http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			loginUserJson, _ := json.Marshal(tc.loginUser)

			// create a request for the login endpoint
			req, _ := http.NewRequest("POST", "/api/user/login/", bytes.NewBuffer(loginUserJson))
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

func (s *IntegrationSuite) TestRegister() {

	testCases := []struct {
		name         string
		creationUser *entity.User
		status       int
	}{
		{
			name:         "Already exists",
			creationUser: &entity.User{Email: "correct@example.com", Password: "password"},
			status:       http.StatusUnauthorized,
		},
		{
			name:         "Invalid input",
			creationUser: &entity.User{Email: "not_an_email", Password: "password"},
			status:       http.StatusBadRequest,
		},
		{
			name:         "Correctly Created",
			creationUser: &entity.User{Email: "new@example.com", Password: "password"},
			status:       http.StatusCreated,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {
			loginUserJson, _ := json.Marshal(tc.creationUser)

			// create a request for the register endpoint
			req, _ := http.NewRequest("POST", "/api/user/", bytes.NewBuffer(loginUserJson))
			req.Header.Set("Content-Type", "application/json")

			// serve the request to the test server
			w := httptest.NewRecorder()
			s.srv.ServeHTTP(w, req)

			// assert that the response status code is as expected
			s.Require().Equal(tc.status, w.Code)
		})
	}
}

func (s *IntegrationSuite) TestShowSelf() {

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
			token:  s.params["authToken"],
			status: http.StatusOK,
		},
		{
			name:           "Show barber",
			token:          s.params["barber1Auth"],
			status:         http.StatusOK,
			barberShopsLen: 1,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the self endpoint
			req, _ := http.NewRequest("GET", "/api/user/self/", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+tc.token)

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

				s.Require().Len(res.User.BarberShopIDs, tc.barberShopsLen)

			}
		})

	}
}

func (s *IntegrationSuite) TestDeleteSelf() {

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
			token:  s.params["authToken"],
			status: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

			// create a request for the self endpoint
			req, _ := http.NewRequest("DELETE", "/api/user/self/", nil)
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

func (s *IntegrationSuite) TestUserShowAll() {

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
			token:  s.params["authToken"],
			status: http.StatusUnauthorized,
		},
		{
			name:        "Correctly shown without filter",
			token:       s.params["adminToken"],
			status:      http.StatusOK,
			resultCount: 3,
		},
		{
			name:        "Correctly shown with filter",
			filter:      "filter",
			token:       s.params["adminToken"],
			status:      http.StatusOK,
			resultCount: 1,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

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

func (s *IntegrationSuite) TestUserShow() {

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
			token:  s.params["authToken"],
			ID:     "genericID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly shown user",
			token:  s.params["adminToken"],
			ID:     s.params["authID"],
			status: http.StatusOK,
		},
		{
			name:   "User not exists",
			token:  s.params["adminToken"],
			ID:     "wrong_ID",
			status: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

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

func (s *IntegrationSuite) TestUserDelete() {

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
			token:  s.params["authToken"],
			ID:     "genericID",
			status: http.StatusUnauthorized,
		},
		{
			name:   "Correctly deleted user",
			token:  s.params["adminToken"],
			ID:     s.params["authID"],
			status: http.StatusAccepted,
		},
		{
			name:   "User not exists",
			token:  s.params["adminToken"],
			ID:     "wrong_ID",
			status: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {

		s.T().Run(tc.name, func(t *testing.T) {

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

// BARBER SHOPS

// func (s *IntegrationSuite) TestBarberShopFind(t *testing.T) {

// 	testCases := []struct {
// 		name           string
// 		lat            string
// 		lon            string
// 		radius         string
// 		token          string
// 		status         int
// 		barberShopName string
// 		resultCount    int
// 		res            interface{}
// 	}{
// 		{
// 			name:   "Unauthorized",
// 			status: http.StatusUnauthorized,
// 		},
// 		{
// 			name:        "All barbershops",
// 			status:      http.StatusOK,
// 			token:       s.params["authToken"],
// 			resultCount: 3,
// 		},
// 		{
// 			name:        "No barbershop near where you are",
// 			lat:         "1.1",
// 			lon:         "1.1",
// 			radius:      "1",
// 			token:       s.params["authToken"],
// 			status:      http.StatusOK,
// 			resultCount: 0,
// 		},
// 		{
// 			name:           "No barbershop for this name",
// 			barberShopName: "not_existing_shop",
// 			token:          s.params["authToken"],
// 			status:         http.StatusOK,
// 			resultCount:    0,
// 		},
// 		{
// 			name:           "Two barbershop found for this name",
// 			barberShopName: "not_existing_shop",
// 			token:          s.params["authToken"],
// 			status:         http.StatusOK,
// 			resultCount:    2,
// 		},
// 		{
// 			name:        "Found 1 barbershop near you",
// 			lat:         "1.1",
// 			lon:         "1.1",
// 			radius:      "100",
// 			token:       s.params["authToken"],
// 			status:      http.StatusOK,
// 			resultCount: 1,
// 		},
// 	}

// 	for _, tc := range testCases {

// 		s.T().Run(tc.name, func(t *testing.T) {

// 			// create a request for the self endpoint
// 			var req *http.Request

// 			query := url.Values{
// 				"name":   {tc.name},
// 				"lat":    {tc.lat},
// 				"lon":    {tc.lon},
// 				"radius": {tc.radius},
// 			}

// 			req, _ = http.NewRequest("GET", "/api/barber_shop"+query.Encode(), nil)
// 			req.Header.Set("Content-Type", "application/json")
// 			req.Header.Add("Authorization", "Bearer "+tc.token)

// 			// serve the request to the test server
// 			w := httptest.NewRecorder()
// 			s.srv.ServeHTTP(w, req)

// 			if w.Code == http.StatusAccepted {

// 				body, err := io.ReadAll(w.Body)

// 				// require no error in reading the response
// 				s.Require().Nil(err)

// 				type response struct {
// 					BarberShops []entity.User `json:"barberShops"`
// 				}

// 				var res response

// 				err = json.Unmarshal(body, &res)
// 				s.Require().Nil(err)

// 				// assert that the number of returned user is as expected
// 				s.Require().Equal(tc.resultCount, len(res.BarberShops))
// 			}

// 			// assert that the response status code is as expected
// 			s.Require().Equal(tc.status, w.Code)
// 		})

// 	}
// }

// // TODO : check if the viewShop is stored
// func (s *IntegrationSuite) TestBarberShopShow() {
// 	testCases := []struct {
// 		name   string
// 		token  string
// 		ID     string
// 		status int
// 	}{
// 		{
// 			name:   "Wrongly formatted token",
// 			token:  "wrong_token",
// 			status: http.StatusUnauthorized,
// 		},
// 		{
// 			name:   "Correctly shown user",
// 			token:  s.params["authToken"],
// 			ID:     s.params["authID"], // TODO fix this with a barberShopID
// 			status: http.StatusOK,
// 		},
// 		{
// 			name:   "BarberShop doesn't exist",
// 			token:  s.params["authToken"],
// 			ID:     "wrong_ID",
// 			status: http.StatusNotFound,
// 		},
// 	}

// 	for _, tc := range testCases {

// 		s.T().Run(tc.name, func(t *testing.T) {

// 			// create a request for the self endpoint
// 			var req *http.Request

// 			req, _ = http.NewRequest("GET", "/api/barber_shop/"+tc.ID, nil)
// 			req.Header.Set("Content-Type", "application/json")
// 			req.Header.Add("Authorization", "Bearer "+tc.token)

// 			// serve the request to the test server
// 			w := httptest.NewRecorder()
// 			s.srv.ServeHTTP(w, req)

// 			// assert that the response status code is as expected
// 			s.Require().Equal(tc.status, w.Code)

// 			// check for user len if the request was accepted
// 			if w.Code == http.StatusAccepted {

// 				body, err := io.ReadAll(w.Body)

// 				// require no error in reading the response
// 				s.Require().Nil(err)

// 				type response struct {
// 					Barber entity.BarberShop `json:"barber"`
// 				}

// 				var res response
// 				err = json.Unmarshal(body, &res)
// 				s.Require().Nil(err)
// 			}
// 		})
// 	}
// }
// func (s *IntegrationSuite) TestBarberShopStore() {

// 	testCases := []struct {
// 		name          string
// 		token         string
// 		newBarberShop *entity.BarberShop
// 		status        int
// 	}{
// 		{
// 			name:          "Require login",
// 			newBarberShop: &entity.BarberShop{Name: "barberShop7"},
// 			status:        http.StatusUnauthorized,
// 		},
// 		{
// 			name:          "Require admin",
// 			token:         s.params["authToken"],
// 			newBarberShop: &entity.BarberShop{Name: "barberShop7"},
// 			status:        http.StatusUnauthorized,
// 		},
// 		{
// 			name:          "Already exists",
// 			token:         s.params["adminToken"],
// 			newBarberShop: &entity.BarberShop{Name: "barberShop1"},
// 			status:        http.StatusUnauthorized,
// 		},
// 		{
// 			name:   "Invalid input",
// 			token:  s.params["adminToken"],
// 			status: http.StatusBadRequest,
// 		},
// 		{
// 			name:          "Correctly Created",
// 			token:         s.params["adminToken"],
// 			newBarberShop: &entity.BarberShop{Name: "barberShop7"},
// 			status:        http.StatusCreated,
// 		},
// 	}

// 	for _, tc := range testCases {

// 		s.T().Run(tc.name, func(t *testing.T) {

// 			newBarberShopJson, _ := json.Marshal(tc.newBarberShop)

// 			// create a request for the register endpoint
// 			req, _ := http.NewRequest("POST", "/api/admin/barber_shop/", bytes.NewBuffer(newBarberShopJson))
// 			req.Header.Set("Content-Type", "application/json")

// 			// serve the request to the test server
// 			w := httptest.NewRecorder()
// 			s.srv.ServeHTTP(w, req)

// 			// assert that the response status code is as expected
// 			s.Require().Equal(tc.status, w.Code)
// 		})
// 	}
// }
// func (s *IntegrationSuite) TestBarberShopModifyByID() {
// 	testCases := []struct {
// 		name           string
// 		token          string
// 		ID             string
// 		editBarberShop *entity.BarberShop
// 		status         int
// 	}{
// 		{
// 			name:           "Require login",
// 			editBarberShop: &entity.BarberShop{EmployeesNumber: 2},
// 			ID:             "genericID",
// 			status:         http.StatusUnauthorized,
// 		},
// 		{
// 			name:           "Require barber",
// 			token:          s.params["authToken"],
// 			editBarberShop: &entity.BarberShop{EmployeesNumber: 3},
// 			status:         http.StatusUnauthorized,
// 			ID:             "genericID",
// 		},
// 		{
// 			name:           "Require to be a barber in that shop",
// 			token:          s.params["barber1Auth"],
// 			editBarberShop: &entity.BarberShop{Name: "barberShop1"},
// 			status:         http.StatusUnauthorized,
// 			ID:             s.params["barberShop2ID"],
// 		},
// 		{
// 			name:           "Correctly Modified",
// 			token:          s.params["barber1Auth"],
// 			editBarberShop: &entity.BarberShop{Name: "barberShop1"},
// 			status:         http.StatusAccepted,
// 			ID:             s.params["barberShop1ID"],
// 		},
// 	}

// 	for _, tc := range testCases {

// 		s.T().Run(tc.name, func(t *testing.T) {

// 			editBarberShopJson, _ := json.Marshal(tc.editBarberShop)

// 			// create a request for the register endpoint
// 			req, _ := http.NewRequest("PUT", "/api/barber_shop/"+tc.ID, bytes.NewBuffer(editBarberShopJson))
// 			req.Header.Set("Content-Type", "application/json")

// 			// serve the request to the test server
// 			w := httptest.NewRecorder()
// 			s.srv.ServeHTTP(w, req)

// 			// assert that the response status code is as expected
// 			s.Require().Equal(tc.status, w.Code)
// 		})

// 	}
// }
// func (s *IntegrationSuite) TestBarberShopDeleteByID() {

// 	testCases := []struct {
// 		name   string
// 		token  string
// 		ID     string
// 		status int
// 	}{
// 		{
// 			name:   "Require Login",
// 			status: http.StatusUnauthorized,
// 			ID:     s.params["barberShop2ID"],
// 		},
// 		{
// 			name:   "Require admin",
// 			token:  s.params["authToken"],
// 			ID:     s.params["barberShop2ID"],
// 			status: http.StatusUnauthorized,
// 		},
// 		{
// 			name:   "Correctly Eliminated",
// 			token:  s.params["adminToken"],
// 			ID:     s.params["barberShop2ID"],
// 			status: http.StatusAccepted,
// 		},
// 	}

// 	for _, tc := range testCases {

// 		s.T().Run(tc.name, func(t *testing.T) {

// 			// create a request for the register endpoint
// 			req, _ := http.NewRequest("DELETE", "/api/admin/barber_shop/"+tc.ID, nil)

// 			// serve the request to the test server
// 			w := httptest.NewRecorder()
// 			s.srv.ServeHTTP(w, req)

// 			// assert that the response status code is as expected
// 			s.Require().Equal(tc.status, w.Code)
// 		})

// 	}
// }
