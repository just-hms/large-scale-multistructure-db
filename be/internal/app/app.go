package app

import (
	"fmt"
	"large-scale-multistructure-db/be/internal/controller"
	"large-scale-multistructure-db/be/internal/controller/middleware"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/internal/usecase/repo"
	"large-scale-multistructure-db/be/pkg/mongo"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {

	// Repository

	mongo, err := mongo.New()
	if err != nil {
		fmt.Printf("mongo-error: %s", err.Error())
		return
	}

	// UseCase
	useruc := usecase.NewUserUseCase(
		repo.NewUserRepo(mongo),
	)

	// HTTP router

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	m := middleware.NewMiddlewareRoutes(useruc)
	ur := controller.NewUserRoutes(useruc)

	users := router.Group("/user")
	{
		router.POST("/", m.RequireAdmin, ur.CreateUser)
		users.GET("/login", ur.Login)
	}

	router.Run()

}
