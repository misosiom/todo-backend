package routes

import (
	"project/controllers"
	"github.com/gin-gonic/gin"
)

// ルートの設定
func SetupRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	{
		todoGroup.GET("/", controllers.GetTodos)
		todoGroup.POST("/", controllers.CreateTodo)
		todoGroup.GET("/:id", controllers.GetTodoByID)
		todoGroup.PUT("/:id", controllers.UpdateTodo)
		todoGroup.DELETE("/:id", controllers.DeleteTodo)
	}
}
