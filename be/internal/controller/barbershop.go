package controller

import (
	"fmt"
	"net/http"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller/middleware"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type BarberShopRoutes struct {
	barberShopUseCase usecase.BarberShop
	calendarUseCase   usecase.Calendar
}

func NewBarberShopRoutes(uc usecase.BarberShop, cl usecase.Calendar) *BarberShopRoutes {
	return &BarberShopRoutes{
		barberShopUseCase: uc,
		calendarUseCase:   cl,
	}
}

func (br *BarberShopRoutes) Find(ctx *gin.Context) {

	barbers, err := br.barberShopUseCase.Find(
		ctx,
		ctx.Param("lat"),
		ctx.Param("lon"),
		ctx.Param("name"),
		ctx.Param("radius"),
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"barberShops": barbers})
}

// TODO add more things
type CreateBarbershopInput struct {
	Name            string  `json:"name" binding:"required"`
	Latitude        float64 `json:"Latitude" binding:"required"`
	Longitude       float64 `json:"Longitude" binding:"required"`
	EmployeesNumber int     `json:"employees_number" binding:"required"`
}

func (br *BarberShopRoutes) Create(ctx *gin.Context) {

	input := CreateBarbershopInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := br.barberShopUseCase.Store(ctx, &entity.BarberShop{
		Name:      input.Name,
		Latitude:  fmt.Sprintf("%f", input.Latitude),
		Longitude: fmt.Sprintf("%f", input.Longitude),
		Employees: input.EmployeesNumber,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (br *BarberShopRoutes) Show(ctx *gin.Context) {

	ID := ctx.Param("shopid")

	// TODO mark it with the middleware
	token := middleware.ExtractTokenFromRequest(ctx)
	tokenID, err := jwt.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	// TODO mark it with the middleware

	barberShop, err := br.barberShopUseCase.GetByID(ctx, tokenID, ID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": barberShop})
}

type ModifyBarberShopInput struct {
	Name      string  `json:"name"`
	Lat       float64 `json:"Latitude"`
	Lon       float64 `json:"Longitude"`
	Employees int     `json:"employees"`
}

func (br *BarberShopRoutes) Modify(ctx *gin.Context) {

	ID := ctx.Param("shopid")

	input := ModifyBarberShopInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := br.barberShopUseCase.ModifyByID(ctx, ID, &entity.BarberShop{
		Name:      input.Name,
		Latitude:  fmt.Sprintf("%f", input.Lat),
		Longitude: fmt.Sprintf("%f", input.Lon),
		Employees: input.Employees,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

func (br *BarberShopRoutes) Delete(ctx *gin.Context) {

	ID := ctx.Param("shopid")

	err := br.barberShopUseCase.DeleteByID(ctx, ID)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

func (br *BarberShopRoutes) Calendar(ctx *gin.Context) {

	ID := ctx.Param("shopid")
	slots, err := br.calendarUseCase.GetByBarberShopID(ctx, ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"calendar": slots})

}
