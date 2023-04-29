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
