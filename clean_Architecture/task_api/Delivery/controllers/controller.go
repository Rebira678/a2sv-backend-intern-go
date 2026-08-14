package controllers

import (
	"net/http"

	"task-api/Domain"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	taskUsecase domain.TaskUsecase
	userUsecase domain.UserUsecase
}

func NewController(taskUc domain.TaskUsecase, userUc domain.UserUsecase) *Controller {
	return &Controller{
		taskUsecase: taskUc,
		userUsecase: userUc,
	}
}

// Task methods
func (c *Controller) CreateTask(ctx *gin.Context) {
	var task domain.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	task.UserID = user.(*domain.User).ID

	if err := c.taskUsecase.CreateTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create task"})
		return
	}

	ctx.JSON(http.StatusCreated, task)
}

func (c *Controller) GetTasks(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := user.(*domain.User).ID

	tasks, err := c.taskUsecase.GetTasksByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c *Controller) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := c.taskUsecase.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	user, exists := ctx.Get("user")
	if !exists || task.UserID != user.(*domain.User).ID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *Controller) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	
	task, err := c.taskUsecase.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	user, exists := ctx.Get("user")
	if !exists || task.UserID != user.(*domain.User).ID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	var updatedTask domain.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	if err := c.taskUsecase.UpdateTask(id, &updatedTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not update task"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
}

func (c *Controller) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := c.taskUsecase.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	user, exists := ctx.Get("user")
	if !exists || task.UserID != user.(*domain.User).ID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	if err := c.taskUsecase.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete task"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

// User methods
func (c *Controller) Register(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	if err := c.userUsecase.Register(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

func (c *Controller) Login(ctx *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := c.userUsecase.Login(credentials.Username, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
