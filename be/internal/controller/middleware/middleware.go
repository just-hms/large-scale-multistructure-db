package middleware

import (
	"large-scale-multistructure-db/be/internal/usecase"
	"large-scale-multistructure-db/be/pkg/jwt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MiddlewareRoutes struct {
	userUseCase usecase.User
}

func NewMiddlewareRoutes(uc usecase.User) *MiddlewareRoutes {
	return &MiddlewareRoutes{
		userUseCase: uc,
	}
}

func ExtractTokenFromRequest(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token
}

func (mr *MiddlewareRoutes) RequireAuth(ctx *gin.Context) {

	token := ExtractTokenFromRequest(ctx)
	tokenID, err := jwt.ExtractTokenID(token)

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
	tokenID, err := jwt.ExtractTokenID(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := mr.userUseCase.GetByID(ctx, tokenID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if !user.IsAdmin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	ctx.Next()
}

func (mr *MiddlewareRoutes) Self(ctx *gin.Context) {

	token := ExtractTokenFromRequest(ctx)
	tokenID, _ := jwt.ExtractTokenID(token)

	ctx.Params = append(ctx.Params, gin.Param{
		Key:   "id",
		Value: strconv.FormatUint(uint64(tokenID), 10),
	})

}
