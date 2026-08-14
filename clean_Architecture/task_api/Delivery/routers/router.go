package routers

import (
	"task-api/Delivery/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controller *controllers.Controller, authMiddleware gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	// Protected routes
	protected := router.Group("/")
	protected.Use(authMiddleware)
	{
		protected.POST("/tasks", controller.CreateTask)
		protected.GET("/tasks", controller.GetTasks)
		protected.GET("/tasks/:id", controller.GetTaskByID)
		protected.PUT("/tasks/:id", controller.UpdateTask)
		protected.DELETE("/tasks/:id", controller.DeleteTask)
	}

	return router
}
