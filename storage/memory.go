package storage

import (
	"errors"
	"project/models"
)

// タスク管理用のメモリ内データベース
var todos = []models.Todo{}
var idCounter = 1

// タスク一覧を取得
func GetAllTodos() []models.Todo {
	return todos
}

// ID でタスクを取得
func GetTodoByID(id int) (*models.Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("Todo not found")
}

// タスクを追加
func AddTodo(todo models.Todo) models.Todo {
	todo.ID = idCounter
	idCounter++
	todos = append(todos, todo)
	return todo
}

// タスクを更新
func UpdateTodo(id int, updatedTodo models.Todo) (*models.Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Completed = updatedTodo.Completed
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

// タスクを削除
func DeleteTodo(id int) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Todo not found")
}
