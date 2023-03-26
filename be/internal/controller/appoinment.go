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
		UserID:       user.ID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})

}

func (ur *AppointmentRoutes) DeleteAppointment(ctx *gin.Context) {

	SHOP_ID := ctx.Param("shopid")
	ID := ctx.Param("appointmentid")

	err := ur.appoinmentUseCase.Cancel(ctx, &entity.Appointment{
		BarbershopID: SHOP_ID,
		ID:           ID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

func (ur *AppointmentRoutes) SetHolidays(ctx *gin.Context) {

}
