package controller

import (
	"net/http"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller/middleware"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type AppointmentRoutes struct {
	appoinmentUseCase usecase.Appointment
	userUseCase       usecase.User
}

func NewAppointmentRoutes(uc usecase.Appointment, us usecase.User) *AppointmentRoutes {
	return &AppointmentRoutes{
		appoinmentUseCase: uc,
		userUseCase:       us,
	}
}

type BookAppointmentInput struct {
	DateTime time.Time `json:"dateTime"`
}

// Book handles a POST request to book a new appointment.
// @Summary Book a new appointment
// @Description Books a new appointment for the current user.
// @Tags Appointment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param shopid path string true "The ID of the barbershop"
// @Param input body BookAppointmentInput true "The appointment details"
// @Success 201 {object} string ""
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Router /barbershop/{shopid}/appointment [post]
func (ur *AppointmentRoutes) Book(ctx *gin.Context) {

	input := BookAppointmentInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := middleware.ExtractTokenFromRequest(ctx)
	tokenID, err := jwt.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ID := ctx.Param("shopid")

	err = ur.appoinmentUseCase.Book(ctx, &entity.Appointment{
		Start:        input.DateTime,
		BarbershopID: ID,
		UserID:       tokenID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

// @Summary Deletes the current user's appointment
// @Description Deletes the appointment of the current user
// @Tags Appointment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 202 {object}  string ""
// @Failure 400 {object}  string "Bad Request"
// @Failure 401 {object}  string "Unauthorized"
// @Router /user/self/appointment [delete]
func (ur *AppointmentRoutes) DeleteSelfAppointment(ctx *gin.Context) {

	token := middleware.ExtractTokenFromRequest(ctx)
	tokenID, err := jwt.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := ur.userUseCase.GetByID(ctx, tokenID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	err = ur.appoinmentUseCase.Cancel(ctx, user.CurrentAppointment)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})

}

// @Summary Deletes an appointment
// @Description Deletes an appointment at a specific barbershop
// @Tags Appointment
// @Accept json
// @Produce json
// @Param shopid path string true "ID of the barbershop"
// @Param appointmentid path string true "ID of the appointment"
// @Success 202 {object} string ""
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object}  string "Unauthorized"
// @Router /barbershop/{shopid}/appointment/{appointmentid} [delete]
func (ur *AppointmentRoutes) DeleteAppointment(ctx *gin.Context) {

	appointment, err := ur.appoinmentUseCase.GetByIDs(
		ctx,
		ctx.Param("shopid"),
		ctx.Param("appointmentid"),
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ur.appoinmentUseCase.Cancel(ctx, appointment)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
