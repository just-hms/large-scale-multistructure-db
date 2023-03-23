package controller

import (
	"net/http"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"

	"github.com/gin-gonic/gin"
)

type HolidayRoutes struct {
	holidayUseCase usecase.HolidayUseCase
}

func NewHolidayRoutes(uc usecase.HolidayUseCase) *HolidayRoutes {
	return &HolidayRoutes{
		holidayUseCase: uc,
	}
}

type SetHolidaysInput struct {
	Date                 time.Time `json:"date"`
	UnavailableEmployees int       `json:"unavailableEmployees"`
}

func (r *HolidayRoutes) Set(ctx *gin.Context) {
	ID := ctx.Param("id")

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
