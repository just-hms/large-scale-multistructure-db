package controller

import (
	"large-scale-multistructure-db/be/internal/controller/middleware"
	"large-scale-multistructure-db/be/internal/usecase"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(usecases []usecase.Usecase) *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, `{"message" : "ok"}`) })

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
	user := router.Group("/user")
	{
		user.POST("/", ur.Register)                                        // TESTED
		user.POST("/login", ur.Login)                                      // TESTED
		user.GET("/self", mr.RequireAuth, mr.MarkWithAuthID, ur.Show)      // TESTED
		user.DELETE("/self", mr.RequireAuth, mr.MarkWithAuthID, ur.Delete) // TESTED
		user.POST("/lost_password", ur.LostPassword)
		user.POST("/reset_password", ur.ResetPassword)
	}

	admin := router.Group("/admin")
	admin.Use(mr.RequireAdmin)
	{
		admin.GET("/user", ur.ShowAll)       // TESTED
		admin.GET("/user/:id", ur.Show)      // TESTED
		admin.DELETE("/user/:id", ur.Delete) // TESTED
		admin.PUT("/user/:id", ur.Modify)
	}

	barberShop := router.Group("/barber_shop")
	barberShop.Use(mr.RequireAuth)
	{
		barberShop.GET("/", br.Find)
		barberShop.GET("/:id", br.Show)
		barberShop.POST("/", br.Create)
		barberShop.PUT("/", br.Modify)

		// TODO: require barber
		barberShop.DELETE("/", br.Delete)
	}

	return router
}
