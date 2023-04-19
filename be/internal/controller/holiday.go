package controller

import (
	"net/http"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"

	"github.com/gin-gonic/gin"
)

type HolidayRoutes struct {
	holidayUseCase usecase.Holiday
}

func NewHolidayRoutes(uc usecase.Holiday) *HolidayRoutes {
	return &HolidayRoutes{
		holidayUseCase: uc,
	}
}

type SetHolidaysInput struct {
	Date                 time.Time `json:"date"`
	UnavailableEmployees int       `json:"unavailableEmployees"`
}

// Set holiday for a given Barber Shop
// @Summary Set holiday
// @Description Set holiday for a given Barber Shop
// @Tags Holiday
// @Param shopid path string true "Barber Shop ID"
// @Param input body SetHolidaysInput true "Set Holidays Input"
// @Accept json
// @Produce json
// @Success 202 {object} string ""
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Router /barbershop/{shopid}/holiday [post]
func (r *HolidayRoutes) Set(ctx *gin.Context) {
	ID := ctx.Param("shopid")

	input := SetHolidaysInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.holidayUseCase.Set(ctx, ID, input.Date, input.UnavailableEmployees)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
