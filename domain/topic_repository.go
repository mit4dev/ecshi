package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type topicRepository struct {
	coll *mongo.Collection
}

func NewTopicRepository(coll *mongo.Collection) TopicRepository {
	return &topicRepository{coll: coll}
}

func (r *topicRepository) FindById(ctx context.Context, id string) (*Topic, error) {
	var entity Topic

	err := r.coll.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&entity)

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *topicRepository) FindAll(ctx context.Context) ([]*Topic, error) {
	var entities []*Topic

	cursor, err := r.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *topicRepository) Store(ctx context.Context, entity *Topic) (*Topic, error) {
	now := time.Now()
	entity.ID = uuid.NewString()
	entity.CreatedAt = now
	entity.UpdatedAt = now

	result, err := r.coll.InsertOne(ctx, entity)

	if err != nil {
		return nil, err
	}

	entity.ID = result.InsertedID.(string)

	return entity, nil
}

func (r *topicRepository) Search(ctx context.Context, query string) ([]*Topic, error) {
	var entities []*Topic

	// TODO: Text search
	cursor, err := r.coll.Find(ctx, bson.D{{Key: "title", Value: query}})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}
