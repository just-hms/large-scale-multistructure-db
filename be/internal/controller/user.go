package controller

import (
	"context"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	uc usecase.User
}

func NewUserRoutes(uc usecase.User) *UserRoutes {
	return &UserRoutes{
		uc: uc,
	}
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ur *UserRoutes) Login(ctx *gin.Context) {

	input := LoginInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var (
		err   error
		token string
	)

	if _, token, err = ur.uc.Login(context.TODO(), input.Username, input.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}

type CreateUserInput struct {
	Email    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ur *UserRoutes) CreateUser(ctx *gin.Context) {

	input := CreateUserInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ur.uc.Store(context.TODO(), &entity.User{
		Password: input.Password,
		Email:    input.Email,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})

}
