package middleware

import (
	"net/http"
	"strings"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MiddlewareRoutes struct {
	userUseCase  usecase.User
	tokenUseCase usecase.Token
}

func NewMiddlewareRoutes(uc usecase.User, t usecase.Token) *MiddlewareRoutes {
	return &MiddlewareRoutes{
		userUseCase:  uc,
		tokenUseCase: t,
	}
}

func ExtractTokenFromRequest(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token
}

func (mr *MiddlewareRoutes) RequireAuth(ctx *gin.Context) {

	token := ExtractTokenFromRequest(ctx)
	tokenID, err := mr.tokenUseCase.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if _, err := mr.userUseCase.GetByID(ctx, tokenID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx.Next()
}

func (mr *MiddlewareRoutes) RequireAdmin(ctx *gin.Context) {

	token := ExtractTokenFromRequest(ctx)

	tokenID, err := mr.tokenUseCase.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := mr.userUseCase.GetByID(ctx, tokenID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if user.Type != entity.ADMIN {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	ctx.Next()
}

func (mr *MiddlewareRoutes) MarkWithAuthID(ctx *gin.Context) {

	token := ExtractTokenFromRequest(ctx)
	tokenID, err := mr.tokenUseCase.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx.Params = append(ctx.Params, gin.Param{
		Key:   "id",
		Value: tokenID,
	})

}

func (mr *MiddlewareRoutes) RequireBarber(ctx *gin.Context) {

	token := ExtractTokenFromRequest(ctx)
	tokenID, err := mr.tokenUseCase.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := mr.userUseCase.GetByID(ctx, tokenID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if user.Type != entity.BARBER {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	barberShopID := ctx.Param("shopid")

	if barberShopID == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// check if the barber is in ones that the barber owns
	for _, b := range user.OwnedShops {
		if b.ID == barberShopID {
			ctx.Next()
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}
