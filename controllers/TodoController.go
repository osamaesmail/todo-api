package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	db "todo-api/db"
	"todo-api/models"
)

func GetTodos(c echo.Context) error {
	var todos []models.Todo
	db := db.GetDB()
	db.Find(&todos)

	return c.JSON(http.StatusOK, todos)
}

func GetTodo(c echo.Context) error {
	id := c.Param("id")
	var todo models.Todo

	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
	}

	return c.JSON(http.StatusOK, &todo)
}

func CreateTodo(c echo.Context) error {
	var todo models.Todo
	var db = db.GetDB()

	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db.Create(&todo)

	return c.JSON(http.StatusOK, &todo)
}

func UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	var todo models.Todo

	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
	}

	c.Bind(&todo)
	db.Save(&todo)

	return c.JSON(http.StatusOK, &todo)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	var todo models.Todo
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
	}

	db.Delete(&todo)

	return c.JSON(http.StatusOK, map[string]string{"message": "Deleted successfully"})
}