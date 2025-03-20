package models

// タスクのデータ構造
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}
