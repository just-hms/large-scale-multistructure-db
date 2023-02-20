package controller

import (
	"large-scale-multistructure-db/be/internal/controller/middleware"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/pkg/jwt"
	"net/http"
	"time"

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

func (ur *AppointmentRoutes) Book(ctx *gin.Context) {

	// TODO check that is not a barber or admin and that there is no current appointment
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

	err = ur.appoinmentUseCase.Cancel(ctx, &entity.Appointment{
		BarbershopID: user.CurrentAppointment.BarbershopID,
		ID:           user.CurrentAppointment.ID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})

}

type DeleteAppointmentInput struct {
	AppointmentId string `json:"appointmentId"`
}

func (ur *AppointmentRoutes) DeleteAppointment(ctx *gin.Context) {

	input := DeleteAppointmentInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ID := ctx.Param("appointmentid")

	err := ur.appoinmentUseCase.Cancel(ctx, &entity.Appointment{
		BarbershopID: ID,
		ID:           input.AppointmentId,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (ur *AppointmentRoutes) SetHolidays(ctx *gin.Context) {

}
