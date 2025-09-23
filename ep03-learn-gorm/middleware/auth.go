package middleware

import (
	"ep03-learn-gorm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error : ": "missing token"})
			return 
		}

		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error :" : "invalid token"})
			return 
		}

		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}