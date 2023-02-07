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
	)

	for _, uc := range usecases {

		switch u := uc.(type) {

		case *usecase.UserUseCase:
			mr = middleware.NewMiddlewareRoutes(u)
			ur = NewUserRoutes(u)
		}
	}

	// TODO :
	// - fix trailing /
	// - don't return hash password
	// - return the ID

	// link the path to the routes
	users := router.Group("/user")
	{
		users.POST("/", ur.Register)
		users.POST("/login", ur.Login)
		users.GET("/self", mr.RequireAuth, mr.MarkWithAuthID, ur.Show)
		users.DELETE("/self", mr.RequireAuth, mr.MarkWithAuthID, ur.Delete)

		// TODO : test
		users.POST("/lost_password", ur.LostPassword)
		users.POST("/reset_password", ur.ResetPassword)
	}

	admin := router.Group("/admin")
	admin.Use(mr.RequireAdmin)
	{
		admin.GET("/user", ur.ShowAll)
		admin.GET("/user/:id", ur.Show)
		admin.DELETE("/user/:id", ur.Delete)

		// TODO : test
		admin.PUT("/user/:id", ur.Modify)
	}

	return router
}
