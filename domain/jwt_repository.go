package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type jwtRepository struct {
	coll *mongo.Collection
}

func NewJwtRepository(coll *mongo.Collection) JwtRepository {
	return &jwtRepository{coll: coll}
}

func (r *jwtRepository) Store(ctx context.Context, jwt *MongoJwt) (*MongoJwt, error) {
	now := time.Now()
	jwt.ID = uuid.NewString()
	jwt.IsDeleted = false
	jwt.CreatedAt = now

	result, err := r.coll.InsertOne(ctx, jwt)

	if err != nil {
		return nil, err
	}

	jwt.ID = result.InsertedID.(string)

	return jwt, nil
}

func (r *jwtRepository) FindByToken(ctx context.Context, token string) (*MongoJwt, error) {
	var jwt MongoJwt

	filter := bson.M{
		"access_token": token,
	}
	err := r.coll.FindOne(ctx, filter).Decode(&jwt)

	if err != nil {
		return nil, err
	}

	return &jwt, nil
}
