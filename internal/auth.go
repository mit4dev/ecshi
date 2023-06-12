package api

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mit4dev/ecshi/domain"
)

var (
	errUserNotFound       = errors.New("user.not-found")
	errInvalidCredentials = errors.New("login.invalid-credentials")
)

type LoginDto struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (api *api) login(c *gin.Context) {
	var login LoginDto

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := api.userRepo.FindByEmail(c.Request.Context(), login.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errUserNotFound.Error()})
		return
	}

	if !compareHash(user.Password, login.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errInvalidCredentials.Error(), "hashed": user.Password, "password": login.Password})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   user.Id,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_AT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	api.jwtRepo.Store(c.Request.Context(), &domain.MongoJwt{
		UserId:      user.Id,
		AccessToken: tokenString,
	})

	year := 60 * 60 * 24 * 365
	c.SetCookie(cookieAccessTokenKey, tokenString, year, "/", "", false, true)
	c.JSON(http.StatusOK, user)
}

func (api *api) refresh(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not-implemented"})
}

func (api *api) logout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not-implemented"})
}
