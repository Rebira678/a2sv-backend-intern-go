package data

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"task_manager/models"
)

var collection *mongo.Collection

func InitMongoDB() error {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	collection = client.Database("task_db").Collection("tasks")
	return nil
}

func GetAllTasks(ctx context.Context) ([]models.Task, error) {
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	if tasks == nil {
		tasks = []models.Task{}
	}
	return tasks, nil
}

func GetTaskByID(ctx context.Context, idStr string) (*models.Task, error) {
	objID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}

	var task models.Task
	err = collection.FindOne(ctx, bson.D{{"_id", objID}}).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func CreateTask(ctx context.Context, task models.Task) (*models.Task, error) {
	task.ID = bson.NewObjectID()
	if task.Status == "" {
		task.Status = "Pending"
	}
	if task.DueDate.IsZero() {
		task.DueDate = time.Now().AddDate(0, 0, 7)
	}

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(ctx context.Context, idStr string, updatedData models.Task) (*models.Task, error) {
	objID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}

	updateFields := bson.D{}
	if updatedData.Title != "" {
		updateFields = append(updateFields, bson.E{Key: "title", Value: updatedData.Title})
	}
	if updatedData.Description != "" {
		updateFields = append(updateFields, bson.E{Key: "description", Value: updatedData.Description})
	}
	if !updatedData.DueDate.IsZero() {
		updateFields = append(updateFields, bson.E{Key: "due_date", Value: updatedData.DueDate})
	}
	if updatedData.Status != "" {
		updateFields = append(updateFields, bson.E{Key: "status", Value: updatedData.Status})
	}

	if len(updateFields) == 0 {
		return nil, errors.New("no valid fields provided for update")
	}

	update := bson.D{{"$set", updateFields}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedTask models.Task
	err = collection.FindOneAndUpdate(ctx, bson.D{{"_id", objID}}, update, opts).Decode(&updatedTask)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return &updatedTask, nil
}

func DeleteTask(ctx context.Context, idStr string) error {
	objID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid task ID format")
	}

	res, err := collection.DeleteOne(ctx, bson.D{{"_id", objID}})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
