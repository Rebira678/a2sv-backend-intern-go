package router

import (
	"github.com/gin-gonic/gin"

	"task_manager/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.GET("", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTaskByID)
		taskRoutes.POST("", controllers.CreateTask)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}

	return r
}
