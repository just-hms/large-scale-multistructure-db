package controller

import (
	"net/http"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/jwt"

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

// Login logs in a user and returns a JWT token.
// It expects a JSON payload containing the user's email and password.
// It returns a JWT token that can be used for subsequent authenticated requests.
//
// @Summary Logs in a user and returns a JWT token
// @Description Logs in a user with the provided email and password and returns a JWT token for subsequent authenticated requests
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param input body LoginInput true "User email and password"
// @Success 200 {string} map[string]string {"token" : "token"}
// @Failure 400 {object} string	"Bad request"
// @Failure 401 {object} string	"Unauthorized"
// @Router /user/login [post]
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

// Register creates a new user with the given email and password
// @Summary Create new user
// @Description Create new user with the provided email and password
// @Tags User
// @Accept  json
// @Produce  json
// @Param input body RegisterInput true "User credentials"
// @Success 201 {string} string ""
// @Failure 401 {object} string	"Unauthorized"
// @Failure 400 {object} string	"Bad request"
// @Router /user [post]
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

// Show Shows the user informations
// @Summary Show user information by ID
// @Description Get user information by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]entity.User
// @Failure 401 {object} string	"Unauthorized"
// @Failure 404 {object} map[string]string
// @Router /user/self [get]
// @Router /admin/user/{id} [get]
func (ur *UserRoutes) Show(ctx *gin.Context) {

	ID := ctx.Param("id")

	user, err := ur.userUseCase.GetByID(ctx, ID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// ShowAll Shows the list of all users
// @Summary Show list of users
// @Description Get a list of users filtered by email
// @Tags User
// @Accept json
// @Produce json
// @Param email query string false "Email filter"
// @Success 200 {object} map[string][]entity.User
// @Failure 400 {object} map[string]string
// @Failure 401 {object} string	"Unauthorized"
// @Router /admin/user [get]
func (ur *UserRoutes) ShowAll(ctx *gin.Context) {

	email := ctx.Query("email")

	users, err := ur.userUseCase.List(ctx, email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

// Delete deletes a user with the given ID.
// It accepts a path parameter "id" to specify the ID of the user to be deleted.
// @Summary Delete user by ID
// @Description Delete a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 202 {object} string ""
// @Failure 401 {object} string	"Unauthorized"
// @Failure 404 {object} map[string]string
// @Router /admin/user/{id} [delete]
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

// Modify modifies a user with the given ID by updating their email and barbershop associations.
// The "email" field specifies the updated email address of the user.
// The "barbershopsId" field specifies an array of barbershop IDs to associate the user with.
// @Summary Modify user by ID
// @Description Modify a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param input body ModifyUserInput true "Modify user input"
// @Success 202 {object} string ""
// @Failure 400 {object} map[string]string
// @Router /admin/user/{id} [put]
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

	ctx.JSON(http.StatusAccepted, gin.H{})
}

type LostPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

// LostPassword generates a password reset token for the user with the given email address and sends it to them via email.
// It accepts a JSON payload in the request body with the field "email", which specifies the email address of the user to reset the password for.
// If successful, it returns a JSON response with the password reset token.
// @Summary Request password reset
// @Description Generate a password reset token and send it to the user's email address
// @Tags User
// @Accept json
// @Produce json
// @Param input body LostPasswordInput true "Lost password input"
// @Success 202 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /user/lostpassword [post]
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

	// THIS IS A POC YOU SHOULD RETURN 200 AND SEND THE E-MAIL WITH THE TOKEN HERE

	ctx.JSON(http.StatusAccepted, gin.H{"resetToken": resetToken})
}

type ResetPasswordInput struct {
	NewPassword string `json:"newpassword" binding:"required"`
}

// ResetPassword resets the password for a user given a reset token
// @Summary Reset user password
// @Description Reset password for a user given a reset token
// @Tags User
// @Accept json
// @Produce json
// @Param reset_token path string true "Reset token"
// @Param newpassword body string true "New password"
// @Success 202 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /user/resetpassword/{resettoken} [post]
func (ur *UserRoutes) ResetPassword(ctx *gin.Context) {

	token := ctx.Param("resettoken")
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

	if err := ur.userUseCase.ResetPassword(ctx, userID, input.NewPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "password correctly resetted"})
}
