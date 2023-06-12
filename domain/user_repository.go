package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(coll *mongo.Collection) UserRepository {
	coll.Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		})

	coll.Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		})

	return &userRepository{coll: coll}
}

func (r *userRepository) FindById(ctx context.Context, id string) (*User, error) {
	var entity User
	err := r.coll.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&entity)

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *userRepository) Store(ctx context.Context, data *User) (*User, error) {

	_, err := r.coll.InsertOne(ctx, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var entity User

	err := r.coll.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&entity)

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*User, error) {
	var entities []*User

	cursor, err := r.coll.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}
