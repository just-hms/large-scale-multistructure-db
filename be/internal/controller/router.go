package controller

import (
	"large-scale-multistructure-db/be/internal/controller/middleware"
	"large-scale-multistructure-db/be/internal/usecase"
	"net/http"

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

func Router(
	userUsecase usecase.User,
	barberShopUsecase usecase.BarberShop,
	appointmentUsecase usecase.Appointment,
	calendarUsecase usecase.Calendar,
) *gin.Engine {

	router := gin.Default()
	router.Use(CORSAllowAll())

	api := router.Group("/api")

	api.GET("/health/", func(c *gin.Context) { c.JSON(http.StatusOK, `{"message" : "ok"}`) })

	// create the routes based on the given usecases
	mr := middleware.NewMiddlewareRoutes(userUsecase)
	ur := NewUserRoutes(userUsecase)
	br := NewBarberShopRoutes(barberShopUsecase, calendarUsecase)
	ar := NewAppointmentRoutes(appointmentUsecase, userUsecase)

	// TODO :
	// - don't return hash password

	// link the path to the routes
	user := api.Group("/user")
	{
		user.POST("/", ur.Register)                                         // TESTED
		user.POST("/login/", ur.Login)                                      // TESTED
		user.GET("/self/", mr.RequireAuth, mr.MarkWithAuthID, ur.Show)      // TESTED
		user.DELETE("/self/", mr.RequireAuth, mr.MarkWithAuthID, ur.Delete) // TESTED
		user.POST("/lost_password/", ur.LostPassword)
		user.POST("/reset_password/", ur.ResetPassword)

		user.DELETE("/self/appointment", ar.DeleteSelfAppointment)
	}

	admin := api.Group("/admin")
	admin.Use(mr.RequireAdmin)
	{
		admin.GET("/user", ur.ShowAll)       // TESTED
		admin.GET("/user/:id", ur.Show)      // TESTED
		admin.DELETE("/user/:id", ur.Delete) // TESTED
		admin.PUT("/user/:id", ur.Modify)

		admin.POST("/barber_shop/", br.Create)
		admin.DELETE("/barber_shop/:id", br.Delete)

		admin.DELETE("/barber_shop/:id/calendar", br.Calendar)

		admin.DELETE("/barber_shop/:id/appointment", ar.DeleteAppointment)
		admin.DELETE("/barber_shop/:id/holidays", ar.SetHolidays)
	}

	barberShop := api.Group("/barber_shop")
	barberShop.Use(mr.RequireAuth)
	{
		barberShop.GET("", br.Find)
		barberShop.GET("/:id", br.Show)
		barberShop.PUT("/:id", mr.RequireBarber, br.Modify)

		// appointments
		barberShop.PUT("/:id/appointment", ar.Book)

	}

	return router
}
