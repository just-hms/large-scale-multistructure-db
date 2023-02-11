package controller

import (
	"fmt"
	"large-scale-multistructure-db/be/internal/controller/middleware"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BarberShopRoutes struct {
	barberShopUseCase usecase.BarberShop
}

func NewBarberShopRoutes(uc usecase.BarberShop) *BarberShopRoutes {
	return &BarberShopRoutes{
		barberShopUseCase: uc,
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
	Latitude        float64 `json:"latitude" binding:"required"`
	Longitude       float64 `json:"longitude" binding:"required"`
	EmployeesNumber int     `json:"employees_number" binding:"required"`
	Rating          float64 `json:"rating"`
}

func (br *BarberShopRoutes) Create(ctx *gin.Context) {

	input := CreateBarbershopInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := br.barberShopUseCase.Store(ctx, &entity.BarberShop{
		Name: input.Name,
		Coordinates: entity.Coordinates{
			Latitude:  fmt.Sprintf("%f", input.Latitude),
			Longitude: fmt.Sprintf("%f", input.Longitude),
		},
		Employees: input.EmployeesNumber,
		Rating:    input.Rating,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (br *BarberShopRoutes) Show(ctx *gin.Context) {

	ID := ctx.Param("id")

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
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Employees int     `json:"employees"`
	Rating    float64 `json:"rating"`
}

func (br *BarberShopRoutes) Modify(ctx *gin.Context) {

	ID := ctx.Param("id")

	input := ModifyBarberShopInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := br.barberShopUseCase.ModifyByID(ctx, ID, &entity.BarberShop{
		Name: input.Name,
		Coordinates: entity.Coordinates{
			Latitude:  fmt.Sprintf("%f", input.Latitude),
			Longitude: fmt.Sprintf("%f", input.Longitude),
		},
		Employees: input.Employees,
		Rating:    input.Rating,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

func (br *BarberShopRoutes) Delete(ctx *gin.Context) {

	ID := ctx.Param("id")

	err := br.barberShopUseCase.DeleteByID(ctx, ID)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
