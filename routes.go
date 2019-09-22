package main

import (
	"github.com/labstack/echo"
	 TodoController "todo-api/controllers"
)

func initRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	todo := v1.Group("/todo")

	todo.GET("", TodoController.GetTodos)
	todo.GET("/:id", TodoController.GetTodo)
	todo.POST("", TodoController.CreateTodo)
	todo.PUT("/:id", TodoController.UpdateTodo)
	todo.DELETE("/:id", TodoController.DeleteTodo)
}