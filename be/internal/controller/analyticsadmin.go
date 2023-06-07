package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller/middleware"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
)

type AdminAnalyticsRoutes struct {
	adminAnalyticsUseCase usecase.AdminAnalytics
	tokenUseCase          usecase.Token
}

func NewAdminAnalyticsRoutes(aa usecase.AdminAnalytics, t usecase.Token) *AdminAnalyticsRoutes {
	return &AdminAnalyticsRoutes{
		adminAnalyticsUseCase: aa,
		tokenUseCase:          t,
	}
}

// GetAnalytics fetches all the analytics for an Admin
//
// @Summary Retrieve analytics for an Admin
// @Description Retrieves analytics of the whole DB for an Admin.
// @Tags Barbershop
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]entity.BarberAnalytics
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Router /barbershop/{shopid}/analytics [get]
func (aa *AdminAnalyticsRoutes) GetAnalytics(ctx *gin.Context) {

	token := middleware.ExtractTokenFromRequest(ctx)
	_, err := aa.tokenUseCase.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	analytics, err := aa.adminAnalyticsUseCase.GetAdminAnalytics(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"adminAnalytics": analytics})
}
