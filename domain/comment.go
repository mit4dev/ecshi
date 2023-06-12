package domain

import "time"

type Comment struct {
	ID           string     `bson:"_id" json:"id"`
	Content      string     `bson:"content" json:"content"`
	AuthorID     string     `bson:"author_id" json:"author_id"`
	TopicID      string     `bson:"topic_id" json:"topic_id"`
	CreatedAt    time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `bson:"updated_at" json:"updated_at"`
	Replies      []*Comment `bson:"replies" json:"replies"`
	IsDeleted    bool       `bson:"is_deleted" json:"is_deleted"`
	UpvoterIDs   []string   `bson:"upvoter_ids" json:"upvoter_ids"`
	DownvoterIDs []string   `bson:"downvoter_ids" json:"downvoter_ids"`
}

type CommentRepository interface {
	FindByTopicID(id string) ([]*Comment, error)
	FindById(id string) (*Comment, error)

	Store(comment *Comment) (*Comment, error)
}
