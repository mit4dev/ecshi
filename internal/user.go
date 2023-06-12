package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mit4dev/ecshi/domain"
)

type PostUserDto struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
}

func (api *api) getUsers(ctx *gin.Context) {
	users, err := api.userRepo.FindAll(ctx)

	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, users)
}

func (api *api) getUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := api.userRepo.FindById(ctx, id)

	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, user)
}

func (api *api) postUser(ctx *gin.Context) {
	var data PostUserDto

	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(data.Password) <= 8 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password.invalid-length"})
		return
	}

	hashed, error := hash(data.Password)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	now := time.Now()
	user := domain.User{
		Id:       uuid.NewString(),
		Email:    data.Email,
		Password: hashed,
		Username: data.Username,
		Roles:    []domain.UserRole{domain.Reader},

		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	inserted, err := api.userRepo.Store(context.Background(), &user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": inserted})
}
