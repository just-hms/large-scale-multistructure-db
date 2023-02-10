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

func Router(usecases []usecase.Usecase) *gin.Engine {

	router := gin.Default()
	router.Use(CORSAllowAll())

	api := router.Group("/api")

	api.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, `{"message" : "ok"}`) })

	// create the routes based on the given usecases
	var (
		mr *middleware.MiddlewareRoutes
		ur *UserRoutes
		br *BarberShopRoutes
	)

	for _, uc := range usecases {

		switch u := uc.(type) {

		// TODO : check this in the https://github.com/evrone/go-clean-template

		case *usecase.UserUseCase:
			mr = middleware.NewMiddlewareRoutes(u)
			ur = NewUserRoutes(u)
		case *usecase.BarberShopUseCase:
			br = NewBarberShopRoutes(u)
		}
	}

	// TODO :
	// - fix trailing /
	// - don't return hash password
	// - return the ID

	// link the path to the routes
	user := api.Group("/user")
	{
		user.POST("/", ur.Register)                                        // TESTED
		user.POST("/login", ur.Login)                                      // TESTED
		user.GET("/self", mr.RequireAuth, mr.MarkWithAuthID, ur.Show)      // TESTED
		user.DELETE("/self", mr.RequireAuth, mr.MarkWithAuthID, ur.Delete) // TESTED
		user.POST("/lost_password", ur.LostPassword)
		user.POST("/reset_password", ur.ResetPassword)
	}

	admin := api.Group("/admin")
	admin.Use(mr.RequireAdmin)
	{
		admin.GET("/user", ur.ShowAll)       // TESTED
		admin.GET("/user/:id", ur.Show)      // TESTED
		admin.DELETE("/user/:id", ur.Delete) // TESTED
		admin.PUT("/user/:id", ur.Modify)

		admin.POST("/barber_shop", br.Create)
		admin.DELETE("/", br.Delete)
	}

	// TODO: require barber
	barberShop := api.Group("/barber_shop")
	barberShop.Use(mr.RequireAuth)
	{
		barberShop.GET("/", br.Find)
		barberShop.GET("/:id", br.Show)
		barberShop.PUT("/", br.Modify)
	}

	return router
}
