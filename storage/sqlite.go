package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"project/models"
)

var db *gorm.DB

// SQLite データベースの初期化
func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// タスク用のテーブルを作成（存在しない場合のみ）
	db.AutoMigrate(&models.Todo{})
}

// タスク一覧を取得
func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := db.Find(&todos)
	return todos, result.Error
}

// ID でタスクを取得
func GetTodoByID(id int) (*models.Todo, error) {
	var todo models.Todo
	result := db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

// タスクを追加
func AddTodo(todo models.Todo) (models.Todo, error) {
	result := db.Create(&todo)
	return todo, result.Error
}

// タスクを更新
func UpdateTodo(id int, updatedTodo models.Todo) (*models.Todo, error) {
	var todo models.Todo
	result := db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}

	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed
	db.Save(&todo)
	return &todo, nil
}

// タスクを削除
func DeleteTodo(id int) error {
	result := db.Delete(&models.Todo{}, id)
	return result.Error
}