package controller

import (
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/pkg/jwt"
	"net/http"

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
	Email    string `json:"email" binding:"required,email"`
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

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ur *UserRoutes) Register(ctx *gin.Context) {

	input := RegisterInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ur.userUseCase.Store(ctx, &entity.User{
		Password: input.Password,
		Email:    input.Email,
		Type:     entity.USER,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (ur *UserRoutes) Show(ctx *gin.Context) {

	ID := ctx.Param("id")

	user, err := ur.userUseCase.GetByID(ctx, ID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (ur *UserRoutes) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	if err := ur.userUseCase.DeleteByID(ctx, ID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "user deleted"})
}

type ModifyUserInput struct {
	Email         string   `json:"email"`
	BarbershopsID []string `json:"barbershopsId"`
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

	ID := ctx.Param("id")

	if err := ur.userUseCase.ModifyByID(ctx, ID, user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ur.userUseCase.EditShopsByIDs(ctx, ID, input.BarbershopsID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "user modified"})

}

type LostPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

func (ur *UserRoutes) LostPassword(ctx *gin.Context) {
	input := LostPasswordInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resetToken, err := ur.userUseCase.LostPassword(ctx, input.Email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// THIS IS A POC YOU SHOULD RETUNR 200 AND SEND THE E-MAIL WITH THE TOKEN HERE

	ctx.JSON(http.StatusAccepted, gin.H{"resetToken": resetToken})
}

type ResetPasswordInput struct {
	Password string `json:"newpassword" binding:"required"`
}

func (ur *UserRoutes) ResetPassword(ctx *gin.Context) {

	token := ctx.Param("reset_token")
	userID, err := jwt.ExtractTokenID(token)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := ResetPasswordInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ur.userUseCase.ResetPassword(ctx, userID, input.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// THIS IS A POC YOU SHOULD RETUNR 200 AND SEND THE E-MAIL WITH THE TOKEN HERE

	ctx.JSON(http.StatusAccepted, gin.H{"message": "password correctly resetted"})
}
