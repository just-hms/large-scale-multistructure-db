package controller

import (
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

type FindBarbershopInput struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Radius    float64 `json:"radius"`
}

// Find handles a POST request to find barbershops within a certain radius.
//
// @Summary Find barbershops within a certain radius
// @Description Finds barbershops within a certain radius of the given coordinates and name.
// @Tags Barbershop
// @Accept json
// @Produce json
// @Param input body FindBarbershopInput true "The input parameters for finding barbershops"
// @Success 200 {object} map[string][]entity.BarberShop
// @Failure 401 {object} string "Unauthorized"
// @Failure 400 {object} string "Bad request"
// @Router /barbershop [post]
func (br *BarberShopRoutes) Find(ctx *gin.Context) {

	input := FindBarbershopInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	barbers, err := br.barberShopUseCase.Find(
		ctx,
		input.Latitude,
		input.Longitude,
		input.Name,
		input.Radius,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"barberShops": barbers})
}

type CreateBarbershopInput struct {
	Name            string  `json:"name" binding:"required"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	EmployeesNumber int     `json:"employees_number"`
}

// Create handles a POST request to create a new barbershop.
//
// @Summary Create a new barbershop
// @Description Creates a new barbershop with the given name, location and number of employees.
// @Tags Barbershop
// @Accept json
// @Produce json
// @Param input body CreateBarbershopInput true "The input parameters for creating a new barbershop"
// @Success 201 {object} string ""
// @Failure 401 {object} string "Unauthorized"
// @Failure 400 {object} string "Bad request"
// @Router /admin/barbershop [post]
func (br *BarberShopRoutes) Create(ctx *gin.Context) {

	input := CreateBarbershopInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := br.barberShopUseCase.Store(ctx, &entity.BarberShop{
		Name:      input.Name,
		Location:  entity.NewLocation(input.Latitude, input.Longitude),
		Employees: input.EmployeesNumber,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

// Show handles a GET request to retrieve details of a barbershop.
//
// @Summary Retrieve details of a barbershop
// @Description Retrieves details of a barbershop for the given shop ID.
// @Tags Barbershop
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param shopid path string true "The ID of the barbershop"
// @Success 200 {object} map[string]entity.BarberShop
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Not Found"
// @Router /barbershop/{shopid} [get]
func (br *BarberShopRoutes) Show(ctx *gin.Context) {

	ID := ctx.Param("shopid")

	token := middleware.ExtractTokenFromRequest(ctx)
	tokenID, err := jwt.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	barberShop, err := br.barberShopUseCase.GetByID(ctx, tokenID, ID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"barbershop": barberShop})
}

type ModifyBarberShopInput struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Employees int     `json:"employees"`
}

// Modify handles a PUT request to modify details of a barbershop.
//
// @Summary Modify details of a barbershop
// @Description Modifies details of a barbershop for the given shop ID.
// @Tags Barbershop
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param shopid path string true "The ID of the barbershop"
// @Param input body ModifyBarberShopInput true "The updated barbershop details"
// @Success 202 {object} string ""
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Router /admin/barbershop/{shopid} [put]
func (br *BarberShopRoutes) Modify(ctx *gin.Context) {

	ID := ctx.Param("shopid")

	input := ModifyBarberShopInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := br.barberShopUseCase.ModifyByID(ctx, ID, &entity.BarberShop{
		Name:      input.Name,
		Location:  entity.NewLocation(input.Latitude, input.Longitude),
		Employees: input.Employees,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

// Delete handles a DELETE request to delete a barbershop by ID.
//
// @Summary Delete a barbershop by ID
// @Description Deletes a barbershop by ID.
// @Tags Barbershop
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param shopid path string true "The ID of the barbershop to delete"
// @Success 202 {object} string ""
// @Failure 401 {object} string "Unauthorized"
// @Router /admin/barbershop/{shopid} [delete]
func (br *BarberShopRoutes) Delete(ctx *gin.Context) {

	ID := ctx.Param("shopid")

	err := br.barberShopUseCase.DeleteByID(ctx, ID)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

// Calendar handles a GET request to get the calendar for a barbershop by ID.
//
// @Summary Get the calendar for a barbershop by ID
// @Description Gets the calendar for a barbershop by ID.
// @Tags Barbershop
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param shopid path string true "The ID of the barbershop to get the calendar for"
// @Success 202 {object} map[string][]entity.Slot
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Router /barbershop/{shopid}/calendar [get]
func (br *BarberShopRoutes) Calendar(ctx *gin.Context) {

	ID := ctx.Param("shopid")
	slots, err := br.calendarUseCase.GetByBarberShopID(ctx, ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"calendar": slots})

}
