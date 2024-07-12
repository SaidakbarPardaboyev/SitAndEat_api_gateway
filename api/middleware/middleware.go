package middleware

import (
	"api_gateway/api/token"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				fmt.Errorf("authorization header is required"))
			return
		}

		valid, err := token.ValidateToken(auth)
		if err != nil || !valid {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				fmt.Errorf("invalid token: %s", err))
			return
		}

		claims, err := token.ExtractClaims(auth)
		if err != nil || !valid {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				fmt.Errorf("invalid token claims: %s", err))
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
