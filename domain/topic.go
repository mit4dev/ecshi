package domain

import (
	"context"
	"time"
)

type Topic struct {
	ID        string    `bson:"_id" json:"id"`
	Title     string    `bson:"title" json:"title"`
	Content   string    `bson:"content" json:"content"`
	AuthorID  string    `bson:"author_id" json:"author_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	IsDeleted bool      `bson:"is_deleted" json:"is_deleted"`
	IsLocked  bool      `bson:"is_locked" json:"is_locked"`
	IsHidden  bool      `bson:"is_hidden" json:"is_hidden"`
	IsPinned  bool      `bson:"is_pinned" json:"is_pinned"`
	IsDraft   bool      `bson:"is_draft" json:"is_draft"`
}

type TopicRepository interface {
	FindAll(ctx context.Context) ([]*Topic, error)
	FindById(ctx context.Context, id string) (*Topic, error)
	Search(ctx context.Context, query string) ([]*Topic, error)

	Store(ctx context.Context, topic *Topic) (*Topic, error)
}
