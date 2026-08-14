package repositories

import (
	"context"

	"task-api/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) domain.TaskRepository {
	return &taskRepository{
		collection: collection,
	}
}

func (r *taskRepository) Create(task *domain.Task) error {
	task.ID = bson.NewObjectID().Hex() // Automatically generate ID if omitted
	_, err := r.collection.InsertOne(context.Background(), task)
	return err
}

func (r *taskRepository) GetByID(id string) (*domain.Task, error) {
	var task domain.Task
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) GetByUserID(userID string) ([]domain.Task, error) {
	var tasks []domain.Task

	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) Update(task *domain.Task) error {
	filter := bson.M{"_id": task.ID}
	update := bson.M{"$set": task}
	
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *taskRepository) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
