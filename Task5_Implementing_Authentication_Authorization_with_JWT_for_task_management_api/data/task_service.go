package data

import (
	"errors"
	"task_manager/models"
)

var tasks = make(map[string]models.Task)

func GetAllTasks() []models.Task {
	var taskList []models.Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return taskList
}

func GetTaskByID(id string) (*models.Task, error) {
	task, exists := tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return &task, nil
}

func CreateTask(task models.Task) {
	tasks[task.ID] = task
}

func UpdateTask(id string, updatedTask models.Task) error {
	if _, exists := tasks[id]; !exists {
		return errors.New("task not found")
	}
	updatedTask.ID = id
	tasks[id] = updatedTask
	return nil
}

func DeleteTask(id string) error {
	if _, exists := tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}
