package domain

import (
	"context"
	"time"
)

type UserRole string

const (
	Mod    UserRole = "mod"
	Client UserRole = "client"
	Writer UserRole = "writer"
	Reader UserRole = "reader"
)

type User struct {
	Id        string     `bson:"_id" json:"id"`
	Username  string     `bson:"username" json:"username" binding:"required"`
	Password  string     `bson:"password" json:"-"`
	Email     string     `bson:"email" json:"email" binding:"required"`
	Avatar    string     `bson:"avatar" json:"avatar"`
	Roles     []UserRole `bson:"roles" json:"roles"`
	IsDeleted bool       `bson:"is_deleted" json:"is_deleted"`
	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at"`
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]*User, error)
	FindById(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)

	Store(ctx context.Context, data *User) (*User, error)
}
