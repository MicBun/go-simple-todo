package route

import (
	"github.com/MicBun/go-simple-todo/controllers"
	"github.com/MicBun/go-simple-todo/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	todoRoute := r.Group("/todos")
	todoRoute.Use(middleware.JwtAuthMiddleware())
	todoRoute.POST("", controllers.CreateTodo)
	todoRoute.GET("", controllers.GetUserTodos)
	todoRoute.GET("/:id", controllers.GetTodoById)
	todoRoute.PUT("/:id", controllers.CompleteTodo)
	todoRoute.DELETE("/:id", controllers.DeleteTodo)
	todoRoute.GET("/completed", controllers.GetCompletedTodos)
	todoRoute.GET("/uncompleted", controllers.GetUncompletedTodos)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
