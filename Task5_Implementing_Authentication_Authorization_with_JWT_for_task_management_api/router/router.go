package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes (requires authentication)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/tasks", controllers.GetTasks)
		protected.GET("/tasks/:id", controllers.GetTaskByID)

		// Admin-only routes
		admin := protected.Group("/")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.POST("/tasks", controllers.CreateTask)
			admin.PUT("/tasks/:id", controllers.UpdateTask)
			admin.DELETE("/tasks/:id", controllers.DeleteTask)
		}
	}

	return r
}
