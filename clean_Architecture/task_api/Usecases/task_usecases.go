package usecases

import (
	"errors"
	"time"

	"task-api/Domain"
)

type taskUsecase struct {
	repository domain.TaskRepository
}

func NewTaskUsecase(repository domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{
		repository: repository,
	}
}

func (u *taskUsecase) CreateTask(task *domain.Task) error {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	if task.Status == "" {
		task.Status = domain.StatusPending
	}
	return u.repository.Create(task)
}

func (u *taskUsecase) GetTaskByID(id string) (*domain.Task, error) {
	return u.repository.GetByID(id)
}

func (u *taskUsecase) GetTasksByUserID(userID string) ([]domain.Task, error) {
	return u.repository.GetByUserID(userID)
}

func (u *taskUsecase) UpdateTask(id string, task *domain.Task) error {
	existingTask, err := u.repository.GetByID(id)
	if err != nil {
		return errors.New("task not found")
	}

	// Update only allowed fields
	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.Status = task.Status
	existingTask.UpdatedAt = time.Now()

	return u.repository.Update(existingTask)
}

func (u *taskUsecase) DeleteTask(id string) error {
	return u.repository.Delete(id)
}
