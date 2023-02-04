package app

import (
	"fmt"
	"large-scale-multistructure-db/be/internal/controller"
	"large-scale-multistructure-db/be/internal/controller/middleware"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/internal/usecase/auth"
	"large-scale-multistructure-db/be/internal/usecase/repo"
	"large-scale-multistructure-db/be/pkg/mongo"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := Router()
	router.Run()
}

// TODO : fix this to devide the router from the rest, and put it in controllers

func Router() *gin.Engine {
	// Repository

	mongo, err := mongo.New()
	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return gin.Default()
	}

	// redis := redis.New()

	// UseCase
	useruc := usecase.NewUserUseCase(
		repo.NewUserRepo(mongo),
		auth.NewPasswordAuth(),
	)

	// TODO : move this inside the controller
	// HTTP router

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	mr := middleware.NewMiddlewareRoutes(useruc)
	ur := controller.NewUserRoutes(useruc)

	users := router.Group("/user")
	{
		users.GET("/login", ur.Login)
		router.POST("/", mr.RequireAuth, ur.CreateUser)
		users.GET("/self", mr.RequireAuth, mr.Self, ur.Show)
		users.DELETE("/self", mr.RequireAuth, mr.Self, ur.Delete)
	}

	admin := router.Group("/admin")
	admin.Use(mr.RequireAdmin)
	{
		admin.GET("/user", ur.ShowAll)
		admin.GET("/user/:id", ur.Show)
		admin.DELETE("/user/:id", ur.Delete)
		admin.PUT("/user/:id", ur.Modify)
	}

	return router

}
