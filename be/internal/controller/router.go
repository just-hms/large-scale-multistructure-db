package controller

import (
	"net/http"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller/middleware"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"

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

func Router(ucs map[byte]usecase.Usecase, production bool) *gin.Engine {

	if production {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(CORSAllowAll())

	api := router.Group("/api")

	api.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, `{"message" : "ok"}`) })

	// create the routes based on the given usecases
	mr := middleware.NewMiddlewareRoutes(ucs[usecase.USER].(usecase.User))
	ur := NewUserRoutes(ucs[usecase.USER].(usecase.User))

	br := NewBarberShopRoutes(
		ucs[usecase.BARBER_SHOP].(usecase.BarberShop),
		ucs[usecase.CALENDAR].(usecase.Calendar),
	)

	ar := NewAppointmentRoutes(
		ucs[usecase.APPOINTMENT].(*usecase.AppoinmentUseCase),
		ucs[usecase.USER].(*usecase.UserUseCase),
	)

	// link the path to the routes
	user := api.Group("/user")
	{
		user.POST("", ur.Register)
		user.POST("/login", ur.Login)
		user.GET("/self", mr.RequireAuth, mr.MarkWithAuthID, ur.Show)
		user.DELETE("/self", mr.RequireAuth, mr.MarkWithAuthID, ur.Delete)
		user.POST("/lost_password", ur.LostPassword)
		user.POST("/reset_password", ur.ResetPassword)

		user.DELETE("/self/appointment", mr.RequireAuth, ar.DeleteSelfAppointment)
	}

	admin := api.Group("/admin")
	admin.Use(mr.RequireAdmin)
	{
		admin.GET("/user", ur.ShowAll)
		admin.GET("/user/:id", ur.Show)
		admin.DELETE("/user/:id", ur.Delete)
		admin.PUT("/user/:id", ur.Modify)

		admin.POST("/barber_shop", br.Create)
		admin.DELETE("/barber_shop/:shopid", br.Delete)

		admin.DELETE("/barber_shop/:shopid/holidays", ar.SetHolidays)
	}

	barberShop := api.Group("/barber_shop")
	barberShop.Use(mr.RequireAuth)
	{
		barberShop.POST("", br.Find)
		barberShop.GET("/:shopid", br.Show)

		barberShop.PUT("/:shopid", mr.RequireBarber, br.Modify)

		barberShop.DELETE("/:shopid/appointment/:appointmentid", mr.RequireBarber, ar.DeleteAppointment)
		barberShop.POST("/:shopid/appointment", ar.Book)

		barberShop.GET("/:shopid/calendar", br.Calendar)
	}

	return router
}
