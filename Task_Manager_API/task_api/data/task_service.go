package data

import (
	"errors"
	"strconv"
	"time"
	"task_manager/models"
)

var tasks = []models.Task{
	{
		ID:          "1",
		Title:       "Setup Project",
		Description: "Initialize Go module and Gin framework",
		DueDate:     time.Now().AddDate(0, 0, 2),
		Status:      "Completed",
	},
}
var nextID = 2

func GetAllTasks() []models.Task {
	return tasks
}

func GetTaskByID(id string) (*models.Task, error) {
	for _, t := range tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(t models.Task) models.Task {
	t.ID = strconv.Itoa(nextID)
	nextID++
	tasks = append(tasks, t)
	return t
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			updatedTask.ID = id
			tasks[i] = updatedTask
			return &updatedTask, nil
		}
	}
	return nil, errors.New("task not found")
}

func DeleteTask(id string) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
