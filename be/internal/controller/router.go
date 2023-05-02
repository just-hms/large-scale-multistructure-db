package controller

import (
	"net/http"

	_ "github.com/just-hms/large-scale-multistructure-db/be/apidocs"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller/middleware"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func CORSAllowAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7000
// @BasePath /api/
// @schemes http
func Router(ucs map[byte]usecase.Usecase, production bool) *gin.Engine {

	if production {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(CORSAllowAll())

	api := router.Group("/api")

	url := ginSwagger.URL("/api/swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api.GET("/health", Health)

	// create the routes based on the given usecases
	mr := middleware.NewMiddlewareRoutes(
		ucs[usecase.USER].(usecase.User),
		ucs[usecase.TOKEN].(usecase.Token),
	)
	ur := NewUserRoutes(
		ucs[usecase.USER].(usecase.User),
		ucs[usecase.TOKEN].(usecase.Token),
	)

	br := NewBarberShopRoutes(
		ucs[usecase.BARBER_SHOP].(usecase.BarberShop),
		ucs[usecase.CALENDAR].(usecase.Calendar),
		ucs[usecase.TOKEN].(usecase.Token),
	)

	ar := NewAppointmentRoutes(
		ucs[usecase.APPOINTMENT].(usecase.Appointment),
		ucs[usecase.USER].(usecase.User),
		ucs[usecase.TOKEN].(usecase.Token),
	)

	hr := NewHolidayRoutes(
		ucs[usecase.HOLIDAY].(usecase.Holiday),
	)

	gr := NewGeocodingRoutes(
		ucs[usecase.GEOCODING].(usecase.Geocoding),
	)

	api.POST("geocoding/search", mr.RequireAuth, gr.Search)

	user := api.Group("/user")
	{
		user.POST("", ur.Register)
		user.POST("/login", ur.Login)
		user.GET("/self", mr.RequireAuth, mr.MarkWithAuthID, ur.Show)
		user.DELETE("/self", mr.RequireAuth, mr.MarkWithAuthID, ur.Delete)
		user.POST("/lostpassword", ur.LostPassword)
		user.POST("/resetpassword/:resettoken", ur.ResetPassword)

		user.DELETE("/self/appointment", mr.RequireAuth, ar.DeleteSelfAppointment)
	}

	barberShop := api.Group("/barbershop")
	barberShop.Use(mr.RequireAuth)
	{
		barberShop.POST("", br.Find)
		barberShop.GET("/:shopid", br.Show)

		barberShop.PUT("/:shopid", mr.RequireBarber, br.Modify)

		barberShop.DELETE("/:shopid/appointment/:appointmentid", mr.RequireBarber, ar.DeleteAppointment)
		barberShop.POST("/:shopid/appointment", ar.Book)

		barberShop.GET("/:shopid/calendar", br.Calendar)
		barberShop.POST("/:shopid/holiday", mr.RequireBarber, hr.Set)
	}

	admin := api.Group("/admin")
	admin.Use(mr.RequireAdmin)
	{
		admin.GET("/user", ur.ShowAll)
		admin.GET("/user/:id", ur.Show)
		admin.DELETE("/user/:id", ur.Delete)
		admin.PUT("/user/:id", ur.Modify)

		admin.POST("/barbershop", br.Create)
		admin.DELETE("/barbershop/:shopid", br.Delete)
	}

	return router
}

// HealthCheck Show the status of the server
//
// @Summary Show the status of the server
// @Description Get the status of the server
// @Tags Root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
