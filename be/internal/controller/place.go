package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
)

type PlaceRoutes struct {
	place usecase.Place
}

func (ur *PlaceRoutes) Search(ctx *gin.Context) {

}
