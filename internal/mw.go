package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	unauthorized = "unauthorized"
	claimsSubKey = "sub"
)

func (api *api) authGuard(ctx *gin.Context) {
	tokenString, err := ctx.Cookie(cookieAccessTokenKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorized})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_AT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims[claimsSubKey].(string)

		user, err := api.userRepo.FindById(ctx.Request.Context(), userId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorized})
			return
		}
		if user.IsDeleted {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorized})
			return
		}

		jwt, err := api.jwtRepo.FindByToken(ctx.Request.Context(), tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorized})
			return
		}
		if jwt.IsDeleted {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorized})
			return
		}

		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), ctxUserKey, user))
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorized})
	}
}
