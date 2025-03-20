package controllers

import (
	"net/http"
	"strconv"
	"project/models"
	"project/storage"

	"github.com/gin-gonic/gin"
)

// すべてのタスクを取得
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, storage.GetAllTodos())
}

// ID でタスクを取得
func GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := storage.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// 新しいタスクを作成
func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := storage.AddTodo(newTodo)
	c.JSON(http.StatusCreated, todo)
}

// タスクを更新
func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := storage.UpdateTodo(id, updatedTodo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// タスクを削除
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = storage.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
