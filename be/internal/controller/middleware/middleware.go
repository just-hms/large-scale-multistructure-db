package middleware

import (
	"context"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/internal/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type MiddlewareRoutes struct {
	uc usecase.User
}

func NewMiddlewareRoutes(uc usecase.User) *MiddlewareRoutes {
	return &MiddlewareRoutes{
		uc: uc,
	}
}

func extractTokenFromRequest(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token
}

func (mr *MiddlewareRoutes) RequireAuth(ctx *gin.Context) {

	token := extractTokenFromRequest(ctx)

	useruc := usecase.UserUseCase{}

	if _, err := useruc.GetByToken(ctx, token); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx.Next()
}

func (mr *MiddlewareRoutes) RequireAdmin(ctx *gin.Context) {

	token := extractTokenFromRequest(ctx)

	useruc := usecase.UserUseCase{}

	var (
		user *entity.User
		err  error
	)

	if user, err = useruc.GetByToken(context.TODO(), token); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if !user.IsAdmin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	ctx.Next()
}
