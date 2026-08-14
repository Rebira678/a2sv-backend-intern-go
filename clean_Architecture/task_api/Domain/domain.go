package domain

import "time"

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted  TaskStatus = "completed"
)

type Task struct {
	ID          string     `json:"id" bson:"_id,omitempty"`
	Title       string     `json:"title" bson:"title" binding:"required"`
	Description string     `json:"description" bson:"description"`
	Status      TaskStatus `json:"status" bson:"status"`
	UserID      string     `json:"user_id" bson:"user_id"`
	CreatedAt   time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" bson:"updated_at"`
}

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type TaskRepository interface {
	Create(task *Task) error
	GetByID(id string) (*Task, error)
	GetByUserID(userID string) ([]Task, error)
	Update(task *Task) error
	Delete(id string) error
}

type UserRepository interface {
	Create(user *User) error
	GetByUsername(username string) (*User, error)
}

type TaskUsecase interface {
	CreateTask(task *Task) error
	GetTaskByID(id string) (*Task, error)
	GetTasksByUserID(userID string) ([]Task, error)
	UpdateTask(id string, task *Task) error
	DeleteTask(id string) error
}

type UserUsecase interface {
	Register(user *User) error
	Login(username, password string) (string, error)
}

type JwtService interface {
	GenerateToken(user *User) (string, error)
	ValidateToken(tokenString string) (*User, error)
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}
