package domain

import (
	"context"
	"time"
)

type MongoJwt struct {
	ID           string    `bson:"_id" json:"id"`
	UserId       string    `bson:"user_id" json:"user_id"`
	AccessToken  string    `bson:"access_token" json:"access_token"`
	RefreshToken string    `bson:"refresh_token" json:"refresh_token"`
	UserAgent    string    `bson:"user_agent" json:"user_agent"`
	DeviceIp     string    `bson:"device_ip" json:"device_ip"`
	IsDeleted    bool      `bson:"is_deleted" json:"is_deleted"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
}

type JwtRepository interface {
	FindByToken(ctx context.Context, token string) (*MongoJwt, error)
	Store(ctx context.Context, jwt *MongoJwt) (*MongoJwt, error)
}
