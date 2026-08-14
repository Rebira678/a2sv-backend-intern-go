package main

import (
	"context"
	"log"
	"time"

	"task-api/Delivery/controllers"
	"task-api/Delivery/routers"
	"task-api/Infrastructure"
	"task-api/Repositories"
	"task-api/Usecases"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongoDB(uri string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

func main() {
	mongoURI := "mongodb://localhost:27017"

	client, err := ConnectMongoDB(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("task_api")
	taskCollection := db.Collection("tasks")
	userCollection := db.Collection("users")

	// Infrastructure
	jwtSecret := "my-secret-key"
	jwtService := infrastructure.NewJwtService(jwtSecret, "task-api")
	passwordService := infrastructure.NewPasswordService()

	// Repositories
	taskRepo := repositories.NewTaskRepository(taskCollection)
	userRepo := repositories.NewUserRepository(userCollection)

	// Usecases
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo, passwordService, jwtService)

	// Controllers
	controller := controllers.NewController(taskUsecase, userUsecase)

	// Middlewares
	authMiddleware := infrastructure.AuthMiddleware(jwtService)

	// Routers
	r := routers.SetupRouter(controller, authMiddleware)

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
