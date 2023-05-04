package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/controller/middleware"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"
)

type ReviewRoutes struct {
	reviewUseCase usecase.Review
	userUseCase   usecase.User
	tokenUseCase  usecase.Token
}

func NewReviewRoutes(uc usecase.Review, t usecase.Token) *ReviewRoutes {
	return &ReviewRoutes{
		reviewUseCase: uc,
		tokenUseCase:  t,
	}
}

type StoreReviewInput struct {
	Rating  int    `json:"rating"`
	Content string `json:"content"`
}

// @Router /barbershop/{shopid}/review [post]
func (rr *ReviewRoutes) Store(ctx *gin.Context) {

	input := StoreReviewInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := middleware.ExtractTokenFromRequest(ctx)
	tokenID, err := rr.tokenUseCase.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	shopID := ctx.Param("shopid")

	err = rr.reviewUseCase.Store(ctx, &entity.Review{
		Rating:  input.Rating,
		Content: input.Content,
		UserID:  tokenID,
	}, shopID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})

}

// @Router /barbershop/{shopid}/review [get]
func (rr *ReviewRoutes) ShowByShop(ctx *gin.Context) {

	shopID := ctx.Param("shopid")

	reviews, err := rr.reviewUseCase.GetByBarberShopID(ctx, shopID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

// @Router /barbershop/{shopid}/review/{reviewid} [delete]
func (rr *ReviewRoutes) Delete(ctx *gin.Context) {

	shopID := ctx.Param("shopid")
	reviewID := ctx.Param("reviewid")

	err := rr.reviewUseCase.DeleteByID(ctx, shopID, reviewID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})

}

type ReviewVoteInput struct {
	Upvote bool `json:"upvote"`
}

// @Router /barbershop/{shopid}/review/{reviewid}/vote [post]
func (rr *ReviewRoutes) Vote(ctx *gin.Context) {

	input := ReviewVoteInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shopID := ctx.Param("shopid")
	reviewID := ctx.Param("reviewid")

	token := middleware.ExtractTokenFromRequest(ctx)
	tokenID, err := rr.tokenUseCase.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if input.Upvote {
		err = rr.reviewUseCase.UpVoteByID(ctx, tokenID, shopID, reviewID)
	} else {
		err = rr.reviewUseCase.DownVoteByID(ctx, tokenID, shopID, reviewID)
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})

}

// @Router /barbershop/{shopid}/review/{reviewid}/vote [delete]
func (rr *ReviewRoutes) RemoveVote(ctx *gin.Context) {

	shopID := ctx.Param("shopid")
	reviewID := ctx.Param("reviewid")

	token := middleware.ExtractTokenFromRequest(ctx)
	tokenID, err := rr.tokenUseCase.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = rr.reviewUseCase.RemoveVoteByID(ctx, tokenID, shopID, reviewID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})

}
