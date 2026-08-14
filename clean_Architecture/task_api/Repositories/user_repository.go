package repositories

import (
	"context"

	"task-api/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) domain.UserRepository {
	return &userRepository{
		collection: collection,
	}
}

func (r *userRepository) Create(user *domain.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
