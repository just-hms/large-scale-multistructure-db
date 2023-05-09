package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
)

type GeocodingRoutes struct {
	geocoding usecase.Geocoding
}

func NewGeocodingRoutes(g usecase.Geocoding) *GeocodingRoutes {
	return &GeocodingRoutes{
		geocoding: g,
	}
}

type SearchInput struct {
	Area string `json:"area" binding:"required"`
}

// Search returns a list of geocoding informations
//
// @Summary Given an area returns an array of geocodes
// @Description Given some information about a place returns a list of possible location with some other information about them in order of importance
// @Tags Geocoding
// @Accept  json
// @Produce  json
// @Param input body SearchInput true "Area to search"
// @Success 200 {object} map[string][]entity.GeocodingInfo
// @Failure 400 {object} string	"Bad request"
// @Failure 401 {object} string	"Unauthorized"
// @Router /geocoding/search [post]
func (ur *GeocodingRoutes) Search(ctx *gin.Context) {
	input := SearchInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ur.geocoding.Search(ctx, input.Area)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"geocodes": res})
}
