package controller

import (
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/pkg/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userUseCase usecase.User
}

func NewUserRoutes(uc usecase.User) *UserRoutes {
	return &UserRoutes{
		userUseCase: uc,
	}
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ur *UserRoutes) Login(ctx *gin.Context) {

	input := LoginInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ur.userUseCase.Login(ctx, &entity.User{
		Password: input.Password,
		Email:    input.Email,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.CreateToken(user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}

type CreateUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ur *UserRoutes) CreateUser(ctx *gin.Context) {

	input := CreateUserInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ur.userUseCase.Store(ctx, &entity.User{
		Password: input.Password,
		Email:    input.Email,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (ur *UserRoutes) Show(ctx *gin.Context) {

	ID, err := strconv.ParseUint(
		ctx.Param("id"),
		10, 64,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ur.userUseCase.GetByID(ctx, uint(ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})

}

func (ur *UserRoutes) ShowAll(ctx *gin.Context) {

	email := ctx.Query("email")

	users, err := ur.userUseCase.List(ctx, email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"users": users})
}

func (ur *UserRoutes) Delete(ctx *gin.Context) {
	// GET PARAMS
	ID, err := strconv.ParseUint(
		ctx.Param("id"),
		10, 64,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ur.userUseCase.DeleteByID(ctx, uint(ID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "user deleted"})
}

type ModifyUserInput struct {
	Email    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ur *UserRoutes) Modify(ctx *gin.Context) {
	input := ModifyUserInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &entity.User{
		Email: input.Email,
	}

	ID, err := strconv.ParseUint(
		ctx.Param("id"),
		10, 64,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ur.userUseCase.ModifyByID(ctx, uint(ID), user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "user modified"})

}
