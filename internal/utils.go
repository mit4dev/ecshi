package api

import (
	"context"

	"github.com/mit4dev/ecshi/domain"
	"golang.org/x/crypto/bcrypt"
)

const (
	ctxUserKey           = "user"
	cookieAccessTokenKey = "access_token"
)

func hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func compareHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func GetCurrentUser(ctx context.Context) *domain.User {
	user, ok := ctx.Value(ctxUserKey).(*domain.User)
	if !ok {
		return nil
	}

	return user
}
