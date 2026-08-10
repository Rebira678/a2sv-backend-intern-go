package main

import (
	"log"

	"task_manager/data"
	"task_manager/router"
)

func main() {
	log.Println("Connecting to MongoDB...")
	if err := data.InitMongoDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB successfully!")

	r := router.SetupRouter()
	log.Println("Starting Task Manager API on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
