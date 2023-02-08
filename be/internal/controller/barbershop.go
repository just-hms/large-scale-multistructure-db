package controller

import (
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

type CreateBarbershopInput struct {
	Name            string  `json:"name" binding:"required"`
	Latitude        float64 `json:"latitude" binding:"required"`
	Longitude       float64 `json:"longitude" binding:"required"`
	EmployeesNumber int     `json:"employees_number" binding:"required"`
	AverageRating   float64 `json:"average_rating"`
}

func (br *BarberShopRoutes) Create(ctx *gin.Context) {

	input := CreateBarbershopInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := br.barberShopUseCase.Store(ctx, &entity.BarberShop{
		Name:            input.Name,
		Latitude:        input.Latitude,
		Longitude:       input.Longitude,
		EmployeesNumber: input.EmployeesNumber,
		AverageRating:   input.AverageRating,
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

type ModifyBarbershopInput struct {
	Name            string  `json:"name" binding:"required"`
	Latitude        float64 `json:"latitude" binding:"required"`
	Longitude       float64 `json:"longitude" binding:"required"`
	EmployeesNumber int     `json:"employees_number" binding:"required"`
	AverageRating   float64 `json:"average_rating"`
}

func (br *BarberShopRoutes) Modify(ctx *gin.Context) {

	ID := ctx.Param("id")

	input := ModifyBarbershopInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := br.barberShopUseCase.ModifyByID(ctx, ID, &entity.BarberShop{
		Name:            input.Name,
		Latitude:        input.Latitude,
		Longitude:       input.Longitude,
		EmployeesNumber: input.EmployeesNumber,
		AverageRating:   input.AverageRating,
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
