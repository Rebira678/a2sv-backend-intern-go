package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Task struct {
	ID          bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	DueDate     time.Time     `json:"due_date" bson:"due_date"`
	Status      string        `json:"status" bson:"status"`
}
